package main

//imports commented out to avoid generating errors for unused

import (
	"context"

	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"

	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"net/http"

	"github.com/gin-gonic/gin"
)

type movie struct {
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
	Watchlist     []movie  `json:"watchlist"`
	Genres        []string `json:"genres"`
	Rating        float32  `json:"rating"`
	Subscriptions []string `json:"subscriptions"`
}

func connectToDB() (client *mongo.Client) {
	if err := godotenv.Load("go.env"); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
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
	router.GET("/watchlist")
	router.Run("localhost:8080")
}
