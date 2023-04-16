import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Movie, moviePosts } from '../user-auth/user';
import { map } from 'rxjs';

interface movieFilters {
  actors: string[],
  max_runtime: number,
  genres: number[],
  streaming_providers: number[]
}

@Injectable({
  providedIn: 'root'
})
export class MovieGeneratorService {

  constructor(private httpClient : HttpClient) { }

  serverURL : string = 'http://localhost:8080'

  getRandomMovieWithFilters(filters : movieFilters) {
    return this.httpClient.post<Movie>(`${this.serverURL}/generate/filters`, filters)
  }

  getRandomMovie() {
    return this.httpClient.get<Movie[]>(`${this.serverURL}/generate`)
  }

  addToWatchList(movie: Movie){
    return this.httpClient.post<Movie>(`${this.serverURL}/watchlist/add`, movie, {withCredentials: true})
  }

  addMoviePosts(post: moviePosts){
    return this.httpClient.post<moviePosts>(`${this.serverURL}/posts`, post, {withCredentials: true})
  }
}
