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
	Title   string
	Runtime float32
}

// the names of fields MUST be uppercase or else MongoDB will NOT store them
type User struct {
	Username  string   `json:"username"`
	Password  string   `json:"password"`
	Watchlist []movie  `json:"watchlist"`
	Genres    []string `json:"genres"`
	Rating    float32  `json:"rating"`
	Providers []string `json:"providers"`
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
			// This error means your query did not match any documents.
			fmt.Printf("username or password is incorrect")
			return
		}
		panic(err)
	}
	context.IndentedJSON(http.StatusOK, retrieved)
	fmt.Printf("login successful!")
	client.Disconnect(context)
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
	router.Run("localhost:8080")
}
