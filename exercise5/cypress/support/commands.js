// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************

// -- This is a parent command --
// Cypress.Commands.add('login', (email, password) => { ... })

// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })

// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })

// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })

// Custom command to test accessibility
Cypress.Commands.add('checkA11y', (context = null, options = null) => {
  if (Cypress.env('A11Y_TEST') === false) {
    return;
  }

  cy.log('Checking accessibility');
  // This would use axe-core if you install the cypress-axe plugin
  // cy.injectAxe();
  // cy.checkA11y(context, options);
});

// Custom command to add a product to cart and verify it's added
Cypress.Commands.add('addProductToCart', (productIndex = 0) => {
  cy.intercept('POST', '**/carts/*/items').as('addToCart');

  cy.get('.card').eq(productIndex).find('button').contains('Add to Cart').click();
  cy.wait('@addToCart');

  // Verify cart count increases
  cy.get('[data-testid="cart-count"]').should('be.visible');
});

// Custom command to fill payment form with test data
Cypress.Commands.add('fillPaymentForm', () => {
  // Fill billing information
  cy.get('input#name').type('John Doe');
  cy.get('input#email').type('john.doe@example.com');
  cy.get('input#address').type('123 Main St');
  cy.get('input#city').type('Anytown');
  cy.get('input#postalCode').type('12345');

  // Fill payment details
  cy.get('input#cardName').type('John Doe');
  cy.get('input#cardNumber').type('4111111111111111');
  cy.get('input#expiryDate').type('12/25');
  cy.get('input#cvv').type('123');
});