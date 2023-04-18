package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bradhe/stopwatch"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

// go test -timeout 30m -v -run ^TestRandomMovie$ bingebuddy.com/m
func TestRandomMovie(t *testing.T) {
	for i := 0; i < 10; i++ {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(httptest.NewRecorder())
		timer := stopwatch.Start()
		randomMovie(context)
		if recorder.Code != 200 {
			t.FailNow()
		}
		timer.Stop()
		time := timer.Milliseconds()
		avgTime := time / 1000
		t.Logf("Average time: %d ms", avgTime)
	}
}

// go test -timeout 30m -v -run ^TestTrueRandomMovie$ bingebuddy.com/m
// escape + i and then :wq to exit vim
// Tests average speed of random movie function ;  go test -timeout 10m -run ^TestRandomMovie$ bingebuddy.com/m
func TestTrueRandomMovie(t *testing.T) {
	//the gin framework allows you to create a test context to pass into functions that require it
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	//must hardcode generator range because test function doesn't go through normal startup process
	largest = 1104380
	smallest = 2
	//naming objects underscores makes go not force you to use them.
	timer := stopwatch.Start()
	for i := 0; i < 10; i++ {
		randomMovie(context)
	}
	timer.Stop()
	time := timer.Milliseconds()
	avgTime := time / 1000
	//error logging doesn't work so you must set a breakpoint here to see average time
	t.Logf("Average time: %d ms", avgTime)
}

// Generates random users, puts them into global array for testing
func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var testUsers map[string]User
var testPosts map[string]Post

func generateUsers(iterations int) {
	//starts with blank map and temp file
	newMap := make(map[string]User)
	file, err := os.Create("users.csv")
	//if file already exists, delete it and make a new one
	if os.IsExist(err) {
		os.Remove("users.csv")
		file, err = os.Create("users.csv")
	}
	defer file.Close() //writes to file so that login test can be run separately
	for i := 0; i < iterations; i++ {
		var user User
		rng := rand.New(rand.NewSource(time.Now().Unix()))
		usernameLength := generateRandomNumber(0, 55, *rng)
		passwordLength := generateRandomNumber(0, 55, *rng)
		user.Username = randSeq(usernameLength)
		user.Password = randSeq(passwordLength)
		_, duplicate := newMap[user.Username]
		for duplicate {
			user.Username = randSeq(usernameLength)
			user.Password = randSeq(passwordLength)
			_, duplicate = newMap[user.Username]
		}
		newMap[user.Username] = user
		profile := user.Username + "," + user.Password + "\n"
		file.WriteString(profile)
		if err != nil {
			panic(err)
		}
	}
	testUsers = newMap
}

func validationBool(user *User) bool {
	ok, _ := validateUser(user)
	return ok
}

// go test -timeout 20m -run ^TestCreateUser$ bingebuddy.com/m
func TestCreateUser(t *testing.T) {
	localMode = true
	//wipes database before creating users
	client := connectToDB()
	database := client.Database("UserInfo").Collection("UserInfo")
	filter := bson.D{{}}
	recorder := httptest.NewRecorder()
	deletionContext, _ := gin.CreateTestContext(recorder)
	database.DeleteMany(deletionContext, filter)
	//generates users
	generateUsers(50)
	for _, value := range testUsers {
		currUser := value
		marshalledUser, _ := json.Marshal(currUser)
		mock := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(mock)

		req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(marshalledUser))
		if err != nil {
			t.FailNow()
		}
		context.Request = req
		createUser(context)

		code := mock.Code
		if code != 200 && validationBool(&currUser) != false {
			t.FailNow()
		}

		//tests whether duplicates are correctly rejected.
		duplicateMock := httptest.NewRecorder()
		duplicateContext, _ := gin.CreateTestContext(duplicateMock)
		duplicateReq, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(marshalledUser))
		if err != nil {
			t.FailNow()
		}
		duplicateContext.Request = duplicateReq
		createUser(duplicateContext)
		duplicatecode := duplicateMock.Code

		if duplicatecode != 400 {
			t.FailNow()
		}
	}
}

