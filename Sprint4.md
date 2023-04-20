# Sprint 4 - Work Completed

__Video Link:__

### Front-end:
  - Our main goals for the final sprint was to complete the generator and "Elite Recommendatios" section of the home page,
  along with adding all the user's movies to the "watched" section of the profile page. After some user experience testing,
  we decided that the home page would not require a "sidenav", instead, we could utilize the space better. In doing so, we 
  adjusted the layout of the profile page and removed the functionality of the "sidenav". Additionally, we decided to remove
  the option of adding movies and posts from the respective sections in the profile page and instead add this opportunity and
  functionality to the home page when a user generates or clicks on the recommended movies. By doing so, we simplify the user
  experience and allow for a cleaner and sharper aesthetic. As for the home page, we added a popup after the user generates or
  selects a movie from the recommendation section. This popup provides the user with all the necessary information of the selected movie
  and allows them to return back to the generator, add the movie to the watch list, and upon adding a movie, the option to write a post
  will open. In the case that the user is not logged in, they will still be able to access all the features of the home page - generator
  and recommendations - but will not be able to access the profile page. To ensure this, we also added user authenticated routing that only
  allows user's that are logged in to access the profile page, negating the effects of simply typing '/profile' at the end of the URL.
  
# Unit Tests (raw code found on GitHub, done with Cypress and Jasmine)
  - Watch section (render and testing to make sure it was sending and getting information from the backend)
  - Post section (render and testing to make sure it was sending and getting information from the backend)
  - NavBar section (within Profile page) (render and testing to make sure it was sending and getting information from the backend)
  - Main component that communicates with all other components (render and testing to make sure it was sending and getting information from the backend)
  - Movie Generator component (if calls functions correctly and gets data)
  - Movie Popup component (if fetches movie successfully, emmits events and get variables)
  - NavBar component in the home page (authentification emmiter and open login)

### Back-End:
 - Overhauled the random movie generator and added a generator function with filters.
 - Implemented input sanitization to prevent XSS attacks.
 - Added function to get posts with a page system.
 - Added a rating system to the posts
 - Implemented character limits and minimum lengths for all user input.
 - Wrote unit tests for the new generator formula, the movie generator with filters, and the add and remove from watchlist functions 
- Reran unit tests from previous sprints.

# Unit Tests
 1. CreateUser function with random input
 2. Login function with random input
 3. Random Movie generator speed benchmark
 4. Legacy Random Movie generator speed benchmark
 5. Random number function
 6. Similar movie feature with random IDs
 7. Post creation
 8. Post deletion
 9. Random Movie Generator with filters

# API Documentation

## New functions

 - randomMovie
 - randomMovieWithFilters
 - getPosts
 - getRandomMoviesList
 - logout
 - santizeUser, sanitizeMovieFields, sanitizePost
 - validateUser, validatePost
 
## Updated functions
 - randomMovie → changed to trueRandomMovie
 - generateRandomNumber
 - createUser
 - login
 - createPost
 - updatePost
 - updateUserInfo 
## createUser (POST)
This function uses the route "/signup" and is intended for new users. It takes a context parameter that contains the user details, binds them to a  struct, marshals that struct, sanitizes the struct to prevent XSS attacks, and creates an entry in the database corresponding to the user that stores all of the user information. This function also includes two basic checks: usernames must be unique and usernames and passwords cannot be empty strings or longer than 50 characters, and it will return a 400 error if these conditions are met. 

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
## login (POST)

The login() function receives a JSON object from the request, which holds the user’s username and password. It uses this information to search for the user in the “UserInfo” Database collection, and generates a unique JWT token for that user if the user's account is found. Before the search happens, the function `sanitizeUser` is called, which removes all HTML tags from the input to prevent XSS attacks. It will only return that token in a cookie within the response to avoid revealing sensitive information such as the user's email or password. On the other hand, if the user is not found, the function just returns an empty JSON user struct and prints in the terminal that the “username or password is incorrect.” The endpoint for this function is “/login."

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
  ## Logout (POST)
  This function prevents further authentication on the front by resetting the authentication cookie. 
  **Example response (OK)**

    {
    "message": {
    "Name": "token",
    "Value": "",
    "Path": "",
    "Domain": "",
    "Expires": "2023-04-17T08:15:36.4712638-04:00",
    "RawExpires": "",
    "MaxAge": 0,
    "Secure": false,
    "HttpOnly": true,
    "SameSite": 0,
    "Raw": "",
    "Unparsed": null
    }
    }

 ## getUserInfo (GET)
This function takes in a valid token within a cookie corresponding to an existing user and retrieves non-sensitive user information such as a user’s watchlist and posts. It is designed to be called in the background by the frontend in order to retrieve the account information after login. It works by reading the token in the cookie, decoding it with a private key for authentication, and fetching the corresponding username. This username is then used in a search query to the user information database, and sensitive information such as the email and password is always removed from the user struct before it is returned to the frontend.
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

  ## updateUserInfo (PUT)
