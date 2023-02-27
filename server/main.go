package main

//imports commented out to avoid generating errors for unused

import (
	"context"
	"strings"

	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title     string   `json:"title"`
	Director  string   `json:"director"`
	Imglink   string   `json:"imglink"`
	Runtime   float32  `json:"runtime"`
	Avgrating float32  `json:"avgrating"`
	Providers []string `json:"providers"`
}

// the names of fields MUST be uppercase or else MongoDB will NOT store them
type User struct {
	Username      string   `json:"username"`
	Password      string   `json:"password"`
	Email         string   `json:"email"`
	Watchlist     []Movie  `json:"watchlist"`
	Posts         []Post   `json:"posts"`
	Genres        []string `json:"genres"`
	Rating        float32  `json:"rating"`
	Subscriptions []string `json:"subscriptions"`
}

type Post struct {
	PostID primitive.ObjectID `json:"id"`
	Title  string             `json:"title"`
	Body   string             `json:"body"`
}

func connectToDB() (client *mongo.Client) {
	if err := godotenv.Load("go.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	//online cluster mongodb+srv://test:1234@cluster0.aruhgq1.mongodb.net/?retryWrites=true&w=majority
	//local cluster URL mongodb://localhost:27017/
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}

// this function is needed to generate a unique token for the logged in user, which is used to authorize the user when wanting to send any requests such as creating a post
func generateToken(currentUser User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["Authorized"] = true
	claims["username"] = currentUser.Username
	claims["expirationDate"] = time.Now().Add(time.Hour * 24).Unix()

	tokString, err := token.SignedString([]byte("sayhellotomylittlefriend"))
	if err != nil {
		return "", err
	}
	return tokString, nil
}
func login(context *gin.Context) {
	client := connectToDB()
	var credentials User
	database := client.Database("UserInfo").Collection("UserInfo")
	if err := context.BindJSON(&credentials); err != nil {
		fmt.Printf("Json binding failed")
	}

	filter := bson.D{{"username", credentials.Username}, {"password", credentials.Password}}
	var retrieved User
	err := database.FindOne(context, filter).Decode(&retrieved)
	//database.Find(context, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// prints debug message and sends back empty JSON struct if password is wrong
			fmt.Printf("username or password is incorrect")
			var emptyStruct User
			context.IndentedJSON(http.StatusOK, emptyStruct)
			return
		}
		panic(err)
	}

	token, err := generateToken(retrieved)
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})

	context.IndentedJSON(http.StatusOK, retrieved)
	fmt.Printf("login successful!")
	client.Disconnect(context)
}

func createUser(context *gin.Context) {
	client := connectToDB()
	database := client.Database("UserInfo").Collection("UserInfo")
	var newUser User
	if err := context.BindJSON(&newUser); err != nil {
		fmt.Printf("JSON bind failed!")
		return //catches null requests and throws error.
	}
	newUser.Posts = []Post{}
	database.InsertOne(context, newUser)
	client.Disconnect(context)
}

func createPost(context *gin.Context) {
	header := context.GetHeader("Authorization") // gets "Bearer token"
	if header == "" {                            // checks if the authorization header is empty or not and throws error if it is
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized User"})
		return
	}
	headerToken := strings.ReplaceAll(header, "Bearer ", "") // gets the token only, which is everything after "Bearer"
	// Now we parse through the token and check that it is valid, if not, then error
	userToken, err := jwt.Parse(headerToken, func(userToken *jwt.Token) (interface{}, error) {
		if _, ok := userToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", userToken.Header["alg"])
		}
		return []byte("sayhellotomylittlefriend"), nil
	})
	if err != nil || !userToken.Valid {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized User"})
		return
	}

	// Using claims and the token, we get the username that has this token
	claims, _ := userToken.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	client := connectToDB() // connect to MongoDB database

	// Create a new post
	var newPost Post
	if err := context.BindJSON(&newPost); err != nil {
		fmt.Printf("JSON bind failed!")
		return
	}

	// Add/insert new created post into database ForumPosts collection ForumPosts for storage
	database := client.Database("ForumPosts").Collection("ForumPosts")
	result, err := database.InsertOne(context, newPost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	postID := result.InsertedID.(primitive.ObjectID)
	userDatabase := client.Database("UserInfo").Collection("UserInfo")
	filter := bson.M{"username": username}
	updateUserPosts := bson.M{"$push": bson.M{"posts": newPost}}
	_, err = userDatabase.UpdateOne(context, filter, updateUserPosts)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add post to user's post array"})
	}

	newPost.PostID = postID
	context.JSON(http.StatusCreated, newPost)
	fmt.Println("Post successfuly created")
	client.Disconnect(context)
}

func main() {
	//database connection boilerplate
	/*
		if err := godotenv.Load("go.env"); err != nil {
			log.Println("No .env file found")
		}
		uri := os.Getenv("MONGODB_URI")
		if uri == "" {
			log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
		}
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()*/

	//alias to easily access database
	//database := client.Database("UserInfo").Collection("UserInfo")
	//comment out insertion of test user if they are already in database
	/*

		newUser := User{Username: "test", Password: "1234"}
		result, err := database.InsertOne(context.TODO(), newUser)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	*/

	//Sets up routing
	router := gin.Default()
	router.GET("/login", login)
	router.POST("/signup", createUser)
	router.POST("/posts", createPost)
	router.Run("localhost:8080")
}
