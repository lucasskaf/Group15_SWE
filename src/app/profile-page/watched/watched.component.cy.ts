import { WatchedComponent } from "./watched.component";

describe('WatchedComponent', () => {
  beforeEach(() => {
    cy.mount(WatchedComponent);
  })

  it ('should verify that WatchedComponent is visible', () => {
    cy.get('.watched-container').should('be.visible');
  })
})