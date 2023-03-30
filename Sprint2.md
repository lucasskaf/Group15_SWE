# Demo Video Presentation
https://youtu.be/7oi2elZxAWM

# Sprint Summary - Frontend
The main focus of this sprint was to establish a connection between the front-end and back-end. As of the end of sprint 2, our project's
front end consists of a main home page and a profile page. The main home page contains two buttons, a home button, which currently offers 
no functionality, and a login button, which when pressed, will prompt the user with our login/register form. In this sprint, we were able to
connect the front-end and the back-end storing the user information after registering, which is stored in our MongoDB server. After the user
does this, the login/register form automatically closes. Additionally, we also began implementation of the profile page, and as of the end of this
sprint, it contains an Angular material navbar and toolbar, which constitutes for the page's top and side navigation bar. There are also two
interactable buttons, one besides the "Binge Buddy" text which toggles the side navigation bar and the "home" button at the far end of the top
navigation bar. Placeholders for the future "My Watched" and "My Posts" sections were also created, but currently hold no functionality.

# Sprint Summary - Backend

 - Updated login function with authentication tokens.
 - Added posting and deleting to profiles with authentication tokens.
 - Added adding and deleting movies to watchlist and added generic functions to update profile information and delete profiles.
 - Implemented a working random movie generator
 - Implemented an experimental function to parse data dumps from the TMDB API to shape input for random movie generator
 - Unit tested profile creation, logins, and the movie generator

# Unit Tests - Frontend

1. Cypress: Testing dropdown button in profile page
2. Login open modal functionality
3. Modal closure functionality

# Unit Tests - Backend

 1. CreateUser function with random input
 2. Login function with random input
 3. Random Movie generator speed benchmark

# API Documentation

## login
  
The login() function receives a JSON object from the request, which holds the user’s username and password. It uses this information to search for the user in the “UserInfo” Database collection, and generates a unique JWT token for that user if the user is found. It will also return that token in the response, along with all the credentials of the user, such as username, password, email, posts, watchlist, genre, rating, and subscriptions. On the other hand, if the user is not found, the function just returns an empty JSON user struct and prints in the terminal that the “username or password is incorrect.” The endpoint for this function is “/login” and the HTTP method is POST 

**Example request:** 

    { 
    “username”: “test”, 
    “password”: “1234 
    } 

**Example Response (OK):**

     { "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRob3JpemVkIjp0cnVlLCJleHBpcmF0aW9uRGF0ZSI6MTY3NzgwODUyMCwidXNlcm5hbWUiOiJoZXJlIn0.vAKlLj3RXmMv6flsAT4pRwEii6a62XcxkZQzsD1LMSg" 
     } 
     {
      "username": "here", 
     "password": "nothere", 
     "email": ""
     , "watchlist": null,
      "posts": [], 
      "genres": null, 
      "rating": 0, 
      "subscriptions": null 
      } 

  **Example Response (Error)**
 

     { 
      "username": "",
       "password": "", 
       "email": "", 
       "watchlist": null, 
       "posts": null, 
       "genres": null, 
       "rating": 0,
        "subscriptions": null
      }

 
## connectToDB
This function is called in almost every function in the API and allows them to access the mongoDB database. It returns a Client object that can be turned into a database object by calling the functions Database and Collection. These functions take in a string that specifies which database and collection in the account that should be accessed. We can only interact with the database by calling CRUD functions such as FindOne on this database object. " The function first attempts to load the URI, or Uniform Resource Identifier, which acts as a link to the database, from the go.env file. If this fails, the function attempts to get the URI from the operating system's environment. Finally, a URI is hardcoded in case the previous steps fail. 

## createUser
This function uses the route "/signup" and is intended for new users. It takes a context parameter that contains the user details, binds them to a  struct, marshals that struct, and creates an entry in the database corresponding to the user that stores all of the user information. This function also includes two basic checks: usernames must be unique and usernames and passwords cannot be empty strings, and it will return a 400 error if these conditions are met. 

**Example request:**

    {
    "username": "test6",
    "password": "1234",
    "email": "",
    "watchlist": null,
    "genres": null,
    "rating": 0,
    "providers": null
    }

**Example response (OK)** 

    {
    "username": "test6",
    "password": "1234",
    "email": "",
    "watchlist": null,
    "genres": null,
    "rating": 0,
    "providers": null
    }

**Example response (Error)**

    {
    "username": "test6",
    "password": "1234",
    "email": "",
    "watchlist": null,
    "genres": null,
    "rating": 0,
    "providers": null
    }

## addToWatchlist
This function uses the route [username]/add and takes JSON with fields such as "Title" and "Director" that give basic information about the movie. [username] specifies the user's watchlist that the movie is added to. The username parameter is stored and used to retrieve the corresponding entry from the database through a filter, while the JSON is marshalled and stored in the database. If the specified username cannot be found, the function will return a 400 error and the original profile.
**Example Request:**

    {
    "title": "Shrek",
    "director": "Andrew Adamson and Vicky Jenson"
    }

**Example Response (OK):**

    {
    "username": "test2",
    "password": "1234",
    "email": "",
    "watchlist": 
    [
    {
    "title": "Shrek",
    "director": "Andrew Adamson and Vicky Jenson",
    "imglink": "",
    "runtime": 0,
    "avgrating": 0,
    "providers": null,
    "databaseid": 0
    }
    ],
    "genres": null,
    "rating": 0,
    "subscriptions": null
    }

