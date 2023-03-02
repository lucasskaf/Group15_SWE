import { ComponentFixture, TestBed } from "@angular/core/testing";
import { NavBarComponent } from "./nav-bar.component";

describe("NavBarComponent", () => {
  let component: NavBarComponent;
  let fixture: ComponentFixture<NavBarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [NavBarComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NavBarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should emit on click (openStatus)', () => {
    const fixture = TestBed.createComponent(NavBarComponent);
    // spy on event emitter
    const component = fixture.componentInstance; 
    spyOn(component.openLogin, 'emit');

    // trigger the click
    const nativeElement = fixture.nativeElement;
    const button = nativeElement.querySelector('button#login-button');
    button.dispatchEvent(new Event('click'));

    fixture.detectChanges();

    expect(component.openLogin.emit).toHaveBeenCalledWith(true);
  })
})