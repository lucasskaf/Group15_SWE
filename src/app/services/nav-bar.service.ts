import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class NavBarService {

  constructor(private http : HttpClient) { }

  serverURL : string = 'http://localhost:8080'

  logout() {
    return this.http.post(`${this.serverURL}/logout`, {}, {withCredentials: true})
  }
}
