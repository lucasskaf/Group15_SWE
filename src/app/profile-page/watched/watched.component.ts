import { Component, EventEmitter, OnInit, Output, Input } from '@angular/core';
import { Emmiters } from 'src/app/emitters/emmiters';
import { Movie } from 'src/app/user-auth/user';
import { MovieGeneratorService } from '../../services/movie-generator.service';

@Component({
  selector: 'app-watched',
  templateUrl: './watched.component.html',
  styleUrls: ['./watched.component.css']
})
export class WatchedComponent implements OnInit {
  @Output() openMovieAddPopup = new EventEmitter<boolean>();
  @Input() addMoviePopupStatus: boolean = false;
  
  movies: Movie[] = [];

  constructor(private movieService: MovieGeneratorService) {}

  ngOnInit() {
    this.movies = Emmiters.watchList;
  }

  onMovieDelete(movie: Movie) {
    this.movieService.removeFromWatchList(movie).subscribe(
      (response) => {
        // remove the movie from the watchList in the backend
        const index = this.movies.indexOf(movie);
        if (index > -1) {
          this.movies.splice(index, 1);
        }
      },
      (error) => {
        console.log(error);
      }
    )
  }

  public onMovieAddClick(): void {
    this.openMovieAddPopup.emit(!this.addMoviePopupStatus);
    this.addMoviePopupStatus = !this.addMoviePopupStatus;
  }
}
