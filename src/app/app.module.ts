import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { ProfilePageModule } from './profile-page/profile-page.module';
import { MovieGeneratorComponent } from './movie-generator/movie-generator.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { UserAuthModule } from './user-auth/user-auth.module';
import { HomePageComponent } from './home-page/home-page.component';
import { MatIconModule } from '@angular/material/icon'
import { Routes } from '@angular/router';
import { routing } from './app-routing.module';

@NgModule({
  declarations: [
    AppComponent,
    NavBarComponent,
    MovieGeneratorComponent,
    HomePageComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    UserAuthModule,
    ProfilePageModule,
    BrowserAnimationsModule,
    MatIconModule,
    routing
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }