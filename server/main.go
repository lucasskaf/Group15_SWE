package main

//imports commented out to avoid generating errors for unused

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"
  
	"bufio"
	"math/rand"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title      string   `json:"title"`
	Director   string   `json:"director"`
	Imglink    string   `json:"imglink"`
	Runtime    float32  `json:"runtime"`
	Avgrating  float32  `json:"avgrating"`
	Providers  []string `json:"providers"`
	DatabaseID int      `json:"databaseid"`
}

// struct for getting IDs from movie database
type parseStruct struct {
	Adult             bool   `json:"adult"`
	Id                int    `json:"id"`
	Original_Language string `json:"original_language"`
	//Original_Title string  `json:"original_title"`
	//Popularity     float32 `json:"popularity"`
	//Video          bool    `json:"video"`
}

// the names of fields MUST be uppercase or else MongoDB will NOT store them
type User struct {
	Username      string   `json:"username"`
	Password      string   `json:"password"`
	Email         string   `json:"email"`
	Watchlist     []Movie  `json:"watchlist"`
	Genres        []string `json:"genres"`
	Rating        float32  `json:"rating"`
	Subscriptions []string `json:"subscriptions"`
}

type ForumPost struct {
	Poster    string         `json:"poster"`
	Timestamp time.Time      `json:"timestamp"` //golang standard struct
	Body      string         `json:"body"`
	Score     int            `json:"score"`
	Comments  []ForumComment `json:"comments"`
}

type ForumComment struct {
	Commenter string    `json:"commenter"`
	Timestamp time.Time `json:"timestamp"`
	Body      string    `json:"body"`
	Score     int       `json:"score"`
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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://test:1234@cluster0.aruhgq1.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	return client
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
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// prints debug message and sends back empty JSON struct if password is wrong
			fmt.Printf("username or password is incorrect")
			var emptyStruct User
			context.IndentedJSON(http.StatusBadRequest, emptyStruct)
			return
		}
		panic(err)
	}
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
	//throws error if username or password are blank
	if newUser.Username == "" || newUser.Password == "" {
		var emptyStruct User
		context.IndentedJSON(http.StatusBadRequest, emptyStruct)
		client.Disconnect(context)
		return
	}

	//checks for duplicate username
	var duplicate User
	filter := bson.D{{"username", newUser.Username}}
	err := database.FindOne(context, filter).Decode(&duplicate)
	if err != mongo.ErrNoDocuments {
		fmt.Printf("username is taken")
		var emptyStruct User
		context.IndentedJSON(http.StatusBadRequest, emptyStruct)
		return
	}
	database.InsertOne(context, newUser)
	client.Disconnect(context)
}

func addToWatchlist(context *gin.Context) {
	username := context.Param("username")
	client := connectToDB()
	database := client.Database("UserInfo").Collection("UserInfo")
	var movie Movie
	if err := context.BindJSON(&movie); err != nil {
		fmt.Printf("JSON bind failed!")
		return //catches null requests and throws error.
	}
	filter := bson.D{{"username", username}}
	var updatedUser User
	database.FindOne(context, filter).Decode(&updatedUser)
	updatedUser.Watchlist = append(updatedUser.Watchlist, movie)
	//FindOneAndUpdate doesn't work
	oldDoc := database.FindOneAndReplace(context, filter, updatedUser)
	//panics if document cannot be updated
	if oldDoc == nil {
		context.IndentedJSON(http.StatusBadRequest, oldDoc)
		fmt.Printf("Dcoument can't be updated")
		return
	}
	context.IndentedJSON(http.StatusOK, updatedUser)
	client.Disconnect(context)
}

func removeFromWatchlist(context *gin.Context) {
	//should take in movie object
	username := context.Param("username")
	client := connectToDB()
	database := client.Database("UserInfo").Collection("UserInfo")
	var movie Movie
	if err := context.BindJSON(&movie); err != nil {
		fmt.Printf("JSON bind failed!")
		return //catches null requests and throws error.
	}
	//filter := bson.D{{"username.watchlist", movie.Title}}
	filter := bson.D{{"username", username}, {"$inc", bson.D{{"$pull", movie.Title}}}}
	result := database.FindOneAndDelete(context, filter)
	//returns error if deletion fails
	if result == nil {
		context.IndentedJSON(http.StatusBadRequest, result)
		client.Disconnect(context)
	}
	context.IndentedJSON(http.StatusOK, result)
	client.Disconnect(context)
}

