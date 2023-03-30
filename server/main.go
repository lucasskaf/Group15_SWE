// package main

// //imports commented out to avoid generating errors for unused

// import (
// 	"context"
// 	"encoding/json"
// 	"io"
// 	"strings"

// 	//"gorm.io/driver/sqlite"
// 	//"gorm.io/gorm"

// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"

// 	"github.com/gin-gonic/gin"
// )

// // this is the movie struct that contains all the different fields for a movie
// type Movie struct {
// 	Title      string   `json:"title"`
// 	Director   string   `json:"director"`
// 	Imglink    string   `json:"imglink"`
// 	Runtime    float32  `json:"runtime"`
// 	Avgrating  float32  `json:"avgrating"`
// 	Providers  []string `json:"providers"`
// 	DatabaseID int      `json:"databaseid"`
// }

// // struct for getting IDs from movie database
// type parseStruct struct {
// 	Adult             bool   `json:"adult"`
// 	Id                int    `json:"id"`
// 	Original_Language string `json:"original_language"`
// 	//Original_Title string  `json:"original_title"`
// 	//Popularity     float32 `json:"popularity"`
// 	//Video          bool    `json:"video"`
// }

// // this is the user struct that contains all the different fields for a certain user
// type User struct {
// 	Username      string   `json:"username"`
// 	Password      string   `json:"password"`
// 	Email         string   `json:"email"`
// 	Watchlist     []Movie  `json:"watchlist"`
// 	Posts         []Post   `json:"posts"`
// 	Genres        []string `json:"genres"`
// 	Rating        float32  `json:"rating"`
// 	Subscriptions []string `json:"subscriptions"`
// }

// type ForumPost struct {
// 	Poster    string         `json:"poster"`
// 	Timestamp time.Time      `json:"timestamp"` //golang standard struct
// 	Body      string         `json:"body"`
// 	Score     int            `json:"score"`
// 	Comments  []ForumComment `json:"comments"`
// }

// type ForumComment struct {
// 	Commenter string    `json:"commenter"`
// 	Timestamp time.Time `json:"timestamp"`
// 	Body      string    `json:"body"`
// 	Score     int       `json:"score"`
// }

// // this is the post struct that contains all the different fields for a certain post
// type Post struct {
// 	PostID primitive.ObjectID `json:"id"`
// 	Title  string             `json:"title"`
// 	Body   string             `json:"body"`
// 	Date   string             `json:"date"`
// }

// // this function connects the server/client to mongodb database whenever it is called
// func connectToDB() (client *mongo.Client) {
// 	if err := godotenv.Load("go.env"); err != nil {
// 		log.Println("No .env file found")
// 	}
// 	uri := os.Getenv("MONGODB_URI")
// 	if uri == "" {
// 		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
// 	}
// 	//online cluster mongodb+srv://test:1234@cluster0.aruhgq1.mongodb.net/?retryWrites=true&w=majority
// 	//local cluster URL mongodb://localhost:27017/

// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://test:1234@cluster0.aruhgq1.mongodb.net/?retryWrites=true&w=majority"))
// 	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

// 	if err != nil {
// 		panic(err)
// 	}
// 	return client
// }

// // this function is needed to generate a unique token for the logged in user, which is used to authorize the user when wanting to send any requests such as creating a post
// func generateToken(currentUser User) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["Authorized"] = true
// 	claims["username"] = currentUser.Username
// 	claims["expirationDate"] = time.Now().Add(time.Hour * 24).Unix()

// 	tokString, err := token.SignedString([]byte("sayhellotomylittlefriend"))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokString, nil
// }

// // this function authenticates the user that is trying to log in and provides the unique token for said user
// func login(context *gin.Context) {
// 	client := connectToDB()
// 	var credentials User
// 	database := client.Database("UserInfo").Collection("UserInfo")
// 	if err := context.BindJSON(&credentials); err != nil {
// 		fmt.Printf("Json binding failed")
// 	}

// 	filter := bson.D{{"username", credentials.Username}, {"password", credentials.Password}}
// 	var retrieved User
// 	err := database.FindOne(context, filter).Decode(&retrieved)
// 	//database.Find(context, filter)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			// prints debug message and sends back empty JSON struct if password is wrong
// 			fmt.Println("username or password is incorrect")
// 			var emptyStruct User
// 			context.IndentedJSON(http.StatusOK, emptyStruct)
// 			return
// 		}
// 		panic(err)
// 	}

// 	token, err := generateToken(retrieved)
// 	if err != nil {
// 		panic(err)
// 	}

// 	context.JSON(http.StatusOK, gin.H{"token": token})

