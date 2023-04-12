import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Movie } from '../user-auth/user';
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

  getRandomMovie(filters : movieFilters) {
    return this.httpClient.post<Movie>(`${this.serverURL}/generate/filters`, filters)
  }
}
