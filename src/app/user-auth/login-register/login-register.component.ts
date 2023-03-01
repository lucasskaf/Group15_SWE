import { Component, Input, Output, EventEmitter } from '@angular/core';
import { loginInteraction } from './login-register'

@Component({
  selector: 'bb-login-register',
  templateUrl: './login-register.component.html',
  styleUrls: ['./login-register.component.css']
})
export class LoginRegisterComponent {
  @Output() isClosed = new EventEmitter()
  @Input() isOpen : boolean = false
  
  ngOnInit(): void {
    const loginAnimation = loginInteraction()
  }

  toogleLogin() {
    this.isOpen = !this.isOpen
    this.isClosed.emit(this.isOpen)
  }
}