// 	context.IndentedJSON(http.StatusOK, retrieved)
// 	fmt.Printf("login successful!")
// 	client.Disconnect(context)
// }

// // this function creates a brand new user and inserts it into the database
// func createUser(context *gin.Context) {
// 	client := connectToDB()
// 	database := client.Database("UserInfo").Collection("UserInfo")
// 	var newUser User
// 	if err := context.BindJSON(&newUser); err != nil {
// 		fmt.Printf("JSON bind failed!")
// 		return //catches null requests and throws error.
// 	}
// 	//throws error if username or password are blank
// 	if newUser.Username == "" || newUser.Password == "" {
// 		var emptyStruct User
// 		context.IndentedJSON(http.StatusBadRequest, emptyStruct)
// 		client.Disconnect(context)
// 		return
// 	}

// 	//checks for duplicate username
// 	var duplicate User
// 	filter := bson.D{{"username", newUser.Username}}
// 	err := database.FindOne(context, filter).Decode(&duplicate)
// 	if err != mongo.ErrNoDocuments {
// 		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Username is taken"})
// 		return
// 	}

// 	database.InsertOne(context, newUser)
// 	client.Disconnect(context)
// }

// func addToWatchlist(context *gin.Context) {
// 	username := context.Param("username")
// 	client := connectToDB()
// 	database := client.Database("UserInfo").Collection("UserInfo")
// 	var movie Movie
// 	if err := context.BindJSON(&movie); err != nil {
// 		fmt.Printf("JSON bind failed!")
// 		return //catches null requests and throws error.
// 	}
// 	filter := bson.D{{"username", username}}
// 	var updatedUser User
// 	database.FindOne(context, filter).Decode(&updatedUser)
// 	updatedUser.Watchlist = append(updatedUser.Watchlist, movie)
// 	//FindOneAndUpdate doesn't work
// 	oldDoc := database.FindOneAndReplace(context, filter, updatedUser)
// 	//panics if document cannot be updated
// 	if oldDoc == nil {
// 		context.IndentedJSON(http.StatusBadRequest, oldDoc)
// 		fmt.Printf("Dcoument can't be updated")
// 		return
// 	}
// 	context.IndentedJSON(http.StatusOK, updatedUser)
// 	client.Disconnect(context)
// }

// func removeFromWatchlist(context *gin.Context) {
// 	//should take in movie object
// 	username := context.Param("username")
// 	client := connectToDB()
// 	database := client.Database("UserInfo").Collection("UserInfo")
// 	var movie Movie
// 	if err := context.BindJSON(&movie); err != nil {
// 		fmt.Printf("JSON bind failed!")
// 		return //catches null requests and throws error.
// 	}
// 	//filter := bson.D{{"username.watchlist", movie.Title}}
// 	filter := bson.D{{"username", username}, {"$inc", bson.D{{"$pull", movie.Title}}}}
// 	result := database.FindOneAndDelete(context, filter)
// 	//returns error if deletion fails
// 	if result == nil {
// 		context.IndentedJSON(http.StatusBadRequest, result)
// 		client.Disconnect(context)
// 	}
// 	context.IndentedJSON(http.StatusOK, result)
// 	client.Disconnect(context)
// }

// func removeUser(context *gin.Context) {
// 	username := context.Param("username")
// 	client := connectToDB()
// 	database := client.Database("UserInfo").Collection("UserInfo")
// 	filter := bson.D{{"username", username}}
// 	result := database.FindOneAndDelete(context, filter)
// 	//returns error if user doesn't exist
// 	if result == nil {
// 		context.IndentedJSON(http.StatusBadRequest, result)
// 	}
// 	context.IndentedJSON(http.StatusOK, result)
// 	client.Disconnect(context)
// }

// // generic function that replaces one user profile in database with an updated one
// func updateUserInfo(context *gin.Context) {
// 	username := context.Param("username")
// 	client := connectToDB()
// 	database := client.Database("UserInfo").Collection("UserInfo")
// 	var updatedUser User
// 	var currProfile User
// 	if err := context.BindJSON(&updatedUser); err != nil {
// 		fmt.Printf("JSON bind failed!")
// 		return //catches null requests and throws error.
// 	}
// 	//checks for blank username and password
// 	if updatedUser.Username == "" || updatedUser.Password == "" {
// 		context.IndentedJSON(http.StatusBadRequest, updatedUser)
// 		client.Disconnect(context)
// 		return
// 	}

// 	duplicateFilter := bson.D{{"username", updatedUser.Username}}
// 	updateFilter := bson.D{{"username", username}}

