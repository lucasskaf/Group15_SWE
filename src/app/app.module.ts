import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ProfilePageModule } from './profile-page/profile-page.module'
import { UserAuthModule } from './user-auth/user-auth.module';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    UserAuthModule,
    ProfilePageModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
