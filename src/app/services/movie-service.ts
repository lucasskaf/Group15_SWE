import { Injectable } from '@angular/core'
import { Movie } from '../common/movie'

@Injectable({
    providedIn: 'root'
})
export class MovieService {
    watchedMovies: Movie[] = [];

    getWatchedMovies() {
        return this.watchedMovies;
    }

    addWatchedMovie(movie: Movie) {
        this.watchedMovies.push(movie);
    }

    deleteWatchedMovie(movie: Movie) {
        const index = this.watchedMovies.indexOf(movie);
        if (index >= 0) {
            this.watchedMovies.splice(index, 1);
        }
    }
}