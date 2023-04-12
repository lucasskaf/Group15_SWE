package main

//imports commented out to avoid generating errors for unused

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"strconv"
	"strings"

	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"

	//"compress/gzip"
	"compress/gzip"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/microcosm-cc/bluemonday"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
)

// this is the movie struct that contains all the different fields for a movie
type Movie struct {
	Adult               bool     `json:"adult"`                                                                                             // Indicates if the movie is adult-rated.
	BackdropPath        string   `json:"backdrop_path,omitempty"`                                                                           // Path to the backdrop image for the movie.
	Budget              int      `json:"budget,omitempty"`                                                                                  // The movie's budget in dollars.
	Genres              []string `json:"genres,omitempty"`                                                                                  // The genres associated with the movie.
	Homepage            string   `json:"homepage,omitempty"`                                                                                // The movie's homepage URL.
	ID                  int      `json:"id,omitempty"`                                                                                      // The movie's unique ID.
	OriginalLanguage    string   `json:"original_language,omitempty"`                                                                       // The movie's original language code.
	OriginalTitle       string   `json:"original_title,omitempty"`                                                                          // The movie's original title.
	Overview            string   `json:"overview,omitempty"`                                                                                // A brief overview of the movie's plot.
	Popularity          float64  `json:"popularity,omitempty"`                                                                              // The movie's popularity score.
	PosterPath          string   `json:"poster_path,omitempty"`                                                                             // Path to the poster image for the movie.
	ProductionCompanies []string `json:"production_companies,omitempty"`                                                                    // The production companies involved in making the movie.
	ProductionCountries []string `json:"production_countries,omitempty"`                                                                    // The countries where the movie was produced.
	ReleaseDate         *string  `json:"release_date,omitempty"`                                                                            // The movie's release date.
	Revenue             *int     `json:"revenue,omitempty"`                                                                                 // The movie's box office revenue in dollars.
	Runtime             *int     `json:"runtime,omitempty"`                                                                                 // The movie's runtime in minutes.
	SpokenLanguages     []string `json:"spoken_languages,omitempty"`                                                                        // The languages spoken in the movie.
	Status              string   `json:"status,omitempty" validate:"oneof=rumored planned in_production post_production released canceled"` // The movie's production status.
	Tagline             string   `json:"tagline,omitempty"`                                                                                 // The movie's tagline.
	Title               string   `json:"title,omitempty"`                                                                                   // The movie's title.
	VoteAverage         float64  `json:"vote_average"`                                                                                      // The average rating given to the movie by users.
	VoteCount           int      `json:"vote_count,omitempty"`                                                                              // The number of user ratings given to the movie.
	UserRating          float32  `json:"user_rating,omitempty"`
}

type MovieResults struct {
	Results []Movie `json:"results"`
}

