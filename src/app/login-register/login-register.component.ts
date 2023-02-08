import { Component } from '@angular/core';

@Component({
  selector: 'app-login-register',
  templateUrl: './login-register.component.html',
  styleUrls: ['./login-register.component.css']
})
export class LoginRegisterComponent {
  container = document.getElementById('container');

  toggleLogIn () {
    this.container?.classList.remove("right-panel-active");
  }
  
  toggleSignUp () {
    this.container?.classList.add("right-panel-active");
  }

}
