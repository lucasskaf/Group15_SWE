import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { loginInteraction } from './login-register'
import { Movie, User } from '../user';

import { FormBuilder, Validators } from '@angular/forms';

import { LoginRegisterService } from 'src/app/services/login-register.service';
import { Router } from '@angular/router';
import { Emmiters } from 'src/app/emitters/emmiters';
import { NgToastService } from 'ng-angular-popup'

@Component({
  selector: 'bb-login-register',
  templateUrl: './login-register.component.html',
  styleUrls: ['./login-register.component.css']
})
export class LoginRegisterComponent implements OnInit {
  isLoginOpen: boolean = false

  getLoginStatus = Emmiters.isLoginOpen.subscribe((resp) => {
    this.isLoginOpen = resp
  })

  formSignUp
  formSignIn

  constructor(private formBuilder : FormBuilder,
    private loginRegisterService : LoginRegisterService,
    private router : Router,
    private toast: NgToastService) {}
  
  ngOnInit(): void {
    const loginAnimation = loginInteraction()

    this.formSignUp = this.formBuilder.group({
      username: this.formBuilder.control('', Validators.required),
      password: this.formBuilder.control('', Validators.required)
    })

    this.formSignIn = this.formBuilder.group({
      username: this.formBuilder.control('', Validators.required),
      password: this.formBuilder.control('', Validators.required)
    })

    // Verify if token is already present in cookies
    this.loginRegisterService.getUser().subscribe(
      {
        next: (userInfo) => {
          console.log('PRESENT')
          Emmiters.authEmmiter.emit(true)
          Emmiters.userData.emit(userInfo.username)
          Emmiters.isLoginOpen.emit(false)            
          Emmiters.watchList = userInfo.watchlist
          console.log(`WATCHLIST: ${userInfo.watchlist}`)
        },
        error: (err) => {
          console.log('Error')
          Emmiters.authEmmiter.emit(false)
        }
      }
    )
  }

  closeLogin() {
    Emmiters.isLoginOpen.emit(false)
  }

  onSubmit(user : User){
    this.loginRegisterService.createUser(user).subscribe((resp) => {
      console.log(resp)
      this.router.navigate(['/'])
    })

    this.closeLogin()
  }

  loginUser(user : User) {
    this.loginRegisterService.loginUser(user).subscribe(
     { 
      next: (resp) => {
        this.loginRegisterService.getUser().subscribe({
          next: (userInfo) => {
            Emmiters.authEmmiter.emit(true)
            Emmiters.isLoginOpen.emit(false)
            Emmiters.userData.emit(userInfo.username)
            // Emmiters.watchList.emit(userInfo.watchlist)
            // this.userWatchlist.emit(userInfo.watchlist)
            Emmiters.watchList = userInfo.watchlist
            console.log(`WATCHLIST: ${userInfo.watchlist}`)
            this.toast.success({detail: "Success", summary: "You were logged in!", duration: 4000})
          }
        })
      },
      error: (err) => {
        Emmiters.authEmmiter.emit(false)
        console.log(err)
      }
    }
    )
  }
}