// go test -timeout 20m -run ^TestLogin$ bingebuddy.com/m
func TestLogin(t *testing.T) {
	localMode = true
	file, err := os.Open("users.csv")
	testUsers := make(map[string]User)
	if os.IsNotExist(err) {
		return
	}
	fileScanner := bufio.NewScanner(file)
	//scans use file line by line
	for fileScanner.Scan() {
		var user User
		lineString := string(fileScanner.Bytes())
		credentials := strings.Split(lineString, ",")
		user.Username = credentials[0]
		user.Password = credentials[1]
		testUsers[user.Username] = user
	}
	//gets file of random credentials created with username
	for key, value := range testUsers {
		mock := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(mock)
		//JSONcredentials := []byte(`{
		//	"username": "` + key + `",
		//	"password": "` + value.Password + `"
		//}`)
		testCredentials := User{
			Username: key,
			Password: value.Password,
		}
		JSONcredentials, _ := json.Marshal(testCredentials)
		request, _ := http.NewRequest("GET", "/login", bytes.NewBuffer(JSONcredentials))
		context.Request = request

		login(context)
		fmt.Println()
		assert.Equal(t, http.StatusOK, mock.Code)
		assert.NotEmpty(t, mock.Body.String())
	}
	os.Remove("user.csv")
}

// tests formula used to generate random IDs
// timeout override command: go test -timeout 10m -run ^TestGenerateRandomNumber$ bingebuddy.com/m
func TestGenerateRandomNumber(t *testing.T) {
	smallest := 7.0
	largest := 113.0
	output := 0
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for output != int(smallest) {
		output = generateRandomNumber(smallest, largest, *rng)
		fmt.Println(output)
		//test automatically fails if out of bounds output is produced
		if output < int(smallest) || output > int(largest) {
			t.FailNow()
		}
	}
	for output != int(largest) {
		output = generateRandomNumber(smallest, largest, *rng)
		fmt.Println(output)
		if output < int(smallest) || output > int(largest) {
			t.FailNow()
		}
	}
}

// Command: go test -timeout 20m -run ^TestGetSimilarMovies$ bingebuddy.com/m
func TestGetSimilarMovies(t *testing.T) {
	for i := 0; i < 2000; i++ {
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		id := generateRandomNumber(2.0, 1000.0, *rng)
		frontHalf := "https://api.themoviedb.org/3/movie/"
		backHalf := "/similar?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US&page=1"
		requestString := frontHalf + fmt.Sprint(id) + backHalf
		resp, err := http.Get(requestString)
		if err != nil {
			t.FailNow()
		}
		defer resp.Body.Close()
		httpBinary, _ := io.ReadAll(resp.Body)
		mock := httptest.NewRecorder()
		mockContext, _ := gin.CreateTestContext(mock)
		//gives the id to the context when calling the function
		mockContext.Params = []gin.Param{
			{
				Key:   "id",
				Value: fmt.Sprint(id),
			},
		}
		getSimilarMovies(mockContext)
		functionBinary, _ := io.ReadAll(mock.Body)
		var r1 MovieResults
		var r2 MovieResults
		json.Unmarshal(httpBinary, &r1)
		json.Unmarshal(functionBinary, &r2)
		if !reflect.DeepEqual(r1, r2) {
			t.FailNow()
		}
	}
}

func postsIdentical(post1 *Post, post2 *Post) bool {
	if post1.Date == post2.Date && post1.Body == post2.Body && post1.Title == post2.Title {
		return true
	}
	return false
}

func generatePosts(executions int) {
	//starts with blank map
	newMap := make(map[string]Post)
	for i := 0; i < executions; i++ {
		var post Post
		rng := rand.New(rand.NewSource(time.Now().Unix()))
		titleLength := generateRandomNumber(0.0, 225.0, *rng)
		bodyLength := generateRandomNumber(0.0, 2700.0, *rng)
		post.Username = "test1234"
		post.Title = randSeq(titleLength)
		post.Body = randSeq(bodyLength)
		newMap[post.Title] = post
	}
	testPosts = newMap
}

