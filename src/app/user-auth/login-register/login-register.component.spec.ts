// import { ComponentFixture, TestBed } from "@angular/core/testing";
// import { LoginRegisterComponent } from "./login-register.component";

// describe("LoginRegisterComponent", () => {
//   let component: LoginRegisterComponent;
//   let fixture: ComponentFixture<LoginRegisterComponent>;

//   beforeEach(async () => {
//     await TestBed.configureTestingModule({
//       declarations: [LoginRegisterComponent],
//     }).compileComponents();
//   });

//   beforeEach(() => {
//     fixture = TestBed.createComponent(LoginRegisterComponent);
//     component = fixture.componentInstance;
//     fixture.detectChanges();
//   });

//   it('should emit closure', () => {
//     const fixture = TestBed.createComponent(LoginRegisterComponent);
//     // spy on event emitter
//     const component = fixture.componentInstance; 
//     spyOn(component.isClosed, 'emit');

//     // trigger the click
//     const nativeElement = fixture.nativeElement;
//     const button = nativeElement.querySelector('button#closeButton');
//     button.dispatchEvent(new Event('click'));

//     fixture.detectChanges();

//     expect(component.isClosed.emit).toHaveBeenCalledWith(true);
//   })
// })
