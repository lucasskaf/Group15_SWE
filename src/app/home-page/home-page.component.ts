import { Component, OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { MovieGeneratorService } from '../services/movie-generator.service';
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
    private movieGeneratorService: MovieGeneratorService
  ){
  }

  ngOnInit(): void {
    this.getMovie()

    Emmiters.userData.subscribe(
      {
        next: (username : string) => {
          this.username = username
        }
      }
    )

    Emmiters.authEmmiter.subscribe(
      {
        next: (auth : boolean) => {
          this.isAuthenticated = auth
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
