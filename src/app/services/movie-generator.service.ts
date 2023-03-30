import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Movie } from '../user-auth/user';
import { map } from 'rxjs';

interface MovieResponse {
  movie : Movie
}

@Injectable({
  providedIn: 'root'
})
export class MovieGeneratorService {

  constructor(private httpClient : HttpClient) { }

  serverURL : string = 'http://localhost:8080'

  getMovie() {
    return this.httpClient.get<MovieResponse>(`${this.serverURL}/generate`).pipe(map(response => {
      return response.movie
    }))
  }
}
