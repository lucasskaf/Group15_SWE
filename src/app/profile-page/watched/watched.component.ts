import { Component, OnInit } from '@angular/core';
import { MovieService } from '../../services/movie-service';
import { MovieFormComponent } from '../movie-form/movie-form.component';
import { Movie } from '../../common/movie';

@Component({
  selector: 'app-watched',
  templateUrl: './watched.component.html',
  styleUrls: ['./watched.component.css']
})
export class WatchedComponent implements OnInit {
  movies: Movie[] = [];

  constructor(private movieService: MovieService) {}

  ngOnInit() {
    this.movies = this.movieService.getWatchedMovies();
  }

  onMovieDelete(movie: Movie) {
    this.movieService.deleteWatchedMovie(movie);
  }

  onSubmit(movie: Movie) {
    console.log(movie);
    this.movieService.addWatchedMovie(movie);
  }
}
