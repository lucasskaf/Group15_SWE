import { Component, EventEmitter, OnInit, Output, Input } from '@angular/core';
import { MovieService } from '../../services/movie-service';
import { MovieComponent } from '../../common/movie/movie.component';
import { AddPopupComponent } from '../add-popup/add-popup.component';

@Component({
  selector: 'app-watched',
  templateUrl: './watched.component.html',
  styleUrls: ['./watched.component.css']
})
export class WatchedComponent implements OnInit {
  @Output() openAddPopup = new EventEmitter<boolean>();
  @Input() addPopupStatus: boolean = false;
  
  movies: MovieComponent[] = [];

  constructor(private movieService: MovieService) {}

  ngOnInit() {
    this.movies = this.movieService.getWatchedMovies();
  }

  onMovieDelete(movie: MovieComponent) {
    this.movieService.deleteWatchedMovie(movie);
  }

  public onAddClick(): void {
    this.openAddPopup.emit(!this.addPopupStatus);
    this.addPopupStatus = !this.addPopupStatus;
  }
}