type ActorResults struct {
	Results []Actor `json:"results"`
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

// this is the user struct that contains all the different fields for a certain user
type User struct {
	Username      string           `json:"username"`
	Password      string           `json:"password"`
	Email         string           `json:"email"`
	Watchlist     []Movie          `json:"watchlist"`
	Posts         []Post           `json:"posts"`
	Genres        []string         `json:"genres"`
	Rating        float32          `json:"rating"`
	Subscriptions []string         `json:"subscriptions"`
	ActiveFilters GeneratorFilters `json:"active_filters"`
}

type ForumComment struct {
	Commenter string    `json:"commenter"`
	Timestamp time.Time `json:"timestamp"`
	Body      string    `json:"body"`
	Score     int       `json:"score"`
}

type GeneratorParameters struct {
	LastUpdated time.Time `json:"lastUpdated"`
	Largest     int       `json:"largest"`
	Smallest    int       `json:"smallest"`
}

type GeneratorFilters struct {
	//Actors and genres have to be comma separated lists of IDs
	Actors     []string `json:"actors"`
	MaxRuntime int      `json:"max_runtime"`
	Genres     []int    `json:"genres"`
	MinRating  float32  `json:"min_rating"`
	Providers  []int    `json:"streaming_providers"`
}

type Actor struct {
	Id int `json:"id"`
}

// global generator parameters
var largest float64
var smallest float64

// toggles database mode for testing - local has no speed limit
var localMode bool

// this is the post struct that contains all the different fields for a certain post
type Post struct {
	PostID   primitive.ObjectID `json:"id"`
	MovieID  string             `json:"movie_id"`
	Username string             `json:"username"`
	Title    string             `json:"title"`
	Body     string             `json:"body"`
	Date     string             `json:"date"`
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
	//online cluster mongodb+srv://test:1234@appdata.1fywcdf.mongodb.net/?retryWrites=true&w=majority
	//local cluster for testing mongodb://localhost:27017/
	var URI string
	if localMode {
		URI = "mongodb://localhost:27017/"
	} else {
		URI = "mongodb+srv://test:1234@appdata.1fywcdf.mongodb.net/?retryWrites=true&w=majority"
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
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
	//sanitizes user profile before searching database
	sanitizeUser(&credentials)
	filter := bson.D{{Key: "username", Value: credentials.Username}, {Key: "password", Value: credentials.Password}}
	var retrieved User
	err := database.FindOne(context, filter).Decode(&retrieved)
	//database.Find(context, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// prints debug message and sends back empty JSON struct if password is wrong
			fmt.Println("username or password is incorrect")
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

	context.JSON(http.StatusOK, gin.H{"token": token})

	//context.IndentedJSON(http.StatusOK, retrieved)
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
		context.JSON(http.StatusAlreadyReported, gin.H{"error": "u r an idiot"})
		return //catches null requests and throws error.
	}

	valid, errString := validateUser(&newUser)
	//throws error if username or password are blank
	if !valid {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		client.Disconnect(context)
		return
	}

	//checks for duplicate username
	var duplicate User
	filter := bson.D{{Key: "username", Value: newUser.Username}}
	err := database.FindOne(context, filter).Decode(&duplicate)
	if err != mongo.ErrNoDocuments {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Username is taken"})
		return
	}
	newUser.Posts = []Post{}
	database.InsertOne(context, newUser)
	context.IndentedJSON(http.StatusOK, newUser)
	client.Disconnect(context)
}

// checks if user meets certain conditions for account creation (put multiple return types in parentheses separated by commas)
func validateUser(user *User) (bool, string) {
	sanitizeUser(user)
	error := ""
	isValid := true
	if user.Username == "" && user.Password == "" {
		isValid = false
		error = "username or password cannot be blank"
		return isValid, error
	}
	userLen := len(user.Username)
	passLen := len(user.Password)
	if userLen < 4 || passLen < 4 {
		error = "username and password must be at least 4 characters"
		return isValid, error
	}
	if userLen > 50 || passLen > 50 {
		isValid = false
		error = "username or password must be less than 50 characters"
		return isValid, error
	}
	return isValid, error
}

func sanitizeUser(user *User) {
	//this policy strips all HTML tags from every part of the user class to prevent XSS attacks
	p := bluemonday.StrictPolicy()
	user.Email = p.Sanitize(user.Email)
	user.Username = p.Sanitize(user.Username)
	user.Password = p.Sanitize(user.Password)
	for i, g := range user.Genres {
		user.Genres[i] = p.Sanitize(g)
	}
	for _, m := range user.Watchlist {
		sanitizeMovieFields(&m, p)
	}
	for i, s := range user.Subscriptions {
		user.Subscriptions[i] = p.Sanitize(s)
	}

}

// sanitizes the fields that the user is likely to know and input
func sanitizeMovieFields(movie *Movie, policy *bluemonday.Policy) {
	//policy can be passed in for greater efficiency, but function can still operate independently
	if policy == nil {
		policy = bluemonday.StrictPolicy()
	}
	movie.Title = policy.Sanitize(movie.Title)
	for i, g := range movie.Genres {
		movie.Genres[i] = policy.Sanitize(g)
	}
}

func addToWatchlist(context *gin.Context) {
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

	client := connectToDB()
	database := client.Database("UserInfo").Collection("UserInfo")
	var movie Movie
	if err := context.BindJSON(&movie); err != nil {
		fmt.Printf("JSON bind failed!")
		return //catches null requests and throws error.
	}
	sanitizeMovieFields(&movie, nil)
	if movie.OriginalTitle == "" {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}
	filter := bson.D{{Key: "username", Value: username}}
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
	filter := bson.D{{Key: "username", Value: username}, {"$inc", bson.D{{"$pull", movie.Title}}}}
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
	client := connectToDB()
	database := client.Database("UserInfo").Collection("UserInfo")
	filter := bson.D{{Key: "username", Value: username}}
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

	duplicateFilter := bson.D{{Key: "username", Value: updatedUser.Username}}
	updateFilter := bson.D{{Key: "username", Value: username}}

	//checks whether desired username already exists
	error := database.FindOne(context, duplicateFilter).Decode(&currProfile)
	if error == mongo.ErrNoDocuments || username == updatedUser.Username {
		database.FindOneAndReplace(context, updateFilter, updatedUser)
		context.IndentedJSON(http.StatusOK, updatedUser)
		client.Disconnect(context)
	} else {
		//throws error if username is duplicate
		context.IndentedJSON(http.StatusBadRequest, currProfile)
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
	appropriate := false
	executions := 1
	var randMovie Movie
	for appropriate == false {
		url := "https://api.themoviedb.org/3/movie/top_rated?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US&page="
		//should produce a random number from 1 to 1000
		randPage := generateRandomNumber(1, 70)
		url = url + fmt.Sprint(randPage)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		var results MovieResults
		binary, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(binary, &results)
		pageSize := len(results.Results)
		executions := 0
		randIndex := generateRandomNumber(0, float64(pageSize-1))
		randMovie = results.Results[randIndex]
		appropriate = filterMovies(&randMovie)
		executions++
	}
	println(executions)
	//returns an empty struct and an error if function failed to produce a random movie.
	if randMovie.Title == "" {
		context.IndentedJSON(http.StatusInternalServerError, randMovie)
	} else {
		context.IndentedJSON(http.StatusOK, randMovie)
	}
}

func randomMovieWithFilters(context *gin.Context) {
	var filters GeneratorFilters
	filters.MaxRuntime = 4294967295
	filters.MinRating = 0
	context.BindJSON(&filters)
	//first assembles actor IDs for query
	var actorIDs []int
	var ActorResults ActorResults

	for i := 0; i < len(filters.Actors); i++ {
		frontHalf := "https://api.themoviedb.org/3/search/person?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US&query="
		backHalf := "&page=1&include_adult=false"
		requestString := frontHalf + url.QueryEscape(filters.Actors[i]) + backHalf
		resp, err := http.Get(requestString)
		if err != nil {
			panic(err)
		}
		binary, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(binary, &ActorResults)
		//checks if requested actor exists
		if len(ActorResults.Results) == 0 {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "no results for actor " + filters.Actors[i]})
		} else {
			actorIDs = append(actorIDs, ActorResults.Results[0].Id)
		}
	}
	requestString := "https://api.themoviedb.org/3/discover/movie?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US&include_adult=false&include_video=false&page=1&"
	//adds the minimum rating
	requestString += ("vote_average.gte=" + fmt.Sprintf("%f", filters.MinRating) + "&with_cast=")
	//loop adds actors to request
	for _, a := range actorIDs {
		requestString += (strconv.Itoa(a) + ",")
	}
	requestString += "&with_genres="
	//loop adds genres to request
	for _, g := range filters.Genres {
		requestString += (strconv.Itoa(g) + ",")
	}
	//specifies maximum runtime
	requestString += ("&with_runtime.lte=" + strconv.Itoa(filters.MaxRuntime))
	//adds streaming providers
	requestString += "&with_watch_providers="
	for _, p := range filters.Providers {
		requestString += (strconv.Itoa(p) + ",")
	}
	resp, err := http.Get(requestString)
	if err != nil {
		panic(err)
	}
	binary, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//makes request with full string
	var resultPage MovieResults
	resp, err = http.Get(requestString)
	if err != nil {
		panic(err)
	}
	binary, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(binary, &resultPage)
	pageSize := len(resultPage.Results)
	if pageSize == 0 {
		context.IndentedJSON(http.StatusOK, gin.H{"error": "No results"})
		return
	} else {
		index := generateRandomNumber(0, float64(pageSize-1))
		result := resultPage.Results[index]
		context.IndentedJSON(http.StatusOK, result)
	}
}

