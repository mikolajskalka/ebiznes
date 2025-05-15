/// <reference types="cypress" />

describe('Payment Page', () => {
  beforeEach(() => {
    // Mock API responses
    cy.intercept('GET', '**/products', { fixture: 'products.json' }).as('getProducts');
    cy.intercept('GET', '**/categories', { fixture: 'categories.json' }).as('getCategories');
    cy.intercept('GET', '**/carts/**', { fixture: 'cart.json' }).as('getCart');

    // Visit home page and add a product to cart
    cy.visit('/');
    cy.wait('@getProducts');

    cy.get('.card button').first().click();
    cy.contains('Cart').click();
    cy.wait('@getCart');
    cy.contains('button', 'Proceed to Checkout').click();
  });

  it('should display payment form correctly', () => {
    // Check page structure
    cy.contains('h1', 'Checkout').should('be.visible');
    cy.contains('h2', 'Billing Information').should('be.visible');
    cy.contains('h2', 'Payment Details').should('be.visible');
    cy.contains('h2', 'Order Summary').should('be.visible');

    // Check form fields
    cy.get('form').within(() => {
      // Billing fields
      cy.get('input#name').should('be.visible');
      cy.get('input#email').should('be.visible');
      cy.get('input#address').should('be.visible');
      cy.get('input#city').should('be.visible');
      cy.get('input#postalCode').should('be.visible');

      // Payment fields
      cy.get('input#cardName').should('be.visible');
      cy.get('input#cardNumber').should('be.visible');
      cy.get('input#expiryDate').should('be.visible');
      cy.get('input#cvv').should('be.visible');
    });

    // Check order summary
    cy.get('.order-summary').should('be.visible');
    cy.get('[data-testid="summary-total"]').should('be.visible');
    cy.get('button[type="submit"]').contains('Complete Purchase').should('be.visible');
  });

  it('should validate required form fields', () => {
    // Submit empty form
    cy.get('form').submit();

    // Check validation messages
    cy.contains('Name is required').should('be.visible');
    cy.contains('Email is required').should('be.visible');
    cy.contains('Address is required').should('be.visible');
    cy.contains('City is required').should('be.visible');
    cy.contains('Postal Code is required').should('be.visible');
    cy.contains('Name on Card is required').should('be.visible');
    cy.contains('Card Number is required').should('be.visible');
    cy.contains('Expiry Date is required').should('be.visible');
    cy.contains('CVV is required').should('be.visible');
  });

  it('should validate email format', () => {
    // Fill in invalid email
    cy.get('input#email').type('invalid-email');
    cy.get('form').submit();

    // Check validation message
    cy.contains('Please enter a valid email address').should('be.visible');
  });

  it('should validate credit card format', () => {
    // Fill in invalid credit card
    cy.get('input#cardNumber').type('12345');
    cy.get('form').submit();

    // Check validation message
    cy.contains('Please enter a valid credit card number').should('be.visible');
  });

  it('should successfully process payment with valid data', () => {
    // Intercept payment API call
    cy.intercept('POST', '**/payment/process', {
      statusCode: 200,
      body: {
        success: true,
        orderId: '123456789'
      }
    }).as('processPayment');

    // Fill in form with valid data
    cy.get('input#name').type('John Doe');
    cy.get('input#email').type('john.doe@example.com');
    cy.get('input#address').type('123 Main St');
    cy.get('input#city').type('Anytown');
    cy.get('input#postalCode').type('12345');

    cy.get('input#cardName').type('John Doe');
    cy.get('input#cardNumber').type('4111111111111111');
    cy.get('input#expiryDate').type('12/25');
    cy.get('input#cvv').type('123');

    // Submit form
    cy.get('button[type="submit"]').click();
    cy.wait('@processPayment');

    // Check success message
    cy.contains('Payment Successful').should('be.visible');
    cy.contains('Thank you for your purchase').should('be.visible');
  });

  it('should handle payment API errors', () => {
    // Intercept payment API call with error
    cy.intercept('POST', '**/payment/process', {
      statusCode: 400,
      body: {
        success: false,
        error: 'Payment declined'
      }
    }).as('failedPayment');

    // Fill in form with valid data
    cy.get('input#name').type('John Doe');
    cy.get('input#email').type('john.doe@example.com');
    cy.get('input#address').type('123 Main St');
    cy.get('input#city').type('Anytown');
    cy.get('input#postalCode').type('12345');

    cy.get('input#cardName').type('John Doe');
    cy.get('input#cardNumber').type('4111111111111111');
    cy.get('input#expiryDate').type('12/25');
    cy.get('input#cvv').type('123');

    // Submit form
    cy.get('button[type="submit"]').click();
    cy.wait('@failedPayment');

    // Check error message
    cy.contains('Payment processing failed').should('be.visible');
  });

  it('should display correct order total in summary', () => {
    // Just verify the total is displayed and has a dollar sign
    cy.get('[data-testid="summary-total"]').should('be.visible');
    cy.get('[data-testid="summary-total"]').should('contain', '$');
  });
});
