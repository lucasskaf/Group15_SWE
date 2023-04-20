import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatIconModule } from '@angular/material/icon';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { routing } from '../app-routing.module';
import { NavBarComponent } from '../nav-bar/nav-bar.component';
import { ProfilePageModule } from '../profile-page/profile-page.module';
import { UserAuthModule } from '../user-auth/user-auth.module';
import { MovieGeneratorComponent } from '../movie-generator/movie-generator.component';
import { CarouselModule } from 'primeng/carousel'
import { MoviePopupComponent } from '../movie-popup/movie-popup.component';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import { RatingModule } from 'primeng/rating';
import {MatSliderModule} from '@angular/material/slider';
import { DialogModule } from 'primeng/dialog';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';


import { HomePageComponent } from './home-page.component';
import { MovieGeneratorService } from '../services/movie-generator.service';

describe('HomePageComponent', () => {
  let component: HomePageComponent;
  let fixture: ComponentFixture<HomePageComponent>;
  let movieComponentService: MovieGeneratorService
  let httpClientSpy: jasmine.SpyObj<HttpClient>
  let USER = {
    username: "own"
  }

  beforeEach(async () => {
    httpClientSpy = jasmine.createSpyObj('HttpClient', ['post'])
    movieComponentService = new MovieGeneratorService(httpClientSpy)
    await TestBed.configureTestingModule({
      imports: [BrowserModule,
        HttpClientModule,
        routing,
        UserAuthModule,
        ProfilePageModule,
        BrowserAnimationsModule,
        MatIconModule,
        CarouselModule,
        MatFormFieldModule,
        MatSelectModule,
        RatingModule,
        MatSliderModule,
        DialogModule,
        ReactiveFormsModule,
        FormsModule,
        MatInputModule
      ],
      declarations: [ HomePageComponent, NavBarComponent, MovieGeneratorComponent, MoviePopupComponent],
      providers: [MovieGeneratorService]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomePageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