func removeUser(context *gin.Context) {
	username := context.Param("username")
	client := connectToDB()
	database := client.Database("UserInfo").Collection("UserInfo")
	filter := bson.D{{"username", username}}
	result := database.FindOneAndDelete(context, filter)
	//returns error if user doesn't exist
	if result == nil {
		context.IndentedJSON(http.StatusBadRequest, result)
	}
	context.IndentedJSON(http.StatusOK, result)
	client.Disconnect(context)
}

// generic function that replaces one user profile in database with an updated one
func updateUserInfo(context *gin.Context) {
	username := context.Param("username")
	client := connectToDB()
	database := client.Database("UserInfo").Collection("UserInfo")
	var updatedUser User
	var currProfile User
	if err := context.BindJSON(&updatedUser); err != nil {
		fmt.Printf("JSON bind failed!")
		return //catches null requests and throws error.
	}
	//checks for blank username and password
	if updatedUser.Username == "" || updatedUser.Password == "" {
		context.IndentedJSON(http.StatusBadRequest, updatedUser)
		client.Disconnect(context)
		return
	}

	duplicateFilter := bson.D{{"username", updatedUser.Username}}
	updateFilter := bson.D{{"username", username}}

	//checks whether desired username already exists
	err := database.FindOne(context, duplicateFilter).Decode(&currProfile)
	if err == mongo.ErrNoDocuments {
		database.FindOneAndReplace(context, updateFilter, updatedUser)
		context.IndentedJSON(http.StatusOK, updatedUser)
		client.Disconnect(context)
	} else {
		//throws error if username is duplicate
		context.IndentedJSON(http.StatusBadRequest, username)
		client.Disconnect(context)
	}
}

/*
	Credit for movie API goes to The Movie DB (TMDB)

"This product uses the TMDB API but is not endorsed or certified by TMDB." - Put this in the frontend
our API key: 010c2ddcdf323db029b6dca4cbfa49de
As of 2/18/2022, the largest possible movie ID is 1088411, while the smallest possible movie ID is 2
*/
func randomMovie(context *gin.Context) {
	//rng uses current time as a seed
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	frontHalf := "https://api.themoviedb.org/3/movie/"
	backHalf := "?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US"
	var resp *http.Response
	var err error
	executions := 0
	//resp is nil by default!

	id := int((rng.Float64() * 1088409) + 2)
	requestString := frontHalf + fmt.Sprint(id) + backHalf
	resp, err = http.Get(requestString)
	if err != nil {
		log.Fatal(err)
	}
	//stores response body in binary
	var body []byte
	appropriate := false
	//If invalid, makes requests until it gets an OK response
	for resp.StatusCode != 200 || !appropriate {
		//replace numbers with variables later- formula is rng times max - min plus min
		id = int((rng.Float64() * 1088409) + 2)
		requestString = frontHalf + fmt.Sprint(id) + backHalf
		resp, err = http.Get(requestString)
		if err != nil {
			log.Fatal(err)
		}
		//filtering mechanism
		binary, err := io.ReadAll(resp.Body)
		var movieData parseStruct
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(binary, &movieData)
		appropriate = filterMovies(&movieData)
		body = binary
		executions++
	}
	//prints out number of subsequent requests made - for testing ONLY
	fmt.Println(executions)
	//converts the binary output intof a string for return
	JSONstring := string(body)
	//takes the string and sends it back to frontend as JSON
	context.JSON(http.StatusOK, JSONstring)
}

func filterMovies(m *parseStruct) bool {
	//checks if movie contains adult content
	if m.Adult {
		return false
	}
	//checks if movie is in English
	en := strings.Contains(m.Original_Language, "en")
	if !en {
		return false
	}
	return true
}
func createPost(context *gin.Context) {

}

