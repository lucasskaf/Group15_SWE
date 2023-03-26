import { Component, OnInit } from '@angular/core';
import { MovieService } from '../../services/movie-service';
import { MovieComponent } from '../../common/movie/movie.component';

@Component({
  selector: 'app-watched',
  templateUrl: './watched.component.html',
  styleUrls: ['./watched.component.css']
})
export class WatchedComponent implements OnInit {
  movies: MovieComponent[] = [];

  printMovies() {
    for (let n of this.movies) {
      console.log(n);
    }
  }

  constructor(private movieService: MovieService) {}

  ngOnInit() {
    this.movies = this.movieService.getWatchedMovies();
    this.printMovies();
  }

  onMovieDelete(movie: MovieComponent) {
    this.movieService.deleteWatchedMovie(movie);
  }
}
