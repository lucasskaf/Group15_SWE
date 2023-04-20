import { WatchedComponent } from "./watched.component";
import { MovieGeneratorService } from "src/app/services/movie-generator.service";
import { HttpClientTestingModule } from "@angular/common/http/testing";
import { TestBed, ComponentFixture } from "@angular/core/testing";
import { PostComponent } from "src/app/common/post/post.component";

describe('WatchedComponent', () => {
  let component: WatchedComponent;
  let fixtureWatched: ComponentFixture<WatchedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [MovieGeneratorService],
      declarations: [WatchedComponent]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixtureWatched = TestBed.createComponent(WatchedComponent);
    component = fixtureWatched.componentInstance;
    fixtureWatched.detectChanges();
  });

  it ('should verify that WatchedComponent is visible', () => {
    cy.get('.watched-container').should('be.visible');
  })
})