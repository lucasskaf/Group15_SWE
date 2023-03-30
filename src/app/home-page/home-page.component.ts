import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { User } from '../user-auth/user';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {
  isLoginOpen = false
  message = 'Home Page'

  constructor(
    private http : HttpClient
  ){}

  ngOnInit(): void {
    this.http.get<User>("http://localhost:8080/user", {withCredentials: true}).subscribe({
      next: (res) => {
        this.message = `HEY ${res.username}`
        Emmiters.userData.emit(res.username)
        Emmiters.authEmmiter.emit(true)
      },
      error: (err) => {
        this.message = 'You are not logged in'
        Emmiters.authEmmiter.emit(false)
      }
    })
  }

  public toogleLoginStatus(loginStatus : boolean): void{
    this.isLoginOpen = loginStatus;
  }
}
