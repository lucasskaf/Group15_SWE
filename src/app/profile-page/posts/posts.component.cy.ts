import { MovieGeneratorService } from "src/app/services/movie-generator.service";
import { HttpClientTestingModule } from "@angular/common/http/testing";
import { TestBed, ComponentFixture } from "@angular/core/testing";
import { PostsComponent } from "../../profile-page/posts/posts.component";

describe('PostComponent', () => {
  let component: PostsComponent;
  let fixturePost: ComponentFixture<PostsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [MovieGeneratorService],
      declarations: [PostsComponent]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixturePost = TestBed.createComponent(PostsComponent);
    component = fixturePost.componentInstance;
    fixturePost.detectChanges();
  });

  it ('should verify that PostComponent is visible', () => {
    cy.get('.posts-container').should('be.visible');
  })
})