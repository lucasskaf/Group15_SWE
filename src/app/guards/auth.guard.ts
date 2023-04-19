import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { LoginRegisterService } from '../services/login-register.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {

  isUserAuthenticated: boolean = true;
  constructor(private loginService: LoginRegisterService) { };

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {

    // returns true if the user is logged in, else, returns false
    console.log(this.loginService.getUserAuth());
    return this.loginService.getUserAuth();
  }

}
