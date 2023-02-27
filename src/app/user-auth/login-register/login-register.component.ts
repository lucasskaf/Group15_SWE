import { Component } from '@angular/core';
import { loginInteraction } from './login-register'
import { closeLogin } from './login-register'

@Component({
  selector: 'app-login-register',
  templateUrl: './login-register.component.html',
  styleUrls: ['./login-register.component.css']
})
export class LoginRegisterComponent {
  
  ngOnInit(): void {
    const functionality = loginInteraction()
    const closeButton = closeLogin()
  }

}