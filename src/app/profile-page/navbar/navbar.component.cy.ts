import { NavbarComponent } from "./navbar.component";

describe('NavbarComponent', () => {
    beforeEach(() => {
        cy.mount(NavbarComponent);
    })

    it('clicks on the home button', () => {
        cy.get('#home-button').click();
    });

    it('clicks on the logout button', () => {
        cy.get('#logout-button').click();
    });
})