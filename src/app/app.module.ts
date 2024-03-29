import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
// import { AppRoutingModule } from './app-routing.module';
import { routing } from './app-routing.module';
import { AppComponent } from './app.component';
import { ProfilePageModule } from './profile-page/profile-page.module';
import { MovieGeneratorComponent } from './movie-generator/movie-generator.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { UserAuthModule } from './user-auth/user-auth.module';
import { HomePageComponent } from './home-page/home-page.component';
import { MatIconModule } from '@angular/material/icon'
import { NgToastModule } from 'ng-angular-popup';
import {MatFormFieldModule} from '@angular/material/form-field';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { MatChipsModule } from '@angular/material/chips';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatSelectModule } from '@angular/material/select'
import { MatSliderModule } from '@angular/material/slider'
import {MatInputModule} from '@angular/material/input';
import { CarouselModule } from 'primeng/carousel'
import { RatingModule } from 'primeng/rating';
import { DialogModule } from 'primeng/dialog';
import { ButtonModule } from 'primeng/button';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { InputTextModule } from 'primeng/inputtext';
import { InputTextareaModule } from 'primeng/inputtextarea';
import { VirtualScrollerModule } from 'primeng/virtualscroller';
import { MoviePopupComponent } from './movie-popup/movie-popup.component';

@NgModule({
  declarations: [
    AppComponent,
    NavBarComponent,
    MovieGeneratorComponent,
    HomePageComponent,
    MoviePopupComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    routing,
    UserAuthModule,
    ProfilePageModule,
    BrowserAnimationsModule,
    MatIconModule,
    MatFormFieldModule,
    NgToastModule,
    ReactiveFormsModule,
    MatChipsModule,
    MatAutocompleteModule,
    MatSelectModule,
    MatSliderModule,
    FormsModule,
    MatInputModule,
    CarouselModule,
    RatingModule,
    DialogModule,
    ButtonModule,
    ToastModule,
    InputTextModule,
    InputTextareaModule,
    VirtualScrollerModule
  ],
  providers: [
    MessageService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }