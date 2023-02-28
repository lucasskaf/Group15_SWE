import { Component, EventEmitter, Output, Input } from '@angular/core';

@Component({
  selector: 'bb-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent {
  @Output() openLogin = new EventEmitter()
  @Input() loginStatus : boolean = false
  // loginStatus = false

  onLoginClick() {
    console.log("Login Opened")
    this.openLogin.emit(!this.loginStatus)
    this.loginStatus = !this.loginStatus
  }
}
