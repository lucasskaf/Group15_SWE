import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import { Movie } from './movie';
// import {environment} from '../environments/environments';

import { Observable, throwError } from 'rxjs';
// import { catchError, retry } from 'rxjs/operators';
// import { map } from 'rxjs/operators';
// import 'rxjs/add/operator/map';

@Injectable()
export class MovieTestService {

  private dataURL: string = "../assets/movies-data-test.json"
  // ${environment.serverUrl}/create

  constructor(private http: HttpClient) { }

  getTitle() : Observable<Movie[]> {
    return this.http.get<Movie[]>(this.dataURL)
  }

}