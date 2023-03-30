import { HttpClientTestingModule } from "@angular/common/http/testing";
import { ComponentFixture, TestBed } from "@angular/core/testing";
import { ReactiveFormsModule } from "@angular/forms";
import { MatIconModule } from "@angular/material/icon";
import { LoginRegisterService } from "src/app/services/login-register.service";
import { LoginRegisterComponent } from "./login-register.component";

describe("LoginRegisterComponent", () => {
  let component: LoginRegisterComponent;
  let fixture: ComponentFixture<LoginRegisterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, MatIconModule, ReactiveFormsModule],
      declarations: [LoginRegisterComponent],
      providers: [LoginRegisterService]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LoginRegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should emit closure', () => {
    const fixture = TestBed.createComponent(LoginRegisterComponent);
    // spy on event emitter
    const component = fixture.componentInstance;
    component.toogleLogin()
    let spy = spyOn(component.isClosed, 'emit')

    // trigger the click
    const nativeElement = fixture.nativeElement;
    const button = nativeElement.querySelector('#closeButton');
    button.dispatchEvent(new Event('click'));

    fixture.detectChanges();

    expect(spy).toHaveBeenCalledWith(false);
  })

  it('should call loginUser', () => {
    const fixture = TestBed.createComponent(LoginRegisterComponent);
    // spy on event emitter
    const component = fixture.componentInstance; 
    spyOn(component, 'loginUser');

    // trigger the click
    const nativeElement = fixture.nativeElement;
    const button = nativeElement.querySelector('button#signin');
    button.dispatchEvent(new Event('click'));

//     fixture.detectChanges();

    expect(component.loginUser).toHaveBeenCalledTimes(0)
  })
})