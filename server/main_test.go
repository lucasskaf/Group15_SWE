package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/bradhe/stopwatch"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
  "strings"
)

func TestRandomMovie(t *testing.T) {
	for i := 0; i < 10000; i++ {
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(httptest.NewRecorder())
		randomMovie(context)
		if recorder.Code != 200 {
			t.Fail()
		}
	}
}

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
	for i := 0; i < 1000; i++ {
		randomMovie(context)
	}
	timer.Stop()
	time := timer.Milliseconds()
	avgTime := time / 5000
	//error logging doesn't work so you must set a breakpoint here to see average time
	fmt.Printf("Average time: %d ms", avgTime)
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

func generateUsers() {
	//starts with blank map
	newMap := make(map[string]User)
	for i := 0; i < 500; i++ {
		var user User
		rng := rand.New(rand.NewSource(time.Now().Unix()))

		usernameLength := int((rng.Float64() * 400))
		passwordLength := int((rng.Float64() * 400))
		user.Username = randSeq(usernameLength)
		user.Password = randSeq(passwordLength)
		newMap[user.Username] = user
	}
	testUsers = newMap
}

func checkDuplicate(user User) bool {
	if user.Username == "" || user.Password == "" {
		return true
	}
	_, duplicate := testUsers[user.Username]
	return duplicate
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
	generateUsers()
	for _, value := range testUsers {
		currUser := value
		marshalledUser, _ := json.Marshal(currUser)
		mock := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(mock)

		req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(marshalledUser))
		if err != nil {
			t.Fail()
		}
		context.Request = req
		createUser(context)

		code := mock.Code
		if code != 200 && checkDuplicate(currUser) == false {
			t.Fail()
		}

		//tests whether duplicates are correctly rejected.
		duplicateMock := httptest.NewRecorder()
		duplicateContext, _ := gin.CreateTestContext(duplicateMock)
		duplicateReq, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(marshalledUser))
		if err != nil {
			t.Fail()
		}
		duplicateContext.Request = duplicateReq
		createUser(duplicateContext)
		duplicatecode := duplicateMock.Code

		if duplicatecode != 400 {
			t.Fail()
		}
	}
}

// go test -timeout 20m -run ^TestLogin$ bingebuddy.com/m
func TestLogin(t *testing.T) {
	localMode = true
	for key, value := range testUsers {
		mock := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(mock)
		JSONcredentials := []byte(`{
			"username": "` + key + `",
			"password": "` + value.Password + `"
		}`)
		request, _ := http.NewRequest("GET", "/login", bytes.NewBuffer(JSONcredentials))
		context.Request = request

		login(context)
		fmt.Println()
		assert.Equal(t, http.StatusOK, mock.Code)
		assert.NotEmpty(t, mock.Body.String())
	}
}

// tests formula used to generate random IDs
// timeout override command: go test -timeout 10m -run ^TestGenerateRandomNumber$ bingebuddy.com/m
func TestGenerateRandomNumber(t *testing.T) {
	smallest := 1.0
	largest := 100.0
	output := 0
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for output != int(smallest) {
		output = generateRandomNumber(smallest, largest, *rng)
		fmt.Println(output)
		//test automatically fails if out of bounds output is produced
		if output < int(smallest) || output > int(largest) {
			t.Fail()
		}
	}
	for output != int(largest) {
		output = generateRandomNumber(smallest, largest, *rng)
		fmt.Println(output)
		if output < int(smallest) || output > int(largest) {
			t.Fail()
		}
	}
}

