import { EventEmitter } from "@angular/core";
import { Movie, moviePosts } from "../user-auth/user";

export class Emmiters {
    static authEmmiter = new EventEmitter<boolean>()
    static userData = new EventEmitter<string>()
    static isLoginOpen = new EventEmitter<boolean>()
    static watchList: Movie[] = []
    static isMovieWatched = new EventEmitter<boolean>()
    static isPopupOpen = new EventEmitter<boolean>()
    static generatedMovie = new EventEmitter<Movie>()
    static userPosts: moviePosts[] = []
}