import { Component, Output, Input, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  @Output() openSidenav = new EventEmitter<boolean>();
  @Input() sidenavStatus: boolean = false;

  @Output() openLogin = new EventEmitter<boolean>();
  @Input() loginStatus: boolean = false;
  

  constructor() { }

  ngOnInit() {

  }

  public onSidenavClick(): void {
    this.openSidenav.emit(!this.sidenavStatus);
    this.sidenavStatus = !this.sidenavStatus
  }

  public onLoginClick() {
    this.openLogin.emit(!this.loginStatus)
    this.loginStatus = !this.loginStatus
  }
}
