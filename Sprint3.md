Work Completed

Front-End:
One of our main focuses this sprint was to successfully authenticate the connection between the user and the back-end when login in and registering for
an account. As of the end of the sprint, we are now able to communicate with the back-end, sending user registration information and receiving
user data to validate the login/register process. Additionally, we worked on the profile page, fixing our issues with the Angular material side
navigation bar, along with further implementing the functionality of the "my watched" and "my post" section. In each of the sections, we are now
able to successfully add movie objects and post objects, which we will then pass on to the back-end for each user. In the
upcoming sprint, we plan on fully incorporating the back-end with the profile section and using the retrieved information to fill out the necessary
user information in the profile section. Additionally, we plan on adding the remove functionality to the respective sections of the profile page.

Video link: https://youtu.be/fGRTOwWzgdM

# Unit Tests (raw code found on GitHub in each respective component (Cypress))

Home Page - testing login and logout, testing registration, testing navigation bar interaction, testing login/register component
Profile Page - "my watched section, add" section input validation for (title, actor, genre, rating, runtime), testing sidenav, testing watched and post sections


Back-End:
 - Fully automated updating generator parameters on startup by parsing daily data dumps from the TMDB API.
 - Implemented a similar movies feature.
 - Added token authentication to logins, adding and removing items from user watchlists and editing user profiles.
 - Added ability to toggle between online and local databases for unit testing
 - Added generic function to update user profiles
 - Updated the formula used to generate random movies.
 - Unit tested the generator formula, similar movies feature, and the creation and deletion of posts.

# Unit Tests

 1. CreateUser function with random input
 2. Login function with random input
 3. Random Movie generator speed benchmark
 4. Random number function
 5. Similar movie feature with random IDs
 6. Post creation
 7. Post deletion
 8. Editing posts

# API Documentation

## New functions

 - updateGeneratorParameters()
 - getSimilarMovies(context *gin.Context)
 - generateRandomNumber(smallest float64 largest float64) int
 - getUserInfo(context *gin.Context)
 - updatePost(context *gin.Context)
 
## Updated functions
 - randomMovie(context *gin.Context)
 - login()
 - createPost()
 - deletePost()
 - addToWatchlist()
 - removeFromWatchlist()
 - updateUserInfo()
 - removeUser()
## Removed functions
 - ScanValidIDs() - Replaced by UpdateGeneratorParameters

 

## login

  
The login() function receives a JSON object from the request, which holds the user’s username and password. It uses this information to search for the user in the “UserInfo” Database collection, and generates a unique JWT token for that user if the user is found. It will only return that token in the response to avoid revealing sensitive information such as the user's email or password. On the other hand, if the user is not found, the function just returns an empty JSON user struct and prints in the terminal that the “username or password is incorrect.” The endpoint for this function is “/login” and the HTTP method is POST 

**Example request:** 

    { 
    “username”: “test”, 
    “password”: “1234 
    } 

