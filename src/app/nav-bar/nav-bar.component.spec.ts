import { HttpClient, HttpClientModule } from "@angular/common/http";
import { HttpClientTestingModule, HttpTestingController } from "@angular/common/http/testing";
import { ComponentFixture, TestBed } from "@angular/core/testing";
import { MatButton } from "@angular/material/button";
import { MatIconModule } from "@angular/material/icon";
import { BrowserModule, By } from "@angular/platform-browser";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { Data } from "@angular/router";
import { of } from "rxjs";
import { routing } from "../app-routing.module";
import { ProfilePageModule } from "../profile-page/profile-page.module";
import { NavBarService } from "../services/nav-bar.service";
import { UserAuthModule } from "../user-auth/user-auth.module";
import { NavBarComponent } from "./nav-bar.component";
import { Emmiters } from "../emitters/emmiters";

describe("NavBarComponent", () => {
  let component: NavBarComponent;
  let fixture: ComponentFixture<NavBarComponent>;
  let navBarService: NavBarService
  let httpClientSpy: jasmine.SpyObj<HttpClient>
  let USER = {
    username: "own"
  }

  beforeEach(async () => {
    httpClientSpy = jasmine.createSpyObj('HttpClient', ['post'])
    navBarService = new NavBarService(httpClientSpy)

    await TestBed.configureTestingModule({
    imports: [BrowserModule,
        HttpClientModule,
        routing,
        UserAuthModule,
        ProfilePageModule,
        BrowserAnimationsModule,
        MatIconModule],
      declarations: [NavBarComponent],
      providers: [NavBarService]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NavBarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should emit on click (openStatus)', () => {
    const fixture = TestBed.createComponent(NavBarComponent);
    fixture.detectChanges();

    const component = fixture.componentInstance;
    const button = fixture.debugElement.query(By.css('#login-button')).nativeElement;

    const spy = spyOn(component, 'onLoginClick');

    button.click();

    expect(spy).toHaveBeenCalled();
  })

  it('should emit authenticated false', () => {
    const fixture = TestBed.createComponent(NavBarComponent);
    fixture.detectChanges();

    const component = fixture.componentInstance;

    component.onLogout()

    const spy = spyOn(component, 'onLogout');

    let auth = component.authethicated

    expect(auth).toBeFalse();
  })
})