func trueRandomMovie(context *gin.Context) {
	frontHalf := "https://api.themoviedb.org/3/movie/"
	backHalf := "?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US"
	var resp *http.Response
	var err error
	executions := 0
	//resp is nil by default!
	//stores response body in binary
	appropriate := false
	var movieData Movie
	//first execution must take place outside of loop
	//If invalid, makes requests until it gets an OK response
	for !appropriate {
		id := generateRandomNumber(smallest, largest)
		requestString := frontHalf + fmt.Sprint(id) + backHalf
		resp, err = http.Get(requestString)
		if err != nil {
			log.Fatal(err)
		}
		//restarts process if ID is invalid
		if resp.StatusCode != 200 {
			continue
		}
		//filtering mechanism
		binary, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(binary, &movieData)
		//known adult ID for filter test : 75312
		appropriate = filterMovies(&movieData)
		executions++
	}
	//prints out number of subsequent requests made - for testing ONLY
	fmt.Println(executions)
	//converts the binary output intof a string for return
	//takes the string and sends it back to frontend as JSON
	context.JSON(http.StatusOK, movieData)
}

func generateRandomNumber(smallest float64, largest float64) int {
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	output := int(((rng.Float64() * (largest - smallest)) + smallest) + 0.5)
	return output
}

func analyzePopularPages() {
	for i := 1; i <= 549; i++ {
		var page MovieResults
		resp, err := http.Get("https://api.themoviedb.org/3/movie/top_rated?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US&page=" + fmt.Sprint(i))
		if err != nil {
			panic(err)
		}
		binary, _ := io.ReadAll(resp.Body)
		json.Unmarshal(binary, &page)
		println(len(page.Results))
	}
}