**Example Response (OK):**

     { 
     "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRob3JpemVkIjp0cnVlLCJleHBpcmF0aW9uRGF0ZSI6MTY3NzgwODUyMCwidXNlcm5hbWUiOiJoZXJlIn0.vAKlLj3RXmMv6flsAT4pRwEii6a62XcxkZQzsD1LMSg" 
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

 ## getUserInfo
This function takes in a valid token corresponding to an existing user and retrieves non-sensitive user information such as a user’s watchlist and posts. It is designed to be called in the background by the frontend in order to retrieve the account information after login. It works by reading the token in the header, decoding it with a private key to authenticate it, and fetching the corresponding username. This username is then used in a search query to the user information database, and sensitive information such as the email and password is always removed from the user struct before it is returned to the frontend.
**Example request:** 

    { 
    "token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRob3JpemVkIjp0cnVlLCJleHBpcmF0aW9uRGF0ZSI6MTY3NzgwODUyMCwidXNlcm5hbWUiOiJoZXJlIn0.vAKlLj3RXmMv6flsAT4pRwEii6a62XcxkZQzsD1LMSg"
    } 
 **Example Response (OK)**

      { 
      "username": "test1234",
       "password": "", 
       "email": "", 
       "watchlist": [], 
       "posts": [], 
       "genres": [], 
       "rating": 7.5,
        "subscriptions": {"Netflix", "Prime Video", "Hulu"}
      }
 **Example Response (Invalid Token)**

    {
    "error": "Unauthorized User"
    }

## connectToDB
This function is called in almost every function in the API and allows them to access the mongoDB database. It returns a Client object that can be turned into a database object by calling the functions Database and Collection. These functions take in a string that specifies which database and collection in the account that should be accessed. We can only interact with the database by calling CRUD functions such as FindOne on this database object. The function loads the URI based on the boolean localMode, which is set in the main method and can be set during unit testing. If localMode is true, the database that this function connects to is the MongoDB database on the machine. If it is false, the program will connect to our online database.

## createUser
This function uses the route "/signup" and is intended for new users. It takes a context parameter that contains the user details, binds them to a  struct, marshals that struct, and creates an entry in the database corresponding to the user that stores all of the user information. This function also includes two basic checks: usernames must be unique and usernames and passwords cannot be empty strings, and it will return a 400 error if these conditions are met. 

**Example request:**

    { 
    "token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRob3JpemVkIjp0cnVlLCJleHBpcmF0aW9uRGF0ZSI6MTY3NzgwODUyMCwidXNlcm5hbWUiOiJoZXJlIn0.vAKlLj3RXmMv6flsAT4pRwEii6a62XcxkZQzsD1LMSg"
    } 
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
   
   **Example Response (Invalid Token)**

    {
    "error": "Unauthorized User"
    }
  ## updateUserInfo
This is a generic function that replaces a given user profile with another one, with limitations. Token authentication is implemented so that users cannot modify other users' profiles. The username of the profile is found by decoding the token and  used to search the database for the correct profile. If the profile is found, the function checks whether the new username is already taken, unless it is identical to the previous username and replaces the user profile in the database and returns it to the frontend. If the username is duplicate, the backend will return an error along with the unchanged current profile.

**Example request:**
	

    { 
    "token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRob3JpemVkIjp0cnVlLCJleHBpcmF0aW9uRGF0ZSI6MTY3NzgwODUyMCwidXNlcm5hbWUiOiJoZXJlIn0.vAKlLj3RXmMv6flsAT4pRwEii6a62XcxkZQzsD1LMSg"
    } 
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
**Example response (Invalid Token)**

    {
    "error": "Unauthorized User"
    }
**Example response (Duplicate username)**

    {
    "username": "test5",
    "password": "1235",
    "email": "",
    "watchlist": null,
    "genres": null,
    "rating": 0,
    "providers": null
    }
    
## removeUser
This function removes the corresponding user from the database when given a valid token. It works by parsing the token and finding the associated username. If the token is invalid, an error is returned. If the user exists, it is deleted, and the profile is returned to the frontend.
**Example Request**

    { 
    "token" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRob3JpemVkIjp0cnVlLCJleHBpcmF0aW9uRGF0ZSI6MTY3NzgwODUyMCwidXNlcm5hbWUiOiJoZXJlIn0.vAKlLj3RXmMv6flsAT4pRwEii6a62XcxkZQzsD1LMSg"
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
  **Example Response (Invalid Token)**

    {
    "error": "Unauthorized User"
    }

## addToWatchlist
This function uses the route [username]/add and takes JSON with fields such as "Title" and "Director" that give basic information about the movie. [username] specifies the user's watchlist that the movie is added to. The username parameter is decoded from the token and used to retrieve the corresponding entry from the database through a filter, while the JSON is marshalled and stored in the database. If the specified username cannot be found or the token is invalid, the function will return a 400 error and the original profile.
**Example Request:**
	
 

    {
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRob3JpemVkIjp0cnVlLCJleHBpcmF0aW9uRGF0ZSI6MTY3NzgwODUyMCwidXNlcm5hbWUiOiJoZXJlIn0.vAKlLj3RXmMv6flsAT4pRwEii6a62XcxkZQzsD1LMSg" 
    } 
     
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
   **Example Response (Invalid Token)**

    {
    "error": "Unauthorized User"
    }

## removeFromWatchlist
This function works similarly to addToWatchlist, using the route [username]/watchlist/remove and taking in a JSON struct with the information of the movie that is to be deleted. The username is used in a filter to find the corresponding entry in the database, and this feature is made fault tolerant because only the title of the movie is used in the search. The token If a match is found, the movie struct is removed from the database, while a 400 error and the original profile data is returned if a deletion can't be confirmed.
**Example Request:**

    { 
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdXRob3JpemVkIjp0cnVlLCJleHBpcmF0aW9uRGF0ZSI6MTY3NzgwODUyMCwidXNlcm5hbWUiOiJoZXJlIn0.vAKlLj3RXmMv6flsAT4pRwEii6a62XcxkZQzsD1LMSg" 
    } 
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
   **Example Response (Invalid Token)**

    {
    "error": "Unauthorized User"
    }

## updateGeneratorParameters
This function finds the largest and smallest ID value, which are then used as bounds in the random number generator in randomMovie, by downloading a data dump from the TMDB API that contains every valid movie ID in JSON and storing the largest and smallest valid IDs in our database. The correct URL is automatically created using the current date, and this function is called automatically on startup. First, it checks when the parameters were last updated. If it has been less than 24 hours since it was last updated, the function only prints a message stating that the database does not need to be updated. TMDB’s data dumps are published daily, so it would be pointless to scan for updates more than once every 24 hours. If the database does need to be updated, the function first makes a GET request to the API and downloads the most recent data dump. The data is written to a temporary file, which is decompressed, and then read. Afterwards, all temporary files are deleted and the new smallest and largest ID values are stored in the database. These ID values will be read in the main function and utilized by the random movie generator.

## randomMovie
This function takes no input and is called on the route "/generate." It makes a request to the TMDB API and returns the raw JSON data of the randomly generated movie to the frontend. Since the TMBD API does not have a native way to return random movies, the random movie generator works by repeatedly making GET requests to the API with randomly generated IDs in the valid range until it receives a valid (200) response. This will return the data of a movie that corresponds to that ID in the API's database. However, there is a filter that removes adult content and ensures that at least one of the movie's original languages is English. The filter works by marshalling the API's raw JSON response into a struct and checking whether it contains adult content and is in English. If these conditions are met, the process restarts. For a more detailed explanation of how a random ID is generated, see generateRandomNumber below.

## generateRandomNumber

 This function is designed to unit test the formula used in the random movie generator without the overhead of the regular generator function. It is identical to the formula used in the random movie generator. The largest and smallest desired values are passed into the function and it returns an integer between and inclusive of those values. The formula works by first generating a number between zero and the difference between the largest and smallest values by multiplying that difference by a random decimal between zero and one using the current time as a seed. The smallest value, which serves as a lower bound, and 0.5 are added to this number. 0.5 is added because it makes rounding up possible during integer conversion since the normal integer conversion only truncates, which is equivalent to always rounding down.

## getSimilarMovie
When given an ID as a parameter, this function returns a JSON object containing the basic information for a group of similar movies to the movie corresponding to the ID of variable size. The function works by sending a GET request to the TMDB API containing the movie ID. The API then returns the first page of a list of similar movie objects that contain basic information about those movies, such as the title and poster link. This information is unmarshaled into a results object that contains a list of the backend’s movie objects and sent to the frontend.
**Example Response**

    "results": [
    {
    "adult": false,
    "id": 102482,
    "original_language": "es",
    "original_title": "La vida nocturna",
    "overview": "Stan lies to his wife about going to a nightclub with Ollie but Mrs. Laurel overhears the plot and outsmarts them both.",
    "popularity": 0.718,
    "poster_path": "/umUDjnY6PTHkKgRbE33KOqEgipd.jpg",
    "release_date": "1930-04-19",
    "title": "The Night Life",
    "vote_average": 8,
    "vote_count": 2
    }, ...//additional responses hidden for brevity

## createPost
The createPost function creates a new post and stores it into the MongoDB database that the backend is connected to. The post that is created belongs to the user who created it and it updates the overall post array of the user, showcasing all the different posts that the logged in user has posted. This function gets the token from the user, which is how it is able to authenticate the user to create a post. If the token is invalid or cannot be found, it basically throws an error stating that the user is unauthorized. With the use of claims and the token, it grabs the username of the logged in user. It also creates a new post struct and takes in the JSON body request, which includes the post's id, title, body, and date published. It is able to authenticate the user to create this post by requiring the valid access token to be passed in the “Authorization” header with the format “Bearer [token].” Once this post has been created, it gets inserted into the “ForumPosts” database collection in the MongoDB database and it also gets inserted into the user’s post array. The endpoint for the createPost() function is “/posts and the HTTP method is POST 

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

**Example Response (Invalid Token)**

     {
     "error": "Unauthorized User" 
     } 

**Example Response (Internal Service Error)**

     { 
     "error": "Failed to create post"
     }
  

## updatePost
This function lets the authenticated logged in user to update a post that they had already created previously. It first checks if the user is authorized to make the update by verifying the JWT token. If the user is authorized, it then checks if the Post ID is valid and if the post exists. If the post does not exist or if the Post ID is invalid, it will throw an error. For the user to actually update their post, they have to be the creator of that post. So, in order to prevent a user from updating other users’ posts, this function then checks if the user is the creator of the post that is being updated. If the user is the creator, the function then updates the post in the database, and also updates the user's post in the user's array. In order to authorize the user, it takes the JWT token with the Bearer prefix in the authorization header to get the user information and then see if this user is the creator of the post.
**Example Request**

    {
    “title”: “First Post, “body”: “This is my first post!”
    }
  
**Example Response (OK)**

    {
    “message”: “Post updated successfully”
    }

**Example Error Message**
	
    }
    “error”: “You cannot update a post that is not yours” 
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

## generateToken

This function generates a random token for the logged in user after being called in the login() function. The token is generated by using claims, signing method, and the jwt.MapClaims. It also uses a secret key that is required to use whenever you want to use the token to authorize a user. It returns the unique token. There is no response or http request method for this function since this function is just called by the login function to assign a unique token to the logged in user.
