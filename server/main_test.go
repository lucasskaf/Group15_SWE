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
	largest := 17.0
	output := 0
	for output != int(smallest) {
		output = generateRandomNumber(smallest, largest)
		fmt.Println(output)
		//test automatically fails if out of bounds output is produced
		if output < int(smallest) || output > int(largest) {
			t.Fail()
		}
	}
	for output != int(largest) {
		output = generateRandomNumber(smallest, largest)
		fmt.Println(output)
		if output < int(smallest) || output > int(largest) {
			t.Fail()
		}
	}
}

// Command: go test -timeout 10m -run ^TestGetSimilarMovies$ bingebuddy.com/m
func TestGetSimilarMovies(t *testing.T) {
	for i := 0; i < 500; i++ {
		id := generateRandomNumber(2.0, 1000.0)
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
	generatePosts(1000)
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
	//for _, p := range user.Posts {
		delReq := httptest.NewRequest("DELETE", "http://localhost:8080", nil)
		context.Params = []gin.Param{
			{
				//Key:   "postID",
				//Value: p.PostID.Hex(),
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
