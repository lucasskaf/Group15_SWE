import { HttpClient } from '@angular/common/http';
import { Component, EventEmitter, Output, Input } from '@angular/core';
import { Router } from '@angular/router';
import { Emmiters } from '../emitters/emmiters';
import { NavBarService } from '../services/nav-bar.service';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent {
  @Output() openLogin = new EventEmitter()
  @Input() loginStatus : boolean = false
  authethicated = false
  username = ""

  constructor(private navBarService: NavBarService,
    private router : Router) { }

  ngOnInit() {
    Emmiters.authEmmiter.subscribe(
      {
        next: (auth : boolean) => {
          this.authethicated = auth
        }
      }
    )
    Emmiters.userData.subscribe(
      {
        next: (usernameResponse : string) => {
          this.username = usernameResponse
        }
      }
    )
  }

  onLoginClick() {
    this.openLogin.emit(!this.loginStatus)
    this.loginStatus = !this.loginStatus
  }

  onLogout() {
    this.navBarService.logout().subscribe(
      () => {
        this.authethicated = false
        Emmiters.authEmmiter.emit(false)
        window.location.reload()
      }
    )
  }
}
