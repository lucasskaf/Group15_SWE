import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MatIconModule } from '@angular/material/icon';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppRoutingModule } from '../app-routing.module';
import { NavBarComponent } from '../nav-bar/nav-bar.component';
import { ProfilePageModule } from '../profile-page/profile-page.module';
import { UserAuthModule } from '../user-auth/user-auth.module';

import { HomePageComponent } from './home-page.component';

describe('HomePageComponent', () => {
  let component: HomePageComponent;
  let fixture: ComponentFixture<HomePageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BrowserModule,
        HttpClientModule,
        AppRoutingModule,
        UserAuthModule,
        ProfilePageModule,
        BrowserAnimationsModule,
        MatIconModule],
      declarations: [ HomePageComponent, NavBarComponent ]
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
