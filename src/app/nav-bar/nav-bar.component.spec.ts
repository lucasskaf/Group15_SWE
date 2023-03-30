import { HttpClient, HttpClientModule } from "@angular/common/http";
import { HttpClientTestingModule, HttpTestingController } from "@angular/common/http/testing";
import { ComponentFixture, TestBed } from "@angular/core/testing";
import { MatButton } from "@angular/material/button";
import { MatIconModule } from "@angular/material/icon";
import { BrowserModule, By } from "@angular/platform-browser";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { Data } from "@angular/router";
import { of } from "rxjs";
import { AppRoutingModule } from "../app-routing.module";
import { ProfilePageModule } from "../profile-page/profile-page.module";
import { NavBarService } from "../services/nav-bar.service";
import { UserAuthModule } from "../user-auth/user-auth.module";
import { NavBarComponent } from "./nav-bar.component";

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
        AppRoutingModule,
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

  it('should call logout', async () => {
    const fixture = TestBed.createComponent(NavBarComponent);
  fixture.detectChanges();

  const component = fixture.componentInstance;

  await fixture.whenStable(); // wait for async operation to complete

  const button = fixture.debugElement.query(By.css('#logout-button')).nativeElement;

  const spy = spyOn(component, 'onLogout');

  button.click();

  expect(spy).toHaveBeenCalled();
    // httpClientSpy.post.and.returnValue(of(USER))
    // const fixture = TestBed.createComponent(NavBarComponent);
    // // spy on event emitter
    // const component = fixture.componentInstance; 
    // spyOn(component, 'onLogout');

    // // trigger the click
    // const nativeElement = fixture.nativeElement;
    // const button = nativeElement.querySelector('#logout-button')
    // button.dispatchEvent(new Event('click'));

    // fixture.detectChanges();

    // expect(component.onLogout).toHaveBeenCalledTimes(0)
  })
})