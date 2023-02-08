import { Component } from '@angular/core';
import { loginInteraction } from './login-register.js'

@Component({
  selector: 'app-login-register',
  templateUrl: './login-register.component.html',
  styleUrls: ['./login-register.component.css']
})
export class LoginRegisterComponent {
  
  ngOnInit(): void {
    const functionality = loginInteraction()
  }

}
