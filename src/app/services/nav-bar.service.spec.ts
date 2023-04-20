import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';
import { MatIconModule } from '@angular/material/icon';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { routing } from '../app-routing.module';
import { ProfilePageModule } from '../profile-page/profile-page.module';
import { UserAuthModule } from '../user-auth/user-auth.module';

import { NavBarService } from './nav-bar.service';

describe('NavBarService', () => {
  let service: NavBarService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [BrowserModule,
        HttpClientModule,
        routing,
        UserAuthModule,
        ProfilePageModule,
        BrowserAnimationsModule,
        MatIconModule],
      providers: [HttpClientModule]
    });
    service = TestBed.inject(NavBarService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