// 	//checks whether desired username already exists
// 	err := database.FindOne(context, duplicateFilter).Decode(&currProfile)
// 	if err == mongo.ErrNoDocuments {
// 		database.FindOneAndReplace(context, updateFilter, updatedUser)
// 		context.IndentedJSON(http.StatusOK, updatedUser)
// 		client.Disconnect(context)
// 	} else {
// 		//throws error if username is duplicate
// 		context.IndentedJSON(http.StatusBadRequest, username)
// 		client.Disconnect(context)
// 	}
// }

// /*
// 	Credit for movie API goes to The Movie DB (TMDB)

// "This product uses the TMDB API but is not endorsed or certified by TMDB." - Put this in the frontend
// our API key: 010c2ddcdf323db029b6dca4cbfa49de
// As of 2/18/2022, the largest possible movie ID is 1088411, while the smallest possible movie ID is 2
// */
// func randomMovie(context *gin.Context) {
// 	//rng uses current time as a seed
// 	rng := rand.New(rand.NewSource(time.Now().Unix()))
// 	frontHalf := "https://api.themoviedb.org/3/movie/"
// 	backHalf := "?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US"
// 	var resp *http.Response
// 	var err error
// 	executions := 0
// 	//resp is nil by default!

// 	id := int((rng.Float64() * 1088409) + 2)
// 	requestString := frontHalf + fmt.Sprint(id) + backHalf
// 	resp, err = http.Get(requestString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	appropriate := false
// 	//If invalid, makes requests until it gets an OK response
// 	for resp.StatusCode != 200 || appropriate == false {
// 		//replace numbers with variables later- formula is rng times max - min plus min
// 		id = int((rng.Float64() * 1088409) + 2)
// 		requestString = frontHalf + fmt.Sprint(id) + backHalf
// 		resp, err = http.Get(requestString)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		//filtering mechanism
// 		binary, err := io.ReadAll(resp.Body)
// 		var movieData parseStruct
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		json.Unmarshal(binary, &movieData)
// 		appropriate = filterMovies(&movieData)
// 		executions++
// 	}
// 	//prints out number of subsequent requests made
// 	fmt.Println(executions)
// 	defer resp.Body.Close()
// 	//reads body of response and converts it into binary
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//converts the binary output into a string for return
// 	JSONstring := string(body)
// 	//takes the string and sends it back to frontend as JSON
// 	context.JSON(http.StatusOK, JSONstring)
// }

// func filterMovies(m *parseStruct) bool {
// 	//checks if movie contains adult content
// 	if m.Adult == false {
// 		return false
// 	}
// 	//checks if movie is in English
// 	en := strings.Contains(m.Original_Language, "en")
// 	if en == false {
// 		return false
// 	}
// 	return true
// }

// // this function creates a new post for the logged in user
// func createPost(context *gin.Context) {
// 	header := context.GetHeader("Authorization") // gets "Bearer token"
// 	if header == "" {                            // checks if the authorization header is empty or not and throws error if it is
// 		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized User"})
// 		return
// 	}
// 	headerToken := strings.ReplaceAll(header, "Bearer ", "") // gets the token only, which is everything after "Bearer"
// 	// Now we parse through the token and check that it is valid, if not, then error
// 	userToken, err := jwt.Parse(headerToken, func(userToken *jwt.Token) (interface{}, error) {
// 		if _, ok := userToken.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", userToken.Header["alg"])
// 		}
// 		return []byte("sayhellotomylittlefriend"), nil
// 	})
// 	if err != nil || !userToken.Valid {
// 		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized User"})
// 		return
// 	}

// 	// Using claims and the token, we get the username that has this token
// 	claims, _ := userToken.Claims.(jwt.MapClaims)
// 	username := claims["username"].(string)

// 	client := connectToDB() // connect to MongoDB database

// 	// Create a new post
// 	var newPost Post
// 	if err := context.BindJSON(&newPost); err != nil {
// 		fmt.Printf("JSON bind failed!")
// 		return
// 	}

// 	date := time.Now().Format("January 2, 2006")
// 	// Add/insert new created post into database ForumPosts collection ForumPosts for storage
// 	postDatabase := client.Database("ForumPosts").Collection("ForumPosts")
// 	result, err := postDatabase.InsertOne(context, bson.M{
// 		"title": newPost.Title,
// 		"body":  newPost.Body,
// 		"date":  date,
// 	})
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
// 		return
// 	}

// 	newPost.PostID = result.InsertedID.(primitive.ObjectID)
// 	newPost.Date = date

