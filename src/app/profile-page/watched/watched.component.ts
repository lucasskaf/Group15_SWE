import { Component, EventEmitter, OnInit, Output, Input } from '@angular/core';
import { MovieService } from '../../services/movie-service';
import { MovieComponent } from '../../common/movie/movie.component';

@Component({
  selector: 'app-watched',
  templateUrl: './watched.component.html',
  styleUrls: ['./watched.component.css']
})
export class WatchedComponent implements OnInit {
  @Output() openMovieAddPopup = new EventEmitter<boolean>();
  @Input() addMoviePopupStatus: boolean = false;
  
  movies: MovieComponent[] = [];

  constructor(private movieService: MovieService) {}

  ngOnInit() {
    this.movies = this.movieService.getWatchedMovies();
  }

  onMovieDelete(movie: MovieComponent) {
    this.movieService.deleteWatchedMovie(movie);
  }

  public onMovieAddClick(): void {
    this.openMovieAddPopup.emit(!this.addMoviePopupStatus);
    this.addMoviePopupStatus = !this.addMoviePopupStatus;
  }
}