/*
Scans local API database to check for largest and smallest possible movie IDs

How to use:
1. Comment out all of the router functions in main
2. Call this function in main
3. Profit
*/
func scanValidIDs() {
	//maybe implement automatic fetch and unzipping
	//ids start at 2 for some reason
	file, err := os.Open("movie_ids_02_18_2023.json")
	if err != nil {
		panic("file opening failed!")
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	largest := 0
	smallest := 4294967295
	for fileScanner.Scan() {
		var lineStruct parseStruct
		//gets line of JSON file
		binaryLine := fileScanner.Bytes()
		//unmarshals binary into a struct
		json.Unmarshal(binaryLine, &lineStruct)
		if lineStruct.Id > largest {
			largest = lineStruct.Id
		}
		if lineStruct.Id < smallest {
			smallest = lineStruct.Id
		}
		//database.InsertOne(context.Background(), lineStruct)

	}
	fmt.Println("Largest: " + fmt.Sprint(largest))
	fmt.Println("Smallest: " + fmt.Sprint(smallest))
}

func main() {
	//Sets up routing

	router := gin.Default()
	router.GET("/login", login)
	router.GET("/generate", randomMovie)
	router.POST("/signup", createUser)
	router.POST("/:username/add", addToWatchlist)
	router.POST("/forum/post", createPost)
	router.PUT("/:username/update", updateUserInfo)
	router.DELETE("/:username/delete", removeUser)
	router.DELETE("/:username/watchlist/remove", removeFromWatchlist)
	router.Run("localhost:8080")

	//scanValidIDs()
  
	Genres    []string  `json:"genres"`
	Title     string    `json:"title"`
	Runtime   float32   `json:"runtime"`
	Rating    int32     `json:"rating"`
	Providers []string  `json:"providers"`  

}

// the names of fields MUST be uppercase or else MongoDB will NOT store them
type User struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
	Watchlist []Movie  `json:"watchlist"`
	// Genres    []string `json:"genres"`
	// Rating    float32  `json:"rating"`
	// Providers []string `json:"providers"`
}

func connectToDB() (client *mongo.Client) {
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
	return client
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
	database.InsertOne(context, newUser)
	client.Disconnect(context)
}

// func watchMovie(context *gin.Context) {
// 	client := connectToDB()
// 	database := client.Database("UserInfo").Collection("MoviesWatched")
// 	var newMovie Movie
// 	if err := context.BindJSON(&newMovie); err != nil {
// 		fmt.Printf("JSON bind failed!")
// 		return //catches null requests and throws error.
// 	}
// 	database.InsertOne(context, newMovie)
// 	client.Disconnect(context)
// }

// func watchlist(context *gin.Context) {
// 	client := connectToDB() // connect to database
// 	database := client.Database("UserInfo").Collection("MoviesWatched") // using the movies watched collection

// 	var newMovie Movie
// 	if err := context.BindJSON(&newMovie); err != nil {
// 		fmt.Printf("JSON bind failed!")
// 		return //catches null requests and throws error.
// 	}
// 	database.InsertOne(context, newMovie)
// 	client.Disconnect(context)

// 	var credentials Movie
// 	if err := context.BindJSON(&credentials); err != nil {
// 		fmt.Printf("Json binding failed")
// 	}
	
// 	filter := bson.D{{"title", credentials.Title}, {"genre", credentials.Genres}}
// 	var retrieved Movie
// 	err := database.FindOne(context, filter).Decode(&retrieved)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			// This means that current movie has not been found for the user, so it wont get added
// 			fmt.Println("movie has not been watched")
// 			var emptyStruct Movie
// 			context.IndentedJSON(http.StatusOK, emptyStruct)
// 			return
// 		}
// 		panic(err)
// 	}
// 	context.IndentedJSON(http.StatusOK, retrieved)
// 	fmt.Printf("Movie watched/added")
// 	client.Disconnect(context)
// }

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
	// router.GET("/watchlist", watchlist)
	
	router.Run("localhost:8080")
	
}