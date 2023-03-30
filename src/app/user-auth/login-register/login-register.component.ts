import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { loginInteraction } from './login-register'
import { User } from '../user';

import { FormBuilder, Validators } from '@angular/forms';

import { LoginRegisterService } from 'src/app/services/login-register.service';
import { Router } from '@angular/router';
import { Emmiters } from 'src/app/emitters/emmiters';

@Component({
  selector: 'bb-login-register',
  templateUrl: './login-register.component.html',
  styleUrls: ['./login-register.component.css']
})
export class LoginRegisterComponent implements OnInit {
  @Output() isClosed = new EventEmitter()
  @Input() isOpen : boolean = false

  formSignUp
  formSignIn

  constructor(private formBuilder : FormBuilder,
    private loginRegisterService : LoginRegisterService,
    private router : Router) {}
  
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
  }

  toogleLogin() {
    this.isOpen = !this.isOpen
    this.isClosed.emit(this.isOpen)
  }

  onSubmit(user : User){
    console.log(`Input User:`)
    console.log(user)

    const postUser : User = {
      "username": "",
      "password" : ""
    }

    postUser.username = user.username
    postUser.password = user.password

    console.log(`Post User:`)
    console.log(postUser)

    this.loginRegisterService.createUser(postUser).subscribe((resp) => {
      console.log(resp)
      this.router.navigate(['/'])
    })

    this.isOpen = false
  }

  loginUser(user : User) {
    this.loginRegisterService.loginUser(user).subscribe((resp) => {
      Emmiters.authEmmiter.emit(true)
      console.log(resp),
      window.location.reload(),
      // this.snackBar.open('Login Sucessful', "", {
      //   duration: 3000
      // }),
      (err) => {
        Emmiters.authEmmiter.emit(false)
        console.log(err)
      }
    })
  }
}
