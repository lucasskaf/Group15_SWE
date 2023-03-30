import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { User } from '../user-auth/user';
import { catchError, Observable, throwError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class LoginRegisterService {
  private handleError(error: HttpErrorResponse) {
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${error.status}, body was: `, error.error);
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error('Something bad happened; please try again later.'));
  }

  constructor(private http : HttpClient) { }

  serverURL : string = 'http://localhost:8080'

  createUser(user : User) : Observable<User> {
    // const httpOptions = {
    //   headers: new Headers({
    //     'Content-Type': 'application/json'
    //   })
    // }

    return this.http.post<User>(`${this.serverURL}/signup`, user)
  }

  // loginUser(user : User) {
  //   return this.http.get(`${this.serverURL}/login`)
  // }
}