// Command: go test -timeout 10m -run ^TestGetSimilarMovies$ bingebuddy.com/m
func TestGetSimilarMovies(t *testing.T) {
	for i := 0; i < 500; i++ {
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		id := generateRandomNumber(2.0, 1000.0, *rng)
		frontHalf := "https://api.themoviedb.org/3/movie/"
		backHalf := "/similar?api_key=010c2ddcdf323db029b6dca4cbfa49de&language=en-US&page=1"
		requestString := frontHalf + fmt.Sprint(id) + backHalf
		resp, err := http.Get(requestString)
		if err != nil {
			t.Fail()
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
			t.Fail()
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

		usernameLength := int((rng.Float64() * 300))
		passwordLength := int((rng.Float64() * 300))
		post.Title = randSeq(usernameLength)
		post.Body = randSeq(passwordLength)
		newMap[post.Title] = post
	}
	testPosts = newMap
}

// command: go test -timeout 30m -run ^TestCreatePost$ bingebuddy.com/m
func TestCreatePost(t *testing.T) {
	//generates token given a known valid username and account
	user := User{
		Username: "test1324",
		Password: "1234",
		Posts:    []Post{},
	}
	localMode = true
	token, _ := generateToken(user)
	generatePosts(300)
	mock := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(mock)
	//resets profile to faciliate repeated testing
	profileBinary, _ := json.Marshal(user)
	updateReq := httptest.NewRequest("GET", "http://localhost:8080/updateUserInfo", bytes.NewBuffer(profileBinary))
	updateReq.Header = map[string][]string{
		"Authorization": {token},
	}
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
	}
	context.Request = request
	getUserInfo(context)
	binaryInfo, _ := io.ReadAll(mock.Body)
	json.Unmarshal(binaryInfo, &user)
	for _, p := range user.Posts {
		origPost := testPosts[p.Title]
		origPost.Date = time.Now().Format("January 2, 2006")
		if !postsIdentical(&origPost, &p) {
			t.Fail()
		}
	}
}

// command: go test -timeout 30m -run ^TestDeletePost$ bingebuddy.com/m
func TestDeletePost(t *testing.T) {
	//generates token given a known valid username and account
	user := User{
		Username: "test1324",
		Password: "1234",
		Posts:    []Post{},
	}
	localMode = true
	token, _ := generateToken(user)
	//gets all  of the test user's post IDs
	mock := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(mock)
	req := httptest.NewRequest("GET", "http://localhost:8080/me", nil)
	req.Header = map[string][]string{
		"Authorization": {token},
	}
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
		}
		context.Request = delReq
		deletePost(context)
	}

	//checks if every post has been deleted
	mock = httptest.NewRecorder()
	context, _ = gin.CreateTestContext(mock)
	req = httptest.NewRequest("GET", "http://localhost:8080/me", nil)
	req.Header = map[string][]string{
		"Authorization": {token},
	}
	context.Request = req
	getUserInfo(context)
	binary, _ = io.ReadAll(mock.Body)
	json.Unmarshal(binary, &user)
	for i := range user.Posts {
		//fails test if the user has any posts remaining that have not been deleted
		if i > 0 {
			t.Fail()
		}
	}

}

