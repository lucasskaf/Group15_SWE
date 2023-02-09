import { Component } from '@angular/core';
import { loginInteraction, closeLogin } from './login-register'

@Component({
  selector: 'app-login-register',
  templateUrl: './login-register.component.html',
  styleUrls: ['./login-register.component.css']
})
export class LoginRegisterComponent {
  
  ngOnInit(): void {
    const loginAnimation = loginInteraction()
    const closeFunctionality = closeLogin()
  }

}