// command: go test -timeout 300m -run ^TestCreatePost$ bingebuddy.com/m
func TestCreatePost(t *testing.T) {
	//generates token given a known valid username and account
	user := User{
		Username: "test1234",
		Password: "1234",
		Posts:    []Post{},
	}
	localMode = true
	token, _ := generateToken(user)
	generatePosts(2000)
	mock := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(mock)
	//resets profile to faciliate repeated testing
	profileBinary, _ := json.Marshal(user)
	creationReq := httptest.NewRequest("POST", "http://localhost:8080/signup", bytes.NewBuffer(profileBinary))
	context.Request = creationReq
	createUser(context)
	updateReq := httptest.NewRequest("PUT", "http://localhost:8080/test1234/update", bytes.NewBuffer(profileBinary))
	context.Request = updateReq
	updateUserInfo(context)
	for i := range testPosts {
		p := testPosts[i]
		binary, _ := json.Marshal(p)
		request := httptest.NewRequest("GET", "http://localhost:8080", bytes.NewBuffer(binary))
		request.Header = map[string][]string{
			"Authorization": {token},
		}
		context.Request = request
		createPost(context)
	}

	mock = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(mock)
	// reader can be nil if it's unnecessary
	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	request.Header = map[string][]string{
		"Authorization": {token},
		"token":         {token},
	}
	c := http.Cookie{
		Name:   "token",
		Value:  token,
		Path:   "",
		Domain: "",
	}
	request.AddCookie(&c)
	context.Request = request
	getUserInfo(context)
	binaryInfo, _ := io.ReadAll(mock.Body)
	json.Unmarshal(binaryInfo, &user)
	for _, p := range user.Posts {
		origPost := testPosts[p.Title]
		origPost.Date = time.Now().Format("January 2, 2006")
		origPost.Username = user.Username
		if !postsIdentical(&origPost, &p) {
			t.FailNow()
		}
	}
}

// command: go test -timeout 30m -run ^TestDeletePost$ bingebuddy.com/m
func TestDeletePost(t *testing.T) {
	//generates token given a known valid username and account
	user := User{
		Username: "test1234",
		Password: "1234",
		Posts:    []Post{},
	}
	localMode = true
	token, _ := generateToken(user)
	//gets all  of the test user's post IDs
	mock := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(mock)
	req := httptest.NewRequest("GET", "http://localhost:8080/user", nil)
	req.Header = map[string][]string{
		"Authorization": {token},
		"token":         {token},
	}
	c := http.Cookie{
		Name:   "token",
		Value:  token,
		Path:   "",
		Domain: "",
	}
	req.AddCookie(&c)
	context.Request = req
	getUserInfo(context)
	binary, _ := io.ReadAll(mock.Body)
	json.Unmarshal(binary, &user)
	//iterates through every post and deletes it
	for _, p := range user.Posts {
		delReq := httptest.NewRequest("DELETE", "http://localhost:8080", nil)
		context.Params = []gin.Param{
			{
				Key:   "postID",
				Value: p.PostID.Hex(),
			},
		}
		delReq.Header = map[string][]string{
			"Authorization": {token},
			"token":         {token},
		}
		c := http.Cookie{
			Name:   "token",
			Value:  token,
			Path:   "",
			Domain: "",
		}
		delReq.AddCookie(&c)
		context.Request = delReq
		deletePost(context)
	}
	mock = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(mock)
	req = httptest.NewRequest("GET", "http://localhost:8080/user", nil)
	req.Header = map[string][]string{
		"Authorization": {token},
		"token":         {token},
	}
	c = http.Cookie{
		Name:   "token",
		Value:  token,
		Path:   "",
		Domain: "",
	}
	req.AddCookie(&c)
	context.Request = req
	getUserInfo(context)
	binary, _ = io.ReadAll(mock.Body)
	json.Unmarshal(binary, &user)
	if len(user.Posts) > 0 {
		t.Fail()
	}

}

type Cast struct {
	Cast []Actor `json:"cast"`
}

type Provider struct {
	Results []US `json:"results"`
}

type US struct {
	Link     string    `json:"link"`
	Rent     []Service `json:"rent"`
	Buy      []Service `json:"buy"`
	Flatrate []Service `json:"flatrate"`
}

