import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginRegisterComponent } from './login-register/login-register.component';
import { MatIconModule } from '@angular/material/icon';

@NgModule({
  declarations: [
    LoginRegisterComponent
  ],
  imports: [
    CommonModule,
    MatIconModule
  ],
  exports: [
    LoginRegisterComponent
  ]
})
export class UserAuthModule { }
