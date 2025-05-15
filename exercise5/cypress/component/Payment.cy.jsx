/// <reference types="cypress" />
import React from 'react';
import { mount } from '@cypress/react';
import Payment from '../../src/components/Payment';
import { CartProvider } from '../../src/context/CartContext';

describe('Payment Component Unit Tests', () => {
  beforeEach(() => {
    // Mock API responses
    cy.intercept('GET', '**/carts/**', { fixture: 'cart.json' }).as('getCart');
    cy.intercept('POST', '**/payment/process', {
      statusCode: 200,
      body: {
        success: true,
        orderId: '123456789'
      }
    }).as('processPayment');

    // Mount the component with the CartProvider
    cy.mount(
      <CartProvider>
        <Payment />
      </CartProvider>
    );

    // Wait for initial data loading
    cy.wait('@getCart');
  });

  it('should render the payment form correctly', () => {
    // Check page structure
    cy.get('h1').should('contain', 'Checkout');
    cy.contains('h2', 'Billing Information').should('be.visible');
    cy.contains('h2', 'Payment Details').should('be.visible');
    cy.contains('h2', 'Order Summary').should('be.visible');

    // Check form fields
    cy.get('form').within(() => {
      // Billing information fields
      cy.get('input#name').should('exist');
      cy.get('input#email').should('exist');
      cy.get('input#address').should('exist');
      cy.get('input#city').should('exist');
      cy.get('input#postalCode').should('exist');

      // Payment details fields
      cy.get('input#cardName').should('exist');
      cy.get('input#cardNumber').should('exist');
      cy.get('input#expiryDate').should('exist');
      cy.get('input#cvv').should('exist');

      // Submit button
      cy.get('button[type="submit"]').should('exist').and('contain', 'Complete Purchase');
    });
  });

  it('should display correct order summary', () => {
    // Check order summary
    cy.fixture('cart').then((cart) => {
      // Calculate expected total
      let expectedTotal = 0;
      cart.items.forEach(item => {
        expectedTotal += item.product.price * item.quantity;
      });

      // Check items in summary
      cy.get('.order-summary').within(() => {
        cart.items.forEach((item) => {
          cy.contains(item.product.name).should('be.visible');
          cy.contains('$' + (item.product.price * item.quantity).toFixed(2)).should('be.visible');
        });

        // Check total
        cy.get('[data-testid="summary-total"]').should('contain', '$' + expectedTotal.toFixed(2));
      });
    });
  });

  it('should validate required fields', () => {
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
    // Fill invalid email
    cy.get('input#email').type('invalid-email');

    // Submit form
    cy.get('form').submit();

    // Check validation message
    cy.contains('Please enter a valid email address').should('be.visible');
  });

  it('should validate credit card number format', () => {
    // Fill invalid card number
    cy.get('input#cardNumber').type('1234');

    // Submit form
    cy.get('form').submit();

    // Check validation message
    cy.contains('Please enter a valid credit card number').should('be.visible');
  });

  it('should validate CVV format', () => {
    // Fill invalid CVV
    cy.get('input#cvv').type('a');

    // Submit form
    cy.get('form').submit();

    // Check validation message
    cy.contains('CVV must be 3 or 4 digits').should('be.visible');
  });

  it('should validate expiry date format', () => {
    // Fill invalid expiry date
    cy.get('input#expiryDate').type('99/99');

    // Submit form
    cy.get('form').submit();

    // Check validation message
    cy.contains('Please enter a valid expiry date (MM/YY)').should('be.visible');
  });

  it('should successfully submit with valid data', () => {
    // Fill form with valid data
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
    cy.contains('Order ID: 123456789').should('be.visible');
    cy.contains('Thank you for your purchase!').should('be.visible');
  });

  it('should handle payment processing error', () => {
    // Override payment intercept with error response
    cy.intercept('POST', '**/payment/process', {
      statusCode: 400,
      body: {
        success: false,
        error: 'Payment declined'
      }
    }).as('failedPayment');

    // Fill form with valid data
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
    cy.contains('Payment Failed').should('be.visible');
    cy.contains('Payment declined').should('be.visible');
  });
});