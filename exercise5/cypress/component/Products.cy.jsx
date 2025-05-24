/// <reference types="cypress" />
import React from 'react';
import Products from '../../src/components/Products';
import { CartProvider } from '../../src/context/CartContext';

describe('Products Component Unit Tests', () => {
  beforeEach(() => {
    // Mock the API responses
    cy.intercept('GET', '**/products', { fixture: 'products.json' }).as('getProducts');
    cy.intercept('GET', '**/categories', { fixture: 'categories.json' }).as('getCategories');
    cy.intercept('POST', '**/carts/*/items', { statusCode: 200 }).as('addToCart');

    // Mount the component with the CartProvider
    cy.mount(
      <CartProvider>
        <Products />
      </CartProvider>
    );

    // Wait for initial data loading
    cy.wait('@getProducts');
    cy.wait('@getCategories');
  });

  it('should render the products correctly', () => {
    // Check that the component renders
    cy.get('h1').should('contain', 'Products');

    // Check form select for categories
    cy.get('.form-select').should('exist');
    cy.get('.form-select option').should('have.length.at.least', 2); // At least "All Categories" + 1 category

    // Check products are rendered
    cy.get('.card').should('have.length.at.least', 1);

    // Check structure of each product card
    cy.get('.card').first().within(() => {
      cy.get('.card-title').should('exist');
      cy.get('.card-text').should('exist');
      cy.get('.card-text.fw-bold').should('exist').and('contain', '$');
      cy.get('button').should('exist');
    });
  });

  it('should display correct product information', () => {
    // Test the first product in detail
    cy.fixture('products').then((products) => {
      const firstProduct = products[0];

      cy.get('.card').first().within(() => {
        // Check product name matches fixture
        cy.get('.card-title').should('contain', firstProduct.name);

        // Check product description matches fixture
        cy.get('.card-text').first().should('contain', firstProduct.description);

        // Check price format and value
        cy.get('.card-text.fw-bold').should('contain', '$' + firstProduct.price.toFixed(2));

        // Check stock status
        cy.get('.card-text.text-muted').should('contain', 'In stock: ' + firstProduct.quantity);

        // Check button is enabled when in stock
        cy.get('button').should('not.be.disabled').and('contain', 'Add to Cart');
      });
    });
  });

  it('should allow filtering products by category', () => {
    // Intercept category filter request
    cy.intercept('GET', '**/products/category/*', { fixture: 'filtered-products.json' }).as('getFilteredProducts');

    // Select first category
    cy.get('.form-select').select(1);
    cy.wait('@getFilteredProducts');

    // Check that filtered products are displayed
    cy.fixture('filtered-products').then((filteredProducts) => {
      cy.get('.card').should('have.length', filteredProducts.length);

      // Check that the first filtered product is displayed correctly
      cy.get('.card').first().within(() => {
        cy.get('.card-title').should('contain', filteredProducts[0].name);
        cy.get('.card-text').first().should('contain', filteredProducts[0].description);
      });
    });
  });

  it('should show disabled button for out-of-stock products', () => {
    // Create a custom intercept for a product with zero quantity
    cy.intercept('GET', '**/products', (req) => {
      req.reply((res) => {
        const products = res.body;
        products[0].quantity = 0;
        res.send(products);
      });
    }).as('getModifiedProducts');

    // Remount the component to use the modified data
    cy.mount(
      <CartProvider>
        <Products />
      </CartProvider>
    );
    cy.wait('@getModifiedProducts');

    // Check first product's button is disabled
    cy.get('.card').first().within(() => {
      cy.get('.text-muted').should('contain', 'Out of stock');
      cy.get('button').should('be.disabled');
    });
  });

  it('should handle adding products to cart', () => {
    // Click the "Add to Cart" button on the first product
    cy.get('.card').first().find('button').click();
    cy.wait('@addToCart');

    // Since we're using CartContext, check that the button state changes temporarily
    cy.get('.card').first().find('button').should('contain', 'Add to Cart');

    // Try adding multiple products to cart
    cy.get('.card').eq(1).find('button').click();
    cy.wait('@addToCart');

    cy.get('.card').eq(2).find('button').click();
    cy.wait('@addToCart');
  });

  it('should handle error during product loading', () => {
    // Create a custom intercept for error
    cy.intercept('GET', '**/products', {
      statusCode: 500,
      body: 'Server error'
    }).as('getProductsError');

    // Remount the component to trigger the error
    cy.mount(
      <CartProvider>
        <Products />
      </CartProvider>
    );
    cy.wait('@getProductsError');

    // Check error state
    cy.contains('Failed to load products').should('be.visible');
  });
});