import { MainComponent } from "./main.component";
import { NavbarComponent } from "../navbar/navbar.component";
import { WatchedComponent } from "../watched/watched.component";
import { PostsComponent } from "../posts/posts.component";
import { HttpClientTestingModule } from "@angular/common/http/testing";
import { TestBed, ComponentFixture } from "@angular/core/testing";

describe('MainComponent', () => {
  let component: MainComponent;
  let fixture: ComponentFixture<MainComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      declarations: [MainComponent, NavbarComponent, WatchedComponent, PostsComponent]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MainComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should render the app-navbar, app-watched, and app-posts components', () => {
    cy.get('app-navbar').should('be.visible');
    cy.get('app-watched').should('be.visible');
    cy.get('app-posts').should('be.visible');
  });
});