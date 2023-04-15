import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginRegisterComponent } from './login-register/login-register.component';
import { MatIconModule } from '@angular/material/icon'
import { ReactiveFormsModule } from '@angular/forms';
import { ToastrModule } from 'ngx-toastr'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { NgToastModule } from 'ng-angular-popup';


@NgModule({
  declarations: [
    LoginRegisterComponent
  ],
  imports: [
    CommonModule,
    MatIconModule,
    ReactiveFormsModule,
    NgToastModule
  ],
  exports: [
    LoginRegisterComponent
  ]
})
export class UserAuthModule { }
