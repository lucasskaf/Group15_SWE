import { PostsComponent } from "./posts.component";

describe('PostsComponent', () => {
  beforeEach(() => {
    cy.mount(PostsComponent);
  })

  it ('should verify that PostsComponent is visible', () => {
    cy.get('.posts-container').should('be.visible');
  })
})