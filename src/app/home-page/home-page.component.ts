import { Component, OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { MovieGeneratorService } from '../services/movie-generator.service';
<<<<<<< Updated upstream
import { LoginRegisterService } from '../services/login-register.service';
=======
>>>>>>> Stashed changes
import { Movie } from '../user-auth/user';


@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  isLoginOpen = false
  isAuthenticated
  username = ""
  message = 'Home Page'
  movieList: Movie[] = []
  userWatchlist: Movie[] = []

  constructor(
<<<<<<< Updated upstream
    private movieGeneratorService: MovieGeneratorService,
    private loginService: LoginRegisterService
  ){}
=======
    private movieGeneratorService: MovieGeneratorService
  ){
  }
>>>>>>> Stashed changes

  ngOnInit(): void {
    this.getMovie()

    Emmiters.userData.subscribe(
      {
        next: (username : string) => {
          this.username = username
<<<<<<< Updated upstream
          this.loginService.setUsername(this.username);
=======
>>>>>>> Stashed changes
        }
      }
    )

    Emmiters.authEmmiter.subscribe(
      {
        next: (auth : boolean) => {
          this.isAuthenticated = auth
<<<<<<< Updated upstream
          this.loginService.setUserAuth(this.isAuthenticated);
=======
>>>>>>> Stashed changes
          this.message = 'Hey '
        }
      }
    )
  }

  getMovie() {
    this.movieGeneratorService.getRandomMovie().subscribe(
      {
        next: (resp) => {
          this.movieList = resp
        },
        error: (err) => {
          console.log('ERROR GENERATING MOVIES')
        }
      }
    )
<<<<<<< Updated upstream
=======
  }
  
  openMoviePopup(movie: Movie){
    Emmiters.generatedMovie.emit(movie)

    for(let i = 0; i < Emmiters.watchList.length; i++){
      if(Emmiters.watchList.at(i)?.id == movie.id){
        Emmiters.isMovieWatched.emit(true);
        console.log('movie was watched');
      }
    }

    Emmiters.isPopupOpen.emit(true);
>>>>>>> Stashed changes
  }

  setUserAuth(auth: boolean) {
    this.isAuthenticated = auth;
  }
  openMoviePopup(movie: Movie){
    Emmiters.generatedMovie.emit(movie)

    for(let i = 0; i < Emmiters.watchList.length; i++){
      if(Emmiters.watchList.at(i)?.id == movie.id){
        Emmiters.isMovieWatched.emit(true);
        console.log('movie was watched');
      }
    }

    Emmiters.isPopupOpen.emit(true);
  }
}