This is a generic function that replaces a given user profile with another one, with limitations. Token authentication is implemented so that users cannot modify other users' profiles. The username of the profile is found by decoding the token and used to search the database for the correct profile. If the profile is found, the function applies the same checks as `createUser`. Usernames still cannot be duplicate  and passwords and usernames must be between  four and 50 characters.

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
    
## removeUser (DELETE)
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

## addToWatchlist (POST)
This function uses the route [username]/add and takes JSON with fields such as "Title" and "Director" that give basic information about the movie. [username] specifies the user's watchlist that the movie is added to. The username parameter is decoded from the token and used to retrieve the corresponding entry from the database through a filter, while the JSON is marshalled and stored in the database. If the specified username cannot be found or the token is invalid, the function will return a 400 error and the original profile. XSS protection is also applied to every field.
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

## removeFromWatchlist (DELETE)
This function works similarly to addToWatchlist, using the route [username]/watchlist/remove and taking in a JSON struct with the information of the movie that is to be deleted and a token corresponding to the account that it is to be deleted from. The function gets the username from this token, which is used in a filter to find the corresponding entry in the database. This feature is made fault tolerant because only the title of the movie is used in the search. The token of a match is found, the movie struct is removed from the database, while a 400 error and the original profile data is returned if a deletion can't be confirmed.
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
## randomMovie (GET)
This function takes no input and is called on the route `/generate`. It works by making a **GET** request to a random page out of the first 150 of the TMDB API's top rated movies. This page is unmarshaled into an object, which contains an array of movie objects containing information about the movie. A random index of this array is usually returned. However, if the movie contains adult content or does not contain English, the page number and page index are randomly generated and a new request is made until the result is suitable to be returned to the frontend. This check is performed by the function `filterMovies`.
## getRandomMoviesList (GET)
This function is designed to return a list of 8 random movie objects for a carousel on the front page. It simply calls randomMovie until eight unique movies are produced and returns the list to the frontend.

## randomMovieWithFilters (POST)
This function is a POST request because it takes in the filter parameters in the request body. Those parameters are: actors, genres, maximum runtime, minimum rating, and streaming providers.  The actors are represented by an array of strings, the genres and providers are represented by an array of integers corresponding to the TMDB API's internal ID (see key below), the maximum runtime is an integer in minutes, and the minimum rating is a float between 1 and 10. After it has unmarshaled the filters, the function then iterates through the actors array, where it searches for the actors' IDs by making **GET** requests to TMDB's API. If there are no results, the API simply skips over that array item.  Then, the function uses string manipulation to assemble a  **GET**  request to TMDB's discover pathway, which will do most of the filtering for BingeBuddy. Results will have every specified actor and genre, and will be on at least one of the specified streaming platforms. Sometimes, it will ignore parameters if there are no results. This necessitates the use of a "visited" system to verify that there are no valid results to a given query. Page numbers that the generator has visited are stored, and the total number of movies visited is also stored. The code that extracts and returns the random movie is enclosed in a loop and will continue to run until either an appropriate result is found or the total number of movies visited is equal to the total number of results. If there are no results, the function returns the error message `"No results"`.

**Filter Keys**
|Genre|	ID  |
|--|--|
|Action  |  28|				    
|Adventure|12  |
|Animation|16|
| Comedy |  35 |
| Crime | 80 |
|Documentary|99|
| Drama | 18 |
| Family |  10751|
|Fantasy|14|
|  History| 13 |
|Horror| 27 |
|Music|10402|
| Mystery |9648  |
|Romance| 10749 |
|Science Fiction|878|
|  TV Movie| 10770 |
|Thriller|  53|
|War|10752|
| Western | 37 |




|Service| ID|
|--|--|
| Netflix | 8
Amazon  Prime  Video |9
|YouTube  Premium|188  |
|Hulu|15|
|  Disney+| 337 |
|HBO  Max| 384 |
|Peacock  Premium|387|
| Crunchyroll | 283 |
|Apple TV Plus|350 |

**Example Request**

    {
    "actors": ["Michelle Yeoh", "ke huy quan"],
    "max_runtime": 180,
    "genres": [878, 12],
    "min_rating": 7.2,
    "providers": [15, 9]
    }
**Response**

    {
    "adult": false,
    "backdrop_path": "/ss0Os3uWJfQAENILHZUdX8Tt1OC.jpg",
    "genre_ids": [28,12,878],
    "id": 545611,
    "original_language": "en",
    "original_title": "Everything Everywhere All at Once",
    "overview": "An aging Chinese immigrant is swept up in an insane adventure, where she alone can save what's important to her by connecting with the lives she could have led in other universes.",
    "popularity": 190.978,
    "poster_path": "/w3LxiVYdWWRvEVdn5RYq6jIqkb1.jpg",
    "release_date": "2022-03-24",
    "title": "Everything Everywhere All at Once",
    "vote_average": 7.9,
    "vote_count": 4414
    }
**Example Response (Error)**

    {
    "error": "No results"
    }



