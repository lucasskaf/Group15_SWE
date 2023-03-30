import { Injectable } from '@angular/core'
import { MovieComponent } from '../common/movie/movie.component';

@Injectable({
    providedIn: 'root'
})
export class MovieService {
    watchedMovies: MovieComponent[] = [];

    getWatchedMovies() {
        return this.watchedMovies;
    }

    addWatchedMovie(movie: MovieComponent) {
        this.watchedMovies.push(movie);
    }

    deleteWatchedMovie(movie: MovieComponent) {
        const index = this.watchedMovies.indexOf(movie);
        if (index >= 0) {
            this.watchedMovies.splice(index, 1);
        }
    }
}