type Service struct {
	ProviderName string `json:"provider_name"`
	ProviderID   int    `json:"provider_id"`
}

func TestRandomMovieWithFilters(t *testing.T) {
	actorNames := []string{
		"Tom Hanks",
		"Leonardo DiCaprio",
		"Tom Cruise",
		"Will Smith",
		"Denzel Washington",
		"Johnny Depp",
		"Brad Pitt",
		"Matt Damon",
		"Samuel L. Jackson",
		"Morgan Freeman",
		"Robert Downey Jr.",
		"Anthony Hopkins",
		"Robert De Niro",
		"Jack Nicholson",
		"George Clooney",
		"Al Pacino",
		"Harrison Ford",
		"Bruce Willis",
		"Hugh Jackman",
		"Liam Neeson",
		"Matthew McConaughey",
		"Keanu Reeves",
		"Christian Bale",
		"Robin Williams",
		"Marlon Brando",
		"Nicolas Cage",
		"Joaquin Phoenix",
		"Arnold Schwarzenegger",
		"Mark Wahlberg",
		"Meryl Streep",
		"Adam Sandler",
		"Dustin Hoffman",
		"Clint Eastwood",
		"Chris Hemsworth",
		"Jamie Foxx",
		"Vin Diesel",
		"Charlton Heston",
		"Antonio Banderas",
		"James Stewart",
		"Gary Cooper",
		"Spencer Tracy",
		"Ben Affleck",
		"John Wayne",
		"Kevin Spacey",
		"Gary Oldman",
		"Kirk Douglas",
		"Don Cheadle",
		"Sandra Bullock",
		"Heath Ledger",
		"Scarlett Johansson",
		"Benedict Cumberbatch",
	}
	streamingServiceIDs := []int{
		8,   //Netflix
		9,   //Prime Video
		188, //YouTube Premium
		15,  //Hulu
		337, //Disney+
		384, //HBO Max
		387, //Peacock Premium
		283, //Crunchyroll
		350, //Apple TV+
	}
	genreIds := []int{
		28,    // Action
		12,    // Adventure
		16,    // Animation
		35,    // Comedy
		80,    // Crime
		99,    // Documentary
		18,    // Drama
		10751, // Family
		14,    // Fantasy
		36,    // History
		27,    // Horror
		10402, // Music
		9648,  // Mystery
		10749, // Romance
		878,   // Science Fiction
		10770, // TV Movie
		53,    // Thriller
		10752, // War
		37,    // Western
	}
	var totalTime time.Duration
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	var noResults []GeneratorFilters
	for i := 0; i < 4000; i++ {
		recorder := httptest.NewRecorder()
		mock, _ := gin.CreateTestContext(recorder)
		//determines number of genres and services
		actorNumber := generateRandomNumber(0, 2, *rng)
		serviceNumber := generateRandomNumber(0, 9.0, *rng)
		genreNumber := generateRandomNumber(0, 3, *rng)
		min_rating := rng.Float32() * 10
		min_rating = float32(int(min_rating*10)) / 100
		max_runtime := generateRandomNumber(60, 300, *rng)
		var actors []string
		var streaming_providers []int
		var genres []int
		for l := 0; l < actorNumber; l++ {
			actors = append(actors, actorNames[generateRandomNumber(0, 49, *rng)])
		}
		for j := 0; j < serviceNumber; j++ {
			id := streamingServiceIDs[generateRandomNumber(0, 8, *rng)]
			streaming_providers = append(streaming_providers, id)
			if id == 9 {
				streaming_providers = append(streaming_providers, 10)
			}
		}
		for k := 0; k < genreNumber; k++ {
			genres = append(genres, genreIds[generateRandomNumber(0, 17, *rng)])
		}
		filters := GeneratorFilters{
			Actors:     actors,
			MaxRuntime: max_runtime,
			Genres:     genres,
			MinRating:  min_rating,
			Providers:  streaming_providers,
		}

		JSONFilters, err := json.Marshal(filters)
		if err != nil {
			panic(err)
		}
		//cannot bind structs directly to context
		req, err := http.NewRequest("GET", "http://localhost:8080/generate/filters", bytes.NewBuffer(JSONFilters))
		mock.Request = req
		timer := stopwatch.Start()
		randomMovieWithFilters(mock)
		timer.Stop()
		totalTime += time.Millisecond
		binary, _ := io.ReadAll(recorder.Result().Body)
		var movie Movie
		json.Unmarshal(binary, &movie)
		//no results case - copies parameters for later use and skips
		if movie.Title == "" && movie.ID == 0 {
			noResults = append(noResults, filters)
			continue
		}
		//now verifies that movie information is correct
		var cast Cast
		//compares immediately accessible components
		if float32(movie.VoteAverage) < min_rating || movie.Runtime > max_runtime {
			t.FailNow()
		}
		//checks genre IDs
		for _, id := range genres {
			if !assert.Contains(t, movie.GenreIDs, id) {
				t.FailNow()
			}
		}
		//checks cast information
		requestString := "https://api.themoviedb.org/3/movie/" + strconv.Itoa(movie.ID) + "/credits?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US"
		resp, _ := http.Get(requestString)
		binary, _ = io.ReadAll(resp.Body)
		json.Unmarshal(binary, &cast)
		//copies cast names to string array
		var castNames []string
		for _, c := range cast.Cast {
			castNames = append(castNames, c.Name)
		}
		for _, n := range actors {
			if !assert.Contains(t, castNames, n) {
				t.FailNow()
			}
		}
		requestString = "https://api.themoviedb.org/3/movie/" + strconv.Itoa(movie.ID) + "/watch/providers?api_key=010c2ddcdf323db029b6dca4cbfa49de"
		resp, _ = http.Get(requestString)
		//checks providers - information comes from JustWatch API
		binary, _ = io.ReadAll(resp.Body)
		respString := string(binary)
		//uses contains on the string because I spent two hours trying to get it to marshal correctly into objects
		serviceCounter := 0
		for _, p := range streaming_providers {
			if strings.Contains(respString, strconv.Itoa(p)) {
				serviceCounter++
			}
		}
		if serviceCounter == 0 && serviceNumber > 0 {
			t.FailNow()
		}
	}
	avgTime := totalTime / 4000
	t.Logf("The average time is %d milliseconds", avgTime)
	//now checks cases with no results by making request again to confirm that there are no results
	for _, c := range noResults {
		var actorIDs []int
		var ActorResults ActorResults
		for i := 0; i < len(c.Actors); i++ {
			frontHalf := "https://api.themoviedb.org/3/search/person?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US&query="
			backHalf := "&page=1&include_adult=false"
			requestString := frontHalf + url.QueryEscape(c.Actors[i]) + backHalf
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
				//context.IndentedJSON(http.StatusOK, gin.H{"error": "no results for actor " + filters.Actors[i]})
				fmt.Printf("No results for actor" + c.Actors[i])
			} else {
				actorIDs = append(actorIDs, ActorResults.Results[0].Id)
			}
		}
		requestString := "https://api.themoviedb.org/3/discover/movie?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US&include_adult=false&include_video=false&"
		//adds the minimum rating
		requestString += ("vote_average.gte=" + fmt.Sprintf("%f", c.MinRating) + "&with_cast=")
		//loop adds actors to request
		for _, a := range actorIDs {
			requestString += (strconv.Itoa(a) + ",")
		}
		requestString += "&with_genres="
		//loop adds genres to request
		for _, g := range c.Genres {
			requestString += (strconv.Itoa(g) + ",")
		}
		//specifies maximum runtime
		requestString += ("&with_runtime.lte=" + strconv.Itoa(c.MaxRuntime))
		//adds streaming providers
		requestString += "&with_watch_providers="
		for _, p := range c.Providers {
			requestString += (strconv.Itoa(p) + "|")
		}
		//needs region flag to filter providers properly
		requestString += "&watch_region=US"
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
		if len(resultPage.Results) > 0 {
			//checks for appropriateness
			for _, m := range resultPage.Results {
				if filterMovies(&m) {
					t.FailNow()
				}
			}
		}
	}
}
