User Stories:

- As a user, registered or not, I want to be recommended movies that fit my preferences in genre, runtime, rating, streaming providers, and actors.
- As an unregistered user, I want the ability to register so that I can save all my watched movies and post my reviews on these movies on the forum page.
- As a registered user, I want my profile to list all my previously watched movies and be able to sort them based on rating and runtime.
- As a registered user, I want to be able to add movies that I've previously seen into the watched section of my profile at any time so that I can freely update this section.
- As a registered user, I want to access all my posts in my profile section so that I can view, edit, and remove them when I want.

Team plan:
Front-end: 
After establishing a functional workspace, we began formulating a plan to address how we should organize and tackle the different
aspects of the Login/Register, Home, and Profile pages. By the end of Sprint 1, we had planned to implement the Login/Register page's
visuals and basic functionality, such as working buttons, email/username inputs, password inputs, and a "remember me" checkbox. Additionally, we also planned on beginning development of the other pages, especially the Home page which should encompass the main functionality of our application. We all worked seperatly on different individual branches to insure that the main branch was implemented with complete features.

Back-end:
For this sprint, we initially planned to set up a basic REST API and implement the login and watchlist on the backend, as well as develop a JSON standard for communication between the frontend and the backend. 

Successful issues:
Front-end:
In Sprint 1, we implemented the main visuals and functionality of the Login/Register page, this includes all the necessary inputs and
working buttons. As for the Home page, we began to plan its implementation, not so much the visual aspects, but how the user should interact with it and how it should function as a result.

Back-end:
In Sprint 1, we were able to successfully implement general code and functionality for the Login/Register page. For the Login page, we connect to the MongoDB database and check if the username and password inputted are registered. If the user has already been created and is trying to login, our program currently outputs “Login was successful” for debugging purposes. For the Register page, we implemented code that creates a new user if it is valid and inserts that user (username, password) into the mongoDB database. 


Non-successful issues:
Front-end:
During the next sprint, we will also add the "remember me" checkbox to the Login/Register page, which is not necessary, but we believe it would increase user ease-of-access and should therefore include it. Originally, we had planned to begin working on the Home page by the end of Sprint 1, but as of currently, we are still planning its development.

Why? 
To be honest, mostly due to time managment. We first thought that we would be able to finish all those features but with the learning curve, the implementations took longer than we thought.

Back-end:
We were not able to successfully implement the backend for a "Remember Me" button. We also weren't able to add support for adding movies to the watchlist because setting up the database took much longer than we anticipated. However, we should be able to implement adding movies and the “Remember Me” button to the watchlist quickly now that the database setup is complete.