// 	userDatabase := client.Database("UserInfo").Collection("UserInfo")
// 	filter := bson.M{"username": username}
// 	updateUserPosts := bson.M{"$push": bson.M{"posts": newPost}}
// 	_, err = userDatabase.UpdateOne(context, filter, updateUserPosts)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add post to user's post array"})
// 	}

// 	context.JSON(http.StatusCreated, newPost)
// 	fmt.Println("Post successfuly created")
// 	client.Disconnect(context)
// }

// // this function deletes a post for the logged in user
// func deletePost(context *gin.Context) {
// 	header := context.GetHeader("Authorization")
// 	if header == "" {
// 		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized User"})
// 		return
// 	}

// 	headerToken := strings.ReplaceAll(header, "Bearer ", "")
// 	userToken, err := jwt.Parse(headerToken, func(userToken *jwt.Token) (interface{}, error) {
// 		if _, ok := userToken.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", userToken.Header["alg"])
// 		}
// 		return []byte("sayhellotomylittlefriend"), nil
// 	})

// 	if err != nil || !userToken.Valid {
// 		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized User"})
// 		return
// 	}

// 	postID := context.Param("postID")
// 	objectID, err := primitive.ObjectIDFromHex(postID)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "This is an invalid post ID"})
// 		return
// 	}

// 	claims, _ := userToken.Claims.(jwt.MapClaims)
// 	username := claims["username"].(string)
// 	client := connectToDB()

// 	userDatabase := client.Database("UserInfo").Collection("UserInfo")
// 	filter := bson.M{"username": username, "posts.postid": objectID}
// 	checker, err := userDatabase.CountDocuments(context, filter)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user's posts"})
// 		return
// 	}
// 	if checker == 0 {
// 		context.JSON(http.StatusForbidden, gin.H{"error": "You cannot delete a post that is not yours!"})
// 		return
// 	}

// 	postDatabase := client.Database("ForumPosts").Collection("ForumPosts")
// 	_, err = postDatabase.DeleteOne(context, bson.M{"_id": objectID})
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post from ForumPosts collection"})
// 		return
// 	}

// 	update := bson.M{"$pull": bson.M{"posts": bson.M{"postid": objectID}}}
// 	_, err = userDatabase.UpdateOne(context, filter, update)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove post from user array"})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "Post deleted successfuly"})
// 	client.Disconnect(context)
// }

// /*
// Scans local API database to check for largest and smallest possible movie IDs
// How to use:
// 1. Comment out all of the router functions in main
// 2. Call this function in main
// 3. Profit
// */
// func scanValidIDs() {
// 	//maybe implement automatic fetch and unzipping
// 	//ids start at 2 for some reason
// 	file, err := os.Open("movie_ids_02_18_2023.json")
// 	if err != nil {
// 		panic("file opening failed!")
// 	}
// 	defer file.Close()
// 	fileScanner := bufio.NewScanner(file)
// 	largest := 0
// 	smallest := 4294967295
// 	for fileScanner.Scan() {
// 		var lineStruct parseStruct
// 		//gets line of JSON file
// 		binaryLine := fileScanner.Bytes()
// 		//unmarshals binary into a struct
// 		json.Unmarshal(binaryLine, &lineStruct)
// 		if lineStruct.Id > largest {
// 			largest = lineStruct.Id
// 		}
// 		if lineStruct.Id < smallest {
// 			smallest = lineStruct.Id
// 		}
// 		//database.InsertOne(context.Background(), lineStruct)

// 	}
// 	fmt.Println("Largest: " + fmt.Sprint(largest))
// 	fmt.Println("Smallest: " + fmt.Sprint(smallest))
// }

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		context.Header("Access-Control-Allow-Origin", "*")
// 		context.Header("Access-Control-Allow-Credentials", "true")
// 		context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		context.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

// 		if context.Request.Method == "OPTIONS" {
// 			context.AbortWithStatus(204)
// 			return
// 		}

// 		context.Next()
// 	}
// }

// func main() {
// 	//Sets up routing

// 	router := gin.Default()
// 	router.Use(CORSMiddleware())
// 	router.GET("/login", login)
// 	router.GET("/generate", randomMovie)
// 	router.POST("/signup", createUser)
// 	router.POST("/:username/add", addToWatchlist)
// 	router.POST("/posts", createPost)
// 	router.DELETE("/posts/:postID", deletePost)
// 	router.PUT("/:username/update", updateUserInfo)
// 	router.DELETE("/:username/delete", removeUser)
// 	router.DELETE("/:username/watchlist/remove", removeFromWatchlist)
// 	router.Run("localhost:8080")
// }