## trueRandomMovie (GET)
This function takes no input. It makes a request to the TMDB API and returns the raw JSON data of the randomly generated movie to the frontend. Since the TMBD API does not have a native way to return random movies, the random movie generator works by repeatedly making GET requests to the API with randomly generated IDs in the valid range until it receives a valid (200) response. This will return the data of a movie that corresponds to that ID in the API's database. However, there is a filter that removes adult content and ensures that at least one of the movie's original languages is English. The filter works by marshalling the API's raw JSON response into a struct and checking whether it contains adult content and is in English. If these conditions are met, the process restarts. For a more detailed explanation of how a random ID is generated, see generateRandomNumber below.

## generateRandomNumber 

 This function was originally designed to unit test the formula used in the random movie generator without the overhead of the regular generator function, but it is now used to generate any random integers within the program. The largest and smallest desired values are passed into the function and it returns an integer between and inclusive of those values. Using the time in nanoseconds as a seed, the formula works by first generating a number between zero and the difference between the largest and smallest values by multiplying that difference by a random decimal between zero and one using the current time as a seed. The smallest value, which serves as a lower bound, and 0.5 are added to this number. 0.5 is added because it makes rounding up possible during integer conversion since the normal integer conversion only truncates, which is equivalent to always rounding down.

## getSimilarMovie (GET)
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

## getPosts (POST)
This function is designed to retrieve posts relating to a specific movie for the frontend. It uses the route `posts/[movie_id]/[page number]`. The posts are ordered from most to least recent and there are 50 posts per page. If a page number is out of bounds, the function simply returns the oldest 50 posts, or all of them if there are fewer than 50. If the page number is invalid, the function simply returns the first page of posts.
**Example Response:**

    [
    {
    "id": "000000000000000000000000",
    "movie_id": "1234",
    "username": "",
    "title": "my 11th post",
    "body": "dsffdasdfasfaffdasdgasdasdgasdgasdgasdgasdgasdgsdgasgddfsafsdfdsafsdgasdgasggdsagsdg",
    "date": ""
    },
    {
    "id": "000000000000000000000000",
    "movie_id": "1234",
    "username": "",
    "title": "my 10th post",
    "body": "dsffdasdfasfaffdasdgasdasdgasdgasdgasdgasdgasdgsdgasgddfsafsdfdsafsdgasdgasggdsagsdg",
    "date": ""
    },... //additional response information omitted for brevity

## createPost (POST)
The createPost function creates a new post and stores it into the MongoDB database that the backend is connected to. Posts consist of a title and a body. A title has a character limit of 200 characters, while a body has a limit of 2500. The post that is created belongs to the user who created it and it updates the overall post array of the user, showcasing all the different posts that the logged in user has posted. This function gets the token from the user, which is how it is able to authenticate the user to create a post. If the token is invalid or cannot be found, it basically throws an error stating that the user is unauthorized. With the use of claims and the token, it grabs the username of the logged in user. It also creates a new post struct and takes in the JSON body request, which includes the post's id, title, body, and date published. It is able to authenticate the user to create this post by requiring the valid access token to be passed in the “Authorization” header with the format “Bearer [token].” The fields of the post struct are sanitized to prevent XSS attacks. Once this post has been created, it gets inserted into the “ForumPosts” database collection in the MongoDB database and it also gets inserted into the user’s post array. The endpoint for the createPost() function is “/posts and the HTTP method is POST 

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
  

## updatePost (PUT)
This function lets the authenticated logged in user to update a post that they had already created previously. The updated post is sanitized to prevent XSS attacks, and the same character limits apply to it. It first checks if the user is authorized to make the update by verifying the JWT token. If the user is authorized, it then checks if the Post ID is valid and if the post exists. If the post does not exist or if the Post ID is invalid, it will throw an error. For the user to actually update their post, they have to be the creator of that post. So, in order to prevent a user from updating other users’ posts, this function then checks if the user is the creator of the post that is being updated. If the user is the creator, the function then updates the post in the database, and also updates the user's post in the user's array. In order to authorize the user, it takes the JWT token with the Bearer prefix in the authorization header to get the user information and then see if this user is the creator of the post.
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

## deletePost (DELETE)
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
## sanitizeUser, sanitizeMovieFields, sanitizePost
These functions utilize the bluemonday library to perform input sanitization. They are only called within other functions and take in user, movie, and post structs, respectively. Internally, each function applies bluemonday's strict policy to the string fields of each object, which removes any HTML tags. The resulting output is checked  afterward to ensure that it complies with character limits in the validate functions (see below)
## validateUser, validateMovie, validatePost
These functions provide an efficient way to both check the whether data in user, movie, and post objects comply with character limits and minimums inside other functions such as createPost and return a helpful and descriptive error message that can be sent back to the frontend and displayed. Similarly to the sanitization functions above, validateUser, validateMovie, and validatePost all take in a pointer to an object. However, these functions all return a boolean and a string. The boolean describes whether the data in the structure is valid. If it isn't, the string will contain an error message, which will be sent back to the frontend in an `error ` field. Otherwise, the function will operate as expected.

**Possible error messages**
 1. Post/Comment/Title/Username/Password cannot be blank
 2. Post/Comment/Title/Username/Password/Movie field cannot exceed character limit 
 3. Username/Password must be at least four characters
