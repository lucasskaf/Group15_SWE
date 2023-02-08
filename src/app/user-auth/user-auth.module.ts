import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginRegisterComponent } from './login-register/login-register.component';



@NgModule({
  declarations: [
    LoginRegisterComponent
  ],
  imports: [
    CommonModule

  ],
  exports: [
    LoginRegisterComponent
  ]
})
export class UserAuthModule { }
