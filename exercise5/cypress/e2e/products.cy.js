/// <reference types="cypress" />

describe('Products Page', () => {
  beforeEach(() => {
    // Intercept API calls
    cy.intercept('GET', '**/products', { fixture: 'products.json' }).as('getProducts');
    cy.intercept('GET', '**/categories', { fixture: 'categories.json' }).as('getCategories');

    // Visit the home page
    cy.visit('/');
  });

  it('should display the products page correctly', () => {
    // Check heading
    cy.contains('h1', 'Products').should('be.visible');
    cy.get('.form-select').should('be.visible');
    cy.wait('@getProducts');
    cy.wait('@getCategories');
  });

  it('should display products from the API', () => {
    cy.wait('@getProducts');

    // Check that products are displayed
    cy.get('.card').should('have.length.at.least', 1);
    cy.get('.card-title').first().should('not.be.empty');
    cy.get('.card-text').first().should('not.be.empty');

    // Check price is displayed correctly
    cy.get('.card-text.fw-bold').first()
      .should('include.text', '$')
      .invoke('text')
      .then((text) => {
        const price = parseFloat(text.replace('$', ''));
        expect(price).to.be.a('number');
        expect(price).to.be.greaterThan(0);
      });
  });

  it('should filter products by category', () => {
    cy.wait('@getCategories');

    // Intercept category filter request
    cy.intercept('GET', '**/products/category/*', { fixture: 'filtered-products.json' }).as('getFilteredProducts');

    // Select a category
    cy.get('.form-select').select(1);
    cy.wait('@getFilteredProducts');

    // Check filtered products are displayed
    cy.get('.card').should('have.length.at.least', 1);
  });

  it('should show "No products found" when no products match the filter', () => {
    // Intercept with empty products
    cy.intercept('GET', '**/products/category/*', { body: [] }).as('getEmptyProducts');

    // Select a category
    cy.get('.form-select').select(1);
    cy.wait('@getEmptyProducts');

    // Check empty state message
    cy.contains('No products found').should('be.visible');
  });

  it('should add a product to the cart', () => {
    cy.wait('@getProducts');

    // Intercept cart add request
    cy.intercept('POST', '**/carts/*/items', { statusCode: 200 }).as('addToCart');

    // Click add to cart button on first product
    cy.get('.card button').first().click();

    // Check cart counter updates
    cy.get('[data-testid="cart-count"]').should('be.visible');
  });
});