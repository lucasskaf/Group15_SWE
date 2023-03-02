import { NavBarComponent } from "src/app/nav-bar/nav-bar.component";
import { NavbarComponent } from "./navbar.component";

describe('NavbarComponent', () => {
    it('can mount', () => {
        cy.mount(NavbarComponent)
    })

    it('clicks the sidenav button', () => {
        cy.mount(NavbarComponent, {
            componentProperties: {
                openSidenav: {
                    emit: cy.spy().as('onClickSpy')
                } as any
            }
        })
        cy.get('#sidenav-button').click();
        cy.get('@onClickSpy').should('have.been.called')
    })
})