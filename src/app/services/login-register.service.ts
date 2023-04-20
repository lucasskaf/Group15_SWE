import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { User } from '../user-auth/user';
import { catchError, lastValueFrom, Observable, throwError } from 'rxjs';
import { Emmiters } from '../emitters/emmiters';

@Injectable({
  providedIn: 'root'
})
export class LoginRegisterService {

  constructor(private http : HttpClient) { }

  userAuthStatus: boolean = false;
  userName: string = "";

  serverURL : string = 'http://localhost:8080'

  createUser(user : User) : Observable<User> {
    return this.http.post<User>(`${this.serverURL}/signup`, user)
  }

  loginUser(user : User) {
    return this.http.post<User>(`${this.serverURL}/login`, user, {withCredentials: true})
  }

  getUser() {
    return this.http.get<User>("http://localhost:8080/user", {withCredentials: true})
  }

  getUserAuth(): boolean {
    return this.userAuthStatus;
  }

  setUserAuth(auth: boolean) {
    this.userAuthStatus = auth;
  }

  getUsername(): string {
    return this.userName;
  }

  setUsername(name: string){
    this.userName = name;
  }
}
