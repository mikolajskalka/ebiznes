/// <reference types="cypress" />
import React from 'react';
import { mount } from '@cypress/react';
import Cart from '../../src/components/Cart';
import { CartProvider } from '../../src/context/CartContext';

describe('Cart Component Unit Tests', () => {
  beforeEach(() => {
    // Mock the API responses
    cy.intercept('GET', '**/carts/**', { fixture: 'cart.json' }).as('getCart');
    cy.intercept('POST', '**/carts/*/items', { statusCode: 200 }).as('addToCart');
    cy.intercept('PUT', '**/carts/*/items/*', { statusCode: 200 }).as('updateCartItem');
    cy.intercept('DELETE', '**/carts/*/items/*', { statusCode: 200 }).as('removeCartItem');

    // Mount the component with the CartProvider
    cy.mount(
      <CartProvider>
        <Cart />
      </CartProvider>
    );

    // Wait for initial data loading
    cy.wait('@getCart');
  });

  it('should render the cart correctly', () => {
    // Check that the component renders
    cy.get('h1').should('contain', 'Your Cart');

    // Check table structure
    cy.get('table').should('exist');
    cy.get('thead').should('exist');
    cy.get('tbody').should('exist');

    // Check column headers
    cy.get('thead th').eq(0).should('contain', 'Product');
    cy.get('thead th').eq(1).should('contain', 'Price');
    cy.get('thead th').eq(2).should('contain', 'Quantity');
    cy.get('thead th').eq(3).should('contain', 'Subtotal');
    cy.get('thead th').eq(4).should('contain', 'Actions');

    // Check cart items are rendered
    cy.get('tbody tr').should('have.length.at.least', 1);
  });

  it('should display correct cart information', () => {
    // Test the cart items in detail
    cy.fixture('cart').then((cart) => {
      const firstItem = cart.items[0];

      cy.get('tbody tr').first().within(() => {
        // Check product name
        cy.get('td').eq(0).should('contain', firstItem.product.name);

        // Check price format and value
        cy.get('td').eq(1).should('contain', '$' + firstItem.product.price.toFixed(2));

        // Check quantity input
        cy.get('td').eq(2).find('input').should('have.value', firstItem.quantity.toString());

        // Check subtotal calculation
        const expectedSubtotal = (firstItem.product.price * firstItem.quantity).toFixed(2);
        cy.get('td').eq(3).should('contain', '$' + expectedSubtotal);

        // Check remove button exists
        cy.get('td').eq(4).find('button').should('contain', 'Remove');
      });
    });
  });

  it('should calculate total price correctly', () => {
    cy.fixture('cart').then((cart) => {
      // Calculate expected total manually
      let expectedTotal = 0;
      cart.items.forEach(item => {
        expectedTotal += item.product.price * item.quantity;
      });

      // Check total is displayed correctly
      cy.get('[data-testid="cart-total"]').should('contain', '$' + expectedTotal.toFixed(2));
    });
  });

  it('should increase quantity when + button is clicked', () => {
    // Get initial quantity
    cy.get('tbody tr').first().find('input').invoke('val').then((initialVal) => {
      const initialQuantity = parseInt(initialVal);

      // Click + button
      cy.get('tbody tr').first().find('button').contains('+').click();
      cy.wait('@updateCartItem');

      // Check new quantity
      cy.get('tbody tr').first().find('input')
        .invoke('val')
        .should('eq', (initialQuantity + 1).toString());
    });
  });

  it('should decrease quantity when - button is clicked', () => {
    // First increase quantity to ensure we can decrease it
    cy.get('tbody tr').first().find('button').contains('+').click();
    cy.wait('@updateCartItem');

    // Get current quantity
    cy.get('tbody tr').first().find('input').invoke('val').then((initialVal) => {
      const initialQuantity = parseInt(initialVal);

      // Click - button
      cy.get('tbody tr').first().find('button').contains('-').click();
      cy.wait('@updateCartItem');

      // Check new quantity
      cy.get('tbody tr').first().find('input')
        .invoke('val')
        .should('eq', (initialQuantity - 1).toString());
    });
  });

  it('should remove item when Remove button is clicked', () => {
    // Count initial items
    cy.get('tbody tr').then($rows => {
      const initialCount = $rows.length;

      // Click remove button on first item
      cy.get('tbody tr').first().find('button').contains('Remove').click();
      cy.wait('@removeCartItem');

      // Check item was removed
      cy.get('tbody tr').should('have.length', initialCount - 1);
    });
  });

  it('should handle empty cart state', () => {
    // Intercept with empty cart
    cy.intercept('GET', '**/carts/**', {
      body: {
        ID: 1,
        userID: 1,
        items: []
      }
    }).as('getEmptyCart');

    // Remount with empty cart
    cy.mount(
      <CartProvider>
        <Cart />
      </CartProvider>
    );
    cy.wait('@getEmptyCart');

    // Check empty state message
    cy.contains('Your cart is empty').should('be.visible');
    cy.get('tbody tr').should('not.exist');

    // Checkout button should be disabled
    cy.get('button').contains('Proceed to Checkout').should('be.disabled');
  });

  it('should proceed to checkout when button clicked', () => {
    // Spy on navigation
    const navigateSpy = cy.spy().as('navigateSpy');

    // Mount with navigation mock
    cy.mount(
      <CartProvider>
        <Cart navigate={navigateSpy} />
      </CartProvider>
    );
    cy.wait('@getCart');

    // Click checkout button
    cy.get('button').contains('Proceed to Checkout').click();

    // Check navigation was called with correct path
    cy.get('@navigateSpy').should('have.been.calledWith', '/payment');
  });
});