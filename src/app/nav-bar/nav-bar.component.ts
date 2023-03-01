import { Component, Output, Input, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent {
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