func filterMovies(m *Movie) bool {
	//checks if movie contains adult content
	if m.Adult {
		return false
	}
	//movie must have a rating above 0
	if m.VoteAverage == 0 {
		return false
	}
	//checks if movie is in English
	en := strings.Contains(m.OriginalLanguage, "en")
	return en
}

func getSimilarMovies(context *gin.Context) {
	id := context.Param("id")
	frontHalf := "https://api.themoviedb.org/3/movie/"
	backHalf := "/similar?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US&page=1"
	requestString := frontHalf + id + backHalf
	resp, err := http.Get(requestString)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var results MovieResults
	json.Unmarshal(body, &results)
	context.JSON(http.StatusOK, results)
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
	valid, errorString := validatePost(&newPost)
	if !valid {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": errorString})
		return
	}
	date := time.Now().Format("January 2, 2006")
	// Add/insert new created post into database ForumPosts collection ForumPosts for storage
	postDatabase := client.Database("ForumPosts").Collection("ForumPosts")
	result, err := postDatabase.InsertOne(context, newPost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	newPost.PostID = result.InsertedID.(primitive.ObjectID)
	newPost.Date = date
	newPost.Username = username

	userDatabase := client.Database("UserInfo").Collection("UserInfo")
	filter := bson.M{"username": username}
	updateUserPosts := bson.M{"$push": bson.M{"posts": newPost}}
	_, err = userDatabase.UpdateOne(context, filter, updateUserPosts)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add post to user's post array"})
	}

	context.JSON(http.StatusCreated, newPost)
	// fmt.Println("Post successfuly created")
	client.Disconnect(context)
}

func sanitizePost(post *Post) {
	policy := bluemonday.NewPolicy()
	policy.AllowStandardURLs()
	policy.AllowRelativeURLs(true)
	policy.AllowImages()
	post.Title = policy.Sanitize(post.Title)
	post.Body = policy.Sanitize(post.Body)
}

