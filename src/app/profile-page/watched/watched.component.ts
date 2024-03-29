import { Component, EventEmitter, OnInit, Output, Input } from '@angular/core';
import { Emmiters } from 'src/app/emitters/emmiters';
import { Movie } from 'src/app/user-auth/user';
import { LoginRegisterService } from 'src/app/services/login-register.service';
import { MovieGeneratorService } from 'src/app/services/movie-generator.service';

@Component({
  selector: 'app-watched',
  templateUrl: './watched.component.html',
  styleUrls: ['./watched.component.css']
})
export class WatchedComponent implements OnInit {
  @Output() openMovieAddPopup = new EventEmitter<boolean>();
  @Input() addMoviePopupStatus: boolean = false;

  movies: Movie[] = [];
  username: string = "";

  constructor(private loginService: LoginRegisterService, private movieService: MovieGeneratorService) { }

  ngOnInit() {
    this.movies = Emmiters.watchList;

    Emmiters.userData.subscribe(
      {
        next: (resp) => {
          this.username = resp.username;
        }
      }
    )
  }

  onMovieDelete(movie: Movie) {
    this.username = this.loginService.getUsername();
    console.log("USERNAME: " + this.username);
    console.log("MOVIE BEING DELETED: " + movie.title);
    this.movieService.removeFromWatchList(movie, this.username).subscribe({
      next: (resp) => {
        // remove the movie from the watchList in the backend
        const index = Emmiters.watchList.indexOf(movie);
        if (index > -1) {
          Emmiters.watchList.splice(index, 1);
        }
      },
      error: (error) => {
        console.log(error);
      }
    });
  }

  // public onMovieAddClick(): void {
  //   this.openMovieAddPopup.emit(!this.addMoviePopupStatus);
  //   this.addMoviePopupStatus = !this.addMoviePopupStatus;
  // }
}
