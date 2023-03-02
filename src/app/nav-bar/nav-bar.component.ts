import { Component, Output, Input, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent {
  @Output() openLogin = new EventEmitter<boolean>();
  @Input() loginStatus: boolean = false;
  
  constructor() { }

  ngOnInit() {

  }

  public onLoginClick() {
    this.openLogin.emit(!this.loginStatus)
    this.loginStatus = !this.loginStatus
  }
}
