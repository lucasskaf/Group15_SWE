import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { loginInteraction } from './login-register'
import { User } from '../user';

import { FormBuilder, Validators } from '@angular/forms';

import { LoginRegisterService } from 'src/app/services/login-register.service';

@Component({
  selector: 'bb-login-register',
  templateUrl: './login-register.component.html',
  styleUrls: ['./login-register.component.css']
})
export class LoginRegisterComponent implements OnInit {
  @Output() isClosed = new EventEmitter()
  @Input() isOpen : boolean = false

  formSignUp

  constructor(private formBuilder : FormBuilder,
    private loginRegisterService : LoginRegisterService) {}
  
  ngOnInit(): void {
    const loginAnimation = loginInteraction()

    this.formSignUp = this.formBuilder.group({
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
    })

    this.isOpen = false
  }
}