func validatePost(post *Post) (bool, string) {
	valid := true
	var error string
	if post.Body == "" || post.Title == "" {
		valid = false
		error = "post title and body cannot be blank"
		return valid, error
	}
	if len(post.Title) > 100 || len(post.Body) > 1000 {
		valid = false
		error = "post title or body is too long"
		return valid, error
	}
	return valid, error
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

// this function updates an already-existing post for the logged in user
func updatePost(context *gin.Context) {
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

	postID := context.Param("postID")
	objectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "This is an invalid post ID"})
		return
	}

	claims, _ := userToken.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	client := connectToDB()

	// Need to get the post that needs to be updated
	postDatabase := client.Database("ForumPosts").Collection("ForumPosts")
	filter := bson.M{"_id": objectID}
	var currentPost Post
	err = postDatabase.FindOne(context, filter).Decode(&currentPost)
	if err != nil { // post does not exist
		context.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
	}
	// Check if the logged in user created post that is trying to be updated
	if currentPost.Username != username {
		context.JSON(http.StatusForbidden, gin.H{"error": "You cannot update a post that is not yours"})
		return
	}

	// Update the current post with whatever logged in user wants
	var updatedPost Post
	if err := context.BindJSON(&updatedPost); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to the parse updated post from request body"})
		return
	}
	valid, errorString := validatePost(&updatedPost)
	if !valid {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": errorString})
		return
	}
	updatedPost.Date = time.Now().Format("January 2, 2006")
	updateMade := bson.M{
		"$set": bson.M{
			"title": updatedPost.Title,
			"body":  updatedPost.Body,
			"date":  updatedPost.Date,
		},
	}

	_, err = postDatabase.UpdateOne(context, filter, updateMade)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	userDatabase := client.Database("UserInfo").Collection("UserInfo")
	updateUserPosts := bson.M{
		"$set": bson.M{
			"posts.$.title": updatedPost.Title,
			"posts.$.body":  updatedPost.Body,
		},
	}
	updateFilter := bson.M{"username": username, "posts.postid": objectID}

	_, err = userDatabase.UpdateOne(context, updateFilter, updateUserPosts)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post in user array"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
	client.Disconnect(context)
}

func getPosts(context *gin.Context) {
	id := context.Param("id")
	page := context.Param("page")
	//converts page number into an integer and handles invalid inputs
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	client := connectToDB()
	database := client.Database("ForumPosts").Collection("ForumPosts")
	filter := bson.D{{Key: "movieid", Value: id}}
	//have to set options to sort posts from most to least recent and limit the amount of retrievals
	opts := options.Find().SetLimit(int64(pageInt) * 50).SetSort(bson.D{{"$natural", -1}})
	//database.FindOne(context, filter).Decode(&post)
	cursor, err := database.Find(context, filter, opts)
	if err == mongo.ErrNoDocuments {
		context.IndentedJSON(http.StatusOK, gin.H{"error": "no posts found"})
	}
	//marshals every result into the array
	var posts []Post
	if err = cursor.All(context, &posts); err != nil {
		panic(err)
	}
	if len(posts) < 50 {
		context.IndentedJSON(http.StatusOK, posts)
	} else {
		lowerBound := (pageInt - 1) * 50
		upperBound := pageInt * 50
		//corrects for out of bounds page requests
		if upperBound > len(posts) {
			upperBound = len(posts)
			lowerBound = upperBound - 50
		}
		posts = posts[lowerBound:upperBound]
		context.IndentedJSON(http.StatusOK, posts)
	}
	client.Disconnect(context)
}

func getUserInfo(context *gin.Context) {
	header := context.GetHeader("Authorization")
	headerToken := strings.ReplaceAll(header, "Bearer ", "")
	userToken, err := jwt.Parse(headerToken, func(userToken *jwt.Token) (interface{}, error) {
		if _, ok := userToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", userToken.Header["alg"])
		}
		return []byte("sayhellotomylittlefriend"), nil
	})

	claims, _ := userToken.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	client := connectToDB()

	database := client.Database("UserInfo").Collection("UserInfo")
	filter := bson.D{{Key: "username", Value: username}}
	var user User
	err = database.FindOne(context, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// prints debug message and sends back empty JSON struct if password is wrong
			fmt.Println("username is invalid")
			var emptyStruct User
			context.IndentedJSON(http.StatusOK, emptyStruct)
			return
		}
		panic(err)
	}
	//obscures sensitive data
	user.Password = ""
	user.Email = ""
	context.IndentedJSON(http.StatusOK, user)
}

