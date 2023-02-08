// import { Component } from '@angular/core';

// @Component({
//   selector: 'app-root',
//   templateUrl: './app.component.html',
//   styleUrls: ['./app.component.css']
// })
// export class AppComponent {
//   title = 'BingeBuddy';
// }

import {Component} from '@angular/core';
import { loginInteraction } from './loginModal.js'

@Component({
  selector: 'LoginModal',
  templateUrl: './loginModal.component.html',
  styleUrls: ['./loginModal.component.css']
})
export class LoginModal {
  title = 'TestMovie'

  role = 'Admin'

  ngOnInit(): void {
    const functionality = loginInteraction()
  }
}
