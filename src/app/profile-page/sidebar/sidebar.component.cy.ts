import { SidebarComponent } from "./sidebar.component";
import { NavbarComponent } from "../navbar/navbar.component"

describe('SidebarComponent', () => {
  beforeEach(() => {
    cy.mount(NavbarComponent);
    cy.mount(SidebarComponent);
  })

  it('mat-drawer should be visible initially', () => {
    cy.mount(SidebarComponent, {
      componentProperties: {
        isOpen: true,
        isClosed: cy.spy().as('isClosedSpy')
      } as any
    })
    cy.get('mat-drawer').should('be.visible');
  })
})