// Checks ID range with API automatically on startup
func updateGeneratorParameters() {
	client := connectToDB()
	database := client.Database("GeneratorParameters").Collection("GeneratorParameters")
	context := context.Background()
	var parameters GeneratorParameters
	/*DELETE THIS LATER - FOR DEBUGGING*/
	firstExecution := false
	//finds parameters
	filter := bson.D{{}}
	database.FindOne(context, filter).Decode(&parameters)
	var lastUpdated time.Duration
	if parameters.Largest != 0 {
		lastUpdated = time.Since(parameters.LastUpdated)
	}
	//performs update if it's been more than 24 hours
	if firstExecution || lastUpdated.Hours() > 24.00 {
		//gets current date for request to API
		Date := time.Now()
		year, month, day := Date.Date()
		//converts date elements to strings
		var monthString = strconv.Itoa(int(month))
		var dayString = strconv.Itoa(day)
		//adds leading zeroes if necessary
		if len(monthString) == 1 {
			monthString = "0" + monthString
		}

		if len(dayString) == 1 {
			dayString = "0" + dayString
		}
		//puts request string together
		requestString := "http://files.tmdb.org/p/exports/movie_ids_" + monthString + "_" + dayString + "_" + strconv.Itoa(year) + ".json.gz"
		//requests file from database
		resp, err := http.Get(requestString)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		//checks to see if data dump has been published yet
		if resp.StatusCode == 200 {
			//creates temporary file for reading
			file, err := os.Create("validIDs")
			if err != nil {
				panic(err)
			}
			//deletes file after scan is finished
			defer os.Remove("validIDs")

			//writes http response body to temp file
			_, err = io.Copy(file, resp.Body)
			if err != nil {
				panic(err)
			}
			//closes initial file writing
			file.Close()
			//opens file again for decompression
			gzipFile, err := os.Open("validIDs")
			if err != nil {
				panic(err)
			}
			//creates a destination for uncompressed file
			out, err := os.Create("Uncompressed.json")
			if err != nil {
				panic(err)
			}
			defer os.Remove("Uncompressed.json")
			//decompresses original file stream
			reader, err := gzip.NewReader(gzipFile)
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(out, reader)
			if err != nil {
				panic(err)
			}
			out.Close()
			reader.Close()
			//opens file again so that scanner will work
			scannerFile, err := os.Open("Uncompressed.json")
			if err != nil {
				panic(err)
			}
			defer scannerFile.Close()
			//scans file line by line
			fileScanner := bufio.NewScanner(scannerFile)
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
			}

			//inserts these parameters into database
			parameters.Largest = largest
			parameters.Smallest = smallest
			parameters.LastUpdated = time.Now()
			database.FindOneAndReplace(context, filter, parameters)
			fmt.Println("Largest: " + fmt.Sprint(largest))
			fmt.Println("Smallest: " + fmt.Sprint(smallest))
			fmt.Println("Database updated!")
		} else {
			fmt.Println("no update to download!")
		}
	} else {
		fmt.Println("database did not need to be updated!")
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control")
		context.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}
		context.Next()
	}
}

func main() {
	//Checks for database updates on startup
	updateGeneratorParameters()
	client := connectToDB()
	database := client.Database("GeneratorParameters").Collection("GeneratorParameters")
	filter := bson.D{{}}
	var parameters GeneratorParameters
	database.FindOne(context.Background(), filter).Decode(&parameters)
	largest = float64(parameters.Largest)
	smallest = float64(parameters.Smallest)
	client.Disconnect(context.Background())
	localMode = false
	//Uncomment out to speed up unit tests by deleting debug message
	os.Setenv("GIN_MODE", "release")
	gin.SetMode(gin.ReleaseMode)
	//Sets up routing
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/login", login)
	router.GET("/me", getUserInfo)
	router.GET("/generate", randomMovie)
	router.GET("/generate/similar/:id", getSimilarMovies)
	router.GET("/posts/:id/:page", getPosts)
	router.POST("/generate/filters", randomMovieWithFilters)
	router.POST("/signup", createUser)
	router.POST("/:username/add", addToWatchlist)
	router.POST("/posts", createPost)
	router.DELETE("/posts/:postID", deletePost)
	router.PUT("/posts/:postID", updatePost)
	router.PUT("/:username/update", updateUserInfo)
	router.DELETE("/:username/delete", removeUser)
	router.DELETE("/:username/watchlist/remove", removeFromWatchlist)
	router.Run("localhost:8080")
}
