package main

import (
  "context"
  "strings"

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

// this is the movie struct that contains all the different fields for a movie
type Movie struct {
  Title     string   `json:"title"`
  Director  string   `json:"director"`
  Imglink   string   `json:"imglink"`
  Runtime   float32  `json:"runtime"`
  Avgrating float32  `json:"avgrating"`
  Providers []string `json:"providers"`
}

// this is the user struct that contains all the different fields for a certain user
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

// this is the post struct that contains all the different fields for a certain post
type Post struct {
  PostID primitive.ObjectID `json:"id"`
  Title  string             `json:"title"`
  Body   string             `json:"body"`
  Date   string             `json:"date"`
}

// this function connects the server/client to mongodb database whenever it is called
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

// this function authenticates the user that is trying to log in and provides the unique token for said user
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

// this function creates a brand new user and inserts it into the database
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

// this function creates a new post for the logged in user
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

  date := time.Now().Format("January 2, 2006")
  // Add/insert new created post into database ForumPosts collection ForumPosts for storage
  postDatabase := client.Database("ForumPosts").Collection("ForumPosts")
  result, err := postDatabase.InsertOne(context, bson.M{
    "title": newPost.Title,
    "body":  newPost.Body,
    "date":  date,
  })
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
    return
  }

  newPost.PostID = result.InsertedID.(primitive.ObjectID)
  newPost.Date = date

  userDatabase := client.Database("UserInfo").Collection("UserInfo")
  filter := bson.M{"username": username}
  updateUserPosts := bson.M{"$push": bson.M{"posts": newPost}}
  _, err = userDatabase.UpdateOne(context, filter, updateUserPosts)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add post to user's post array"})
  }

  context.JSON(http.StatusCreated, newPost)
  fmt.Println("Post successfuly created")
  client.Disconnect(context)
}

// this function deletes a post for the logged in user
func deletePost(context *gin.Context) {
  header := context.GetHeader("Authorization")
  if header == "" {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized User"})
    return
  }

  headerToken := strings.ReplaceAll(header, "Bearer ", "")
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

  postID := context.Param("postID")
  objectID, err := primitive.ObjectIDFromHex(postID)
  if err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": "This is an invalid post ID"})
    return
  }

  claims, _ := userToken.Claims.(jwt.MapClaims)
  username := claims["username"].(string)
  client := connectToDB()

  userDatabase := client.Database("UserInfo").Collection("UserInfo")
  filter := bson.M{"username": username, "posts.postid": objectID}
  checker, err := userDatabase.CountDocuments(context, filter)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check user's posts"})
    return
  }
  if checker == 0 {
    context.JSON(http.StatusForbidden, gin.H{"error": "You cannot delete a post that is not yours!"})
    return
  }

  postDatabase := client.Database("ForumPosts").Collection("ForumPosts")
  _, err = postDatabase.DeleteOne(context, bson.M{"_id": objectID})
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post from ForumPosts collection"})
    return
  }

  update := bson.M{"$pull": bson.M{"posts": bson.M{"postid": objectID}}}
  _, err = userDatabase.UpdateOne(context, filter, update)
  if err != nil {
    context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove post from user array"})
    return
  }

  context.JSON(http.StatusOK, gin.H{"message": "Post deleted successfuly"})
  client.Disconnect(context)
}

func main() {
  //Sets up routing
  router := gin.Default()
  router.GET("/login", login)
  router.POST("/signup", createUser)
  router.POST("/posts", createPost)
  router.DELETE("/posts/:postID", deletePost)
  router.Run("localhost:8080")
}