**Example Response (Error):**

    {
    "username": "test2",
    "password": "1234",
    "email": "",
    "watchlist": null,
    "genres": null,
    "rating": 0,
    "subscriptions": null
    }

## removeFromWatchlist
This function works similarly to addToWatchlist, using the route [username]/watchlist/remove and taking in a JSON struct with the information of the movie that is to be deleted. The username is used in a filter to find the corresponding entry in the database, and this feature is made fault tolerant because only the title of the movie is used in the search. If a match is found, the movie struct is removed from the database, while a 400 error and the original profile data is returned if a deletion can't be confirmed.
**Example Request:**

    {
    "title": "Shrek",
    "director": "Andrew Adamson and Vicky Jenson"
    }

**Example Response (OK)**

    {
    "username": "test2",
    "password": "1234",
    "email": "",
    "watchlist": null,
    "genres": null,
    "rating": 0,
    "subscriptions": null
    }

**Example Response (Error)**

    {
    "username": "test2",
    "password": "1234",
    "email": "",
    "watchlist": 
    [
    {
    "title": "Shrek",
    "director": "Andrew Adamson and Vicky Jenson",
    "imglink": "",
    "runtime": 0,
    "avgrating": 0,
    "providers": null,
    "databaseid": 0
    }
    ],
    "genres": null,
    "rating": 0,
    "subscriptions": null
    }

## scanValidIDs
This function parses a JSON file containing every movie in the TMBD API, our movie API and returns the largest and smallest ID contained in the API database. These values are used to set the parameters of our random movie generator. This JSON file can be downloaded directly from the browser with the link "http://files.tmdb.org/p/exports/movie_ids_[day]_[month]\_[year].json.gz" 
The user must manually decompress the file, but we plan to automate this entire process in the next sprint. The function then iterates through the file line by line and unmarshals each JSON struct into the special struct parseStruct. Then, the ID contained in the parse struct is compared to variables that will contain the largest and smallest ids in the database, respectively. The smallest variable is set to the maximum integer value by default and is replaced by the ID of a specific movie if the ID is less than the smallest variable, while the largest variable  is set to zero (there are no negative IDs) and replaced when the movie ID is larger.

## randomMovie
This function takes no input and is called on the route "/generate." It makes a request to the TMDB API and returns the raw JSON data of the randomly generated movie to the frontend. Since the TMBD API does not have a native way to return random movies, the random movie generator works by repeatedly making GET requests to the API with randomly generated IDs in the valid range until it receives a valid (200) response. This will return the data of a movie that corresponds to that ID in the API's database. However, there is a filter that removes adult content and ensures that at least one of the movie's original languages is English. The filter works by marshalling the API's raw JSON response into a struct and checking whether it contains adult content and is in English. If these conditions are met, the process restarts. 


## createPost
The createPost() function creates a new post and stores it into the MongoDB database that the backend is connected to. The post that is created belongs to the users who created it and it updates the overall post array of the user, showcasing all the different posts that the logged in user has posted. This function gets the token from the user, which is how it is able to authenticate the user to create a post. If the token is invalid or cannot be found, it basically throws an error stating that the user is unauthorized. With the use of claims and the token, it grabs the username of the logged in user. It also creates a new post struct and takes in the JSON body request, which includes the post's id, title, body, and date published. It is able to authenticate the user to create this post by requiring the valid access token to be passed in the “Authorization” header with the format “Bearer [token].” Once this post has been created, it gets inserted into the “ForumPosts” database collection in the MongoDB database and it also gets inserted into the user’s post array. The endpoint for the createPost() function is “/posts and the HTTP method is POST 

**Example request:**

     { 
     “title”: “First Post”, 
     “body”: “This is my first post” 
     } 

 **Example response (OK):**

     {
     "id":"64001c98e662e298c4ee260c", 
     "title":"First Post", 
     "body":"This is my first post", 
     "date":"March 1, 2023"
     } 

**Example Response (Failed Authentication)**

     {
     "error": "Unauthorized User" 
     } 

**Example Response (Internal Service Error)**

     { 
     "error": "Failed to create post"
     }

## deletePost
Similar to the createPost() function, but this function just deletes the function. Besides deleting instead of adding, it follows the same functionality as the createPost() function. It grabs the token for the “Authorization” header to authenticate the user and it uses claims to get the logged in user's username. If the token is invalid or cannot be found, it will return an error stating that the user is unauthorized. It also grabs the post ID from the post that the user wants to delete, and checks for in the mongoDB database. If the post is found and it belongs to the logged in user, it will delete it from the “ForumPosts” database collection and remove it from the user’s post array. If the id is not found or the post does not belong to the user, it will throw an error message. If the post is deleted successfully, it gives a message as a response stating that the post was deleted successfully. Something very important to mention is that when testing in Postman, it is required to put the users token in the authorization header and the post id in the URL in order for it to function successfully. The endpoint for the deletePost() function is “/posts/:postID and the HTTP method is DELETE 
**Example response (OK)** 
{ 
“message”: “Post deleted successfully” 
}
**Example response (Error)**
 { 
"error": "You cannot delete a post that is not yours!"
  }

## generateToken()

This function is simple, it generates a random token for the logged in user after being called in the login() function. The token is generated by using claims, signing method, and the jwt.MapClaims. It also uses a secret key that is required to use whenever you want to use the token to authorize a user. It returns the unique token. There is no response or http request method for this function since this function is just called by the login function to assign a unique token to the logged in user.