// go test -timeout 30m -run ^TestUpdatePost$ bingebuddy.com/m
func TestUpdatePost(t *testing.T) {
  localMode = true
  // Create a test user
  user := User{
    Username: "Albert",
    Password: "Gator",
    Posts:    []Post{},
  }

  // Connect to the test database
  dbURI := "mongodb+srv://test:1234@cluster0.gmfsqnv.mongodb.net/test"
  clientOptions := options.Client().ApplyURI(dbURI)
  client, err := mongo.Connect(context.Background(), clientOptions)
  if err != nil {
    t.Fatalf("Failed to connect to test database: %v", err)
  }
  defer client.Disconnect(context.Background())

  // Insert the test user and post into the test database
  usersCollection := client.Database("UserInfo").Collection("UserInfo")
  _, err = usersCollection.InsertOne(context.Background(), user)
  if err != nil {
    t.Fatalf("Failed to insert user into test database: %v", err)
  }

  postsCollection := client.Database("ForumPosts").Collection("ForumPosts")

	for i := 0; i < 500; i++ {
		 // Create a test post
		 post := Post{
			PostID:   primitive.NewObjectID(),
			Username: "Albert",
			Title:    fmt.Sprintf("Albert's post %d", i),
			Body:     fmt.Sprintf("This is my post number %d", i),
			Date:     time.Now().Format("January 2, 2006"),
		}

		_, err = postsCollection.InsertOne(context.Background(), post)
		if err != nil {
			t.Fatalf("Failed to insert post into test database: %v", err)
		}
		user.Posts = append(user.Posts, post)
	}

	// Generate token for the test user
	tokenString, err := generateToken(user)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	for _, post := range user.Posts {
		// Call the updatePost function with the test data
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
	
		
		c.Request, _ = http.NewRequest(http.MethodPut, "/posts/"+post.PostID.Hex(), nil)
		c.Request.Header.Set("Authorization", "Bearer "+tokenString)

		// Set the updated post data in the request body
		updatedPost := Post{
			PostID: post.PostID,
			Title: "Albert's Updated Post",
			Body:  "This is my updated post",
		}
		updatedPostData, err := json.Marshal(updatedPost)
		if err != nil {
			t.Fatalf("Failed to marshal updated post data: %v", err)
		}
		
		body := strings.NewReader(string(updatedPostData))
		c.Request.Body = ioutil.NopCloser(body)
		c.Request.ContentLength = int64(len(updatedPostData))
		c.Request.Header.Set("Content-Type", "application/json")
	
		// Call the updatePost function
	
		updatePost(c)

		// Check that the post was updated correctly
		var updatedPostFromDB Post
		err = postsCollection.FindOne(context.Background(), bson.M{"postid": post.PostID}).Decode(&updatedPostFromDB)
		if err != nil {
    		t.Fatalf("Failed to retrieve updated post from database: %v", err)
		}

		if updatedPostFromDB.Title != updatedPost.Title {
    		t.Errorf("Failed to update post title: got %s, expected %s", updatedPostFromDB.Title, updatedPost.Title)
		}

		if updatedPostFromDB.Body != updatedPost.Body {
    		t.Errorf("Failed to update post body: got %s, expected %s", updatedPostFromDB.Body, updatedPost.Body)
		}
	}
}

// go test -timeout 30m -run ^TestAddToWatchlist$ bingebuddy.com/m
func TestAddToWatchlist(t *testing.T) {
	localMode = true
	// Create a test user
	user := User{
		Username: "Joel5",
		Password: "Aloma",
	}

	clientOptions := options.Client().ApplyURI("mongodb+srv://test:1234@cluster0.gmfsqnv.mongodb.net/test")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("UserInfo").Collection("UserInfo")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		t.Fatalf("Failed to insert user into database: %v", err)
	}

	// create a new gin context for the test
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	// create a fake JWT token
	tokenString, err := generateToken(user)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}
	cookie := &http.Cookie{
		Name:  "token",
		Value: tokenString,
	}

	for i := 0; i < 50; i++ {
		movie := Movie{
			OriginalTitle: fmt.Sprintf("Test Movie %d", i),
			Overview:      fmt.Sprintf("A test movie %d", i),
			ReleaseDate:   String("2023-04-18"),
		}

		// convert the movie data to JSON
		jsonMovies, err := json.Marshal(movie)
		if err != nil {
			t.Fatalf("Failed to marshal movie data: %v", err)
		}

		// Set the token and JSON request body in the request headers
		req, err := http.NewRequest("POST", "/"+user.Username+"/add", bytes.NewBuffer(jsonMovies))
		req.Header.Set("Content-Type", "application/json")
		req.AddCookie(cookie)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// Bind the request to the gin context
		context.Request = req

		// call the addToWatchlist function
		addToWatchlist(context)
	}

	// check if the response status code is OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	if len(w.Body.Bytes()) == 0 {
		t.Fatal("Empty response body")
	}

	// check if the response body contains the updated user with the added movies in its Watchlist
	var updatedUser User
	if err := json.NewDecoder(w.Body).Decode(&updatedUser); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	length, _ := GetWatchlistLength(user.Username)
	if length != 50 {
		t.Errorf("Expected Watchlist length of 5 but got %d", length)
	}
}

func GetWatchlistLength(username string) (int, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://test:1234@cluster0.gmfsqnv.mongodb.net/test")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return 0, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("UserInfo").Collection("UserInfo")
	filter := bson.M{"username": username}
	var user User
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return 0, fmt.Errorf("failed to find user %s: %v", username, err)
	}

	return len(user.Watchlist), nil
}

func String(v string) *string {
	return &v
}
