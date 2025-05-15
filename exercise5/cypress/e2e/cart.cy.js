/// <reference types="cypress" />

describe('Cart Functionality', () => {
  beforeEach(() => {
    // Intercept API calls
    cy.intercept('GET', '**/products', { fixture: 'products.json' }).as('getProducts');
    cy.intercept('GET', '**/categories', { fixture: 'categories.json' }).as('getCategories');
    cy.intercept('GET', '**/carts/**', { fixture: 'cart.json' }).as('getCart');
    cy.intercept('POST', '**/carts/*/items', { statusCode: 200 }).as('addToCart');

    // Visit the home page
    cy.visit('/');
    cy.wait('@getProducts');

    // Add a product to the cart
    cy.get('.card button').first().click();
  });

  it('should display the cart page correctly', () => {
    // Navigate to cart page
    cy.contains('Cart').click();
    cy.wait('@getCart');

    // Check heading and content
    cy.contains('h1', 'Your Cart').should('be.visible');
    cy.get('table').should('be.visible');
    cy.get('table tbody tr').should('have.length.at.least', 1);
  });

  it('should display correct product information in cart', () => {
    cy.contains('Cart').click();
    cy.wait('@getCart');

    // Check product details
    cy.get('table tbody tr').first().within(() => {
      cy.get('td').eq(0).should('not.be.empty'); // Product name
      cy.get('td').eq(1).invoke('text').should('match', /\$\d+\.\d{2}/); // Price format
      cy.get('td').eq(2).find('input').should('have.value', '1'); // Quantity default
    });
  });

  it('should update quantity when + button is clicked', () => {
    cy.contains('Cart').click();
    cy.wait('@getCart');

    // Intercept quantity update
    cy.intercept('PUT', '**/carts/*/items/*', { statusCode: 200 }).as('updateQuantity');

    // Get initial quantity
    cy.get('table tbody tr').first().find('input').invoke('val').then((initialVal) => {
      const initialQuantity = parseInt(initialVal);

      // Click + button
      cy.get('table tbody tr').first().find('button').contains('+').click();
      cy.wait('@updateQuantity');

      // Check new quantity
      cy.get('table tbody tr').first().find('input')
        .invoke('val')
        .should('eq', (initialQuantity + 1).toString());
    });
  });

  it('should update quantity when - button is clicked', () => {
    cy.contains('Cart').click();
    cy.wait('@getCart');

    // First add quantity to make sure we can decrease
    cy.get('table tbody tr').first().find('button').contains('+').click();

    // Intercept quantity update
    cy.intercept('PUT', '**/carts/*/items/*', { statusCode: 200 }).as('updateQuantity');

    // Get current quantity
    cy.get('table tbody tr').first().find('input').invoke('val').then((initialVal) => {
      const initialQuantity = parseInt(initialVal);

      // Click - button
      cy.get('table tbody tr').first().find('button').contains('-').click();
      cy.wait('@updateQuantity');

      // Check new quantity
      cy.get('table tbody tr').first().find('input')
        .invoke('val')
        .should('eq', (initialQuantity - 1).toString());
    });
  });

  it('should remove item from cart', () => {
    cy.contains('Cart').click();
    cy.wait('@getCart');

    // Count initial items
    cy.get('table tbody tr').then(($rows) => {
      const initialCount = $rows.length;

      // Intercept remove item
      cy.intercept('DELETE', '**/carts/*/items/*', { statusCode: 200 }).as('removeItem');

      // Click remove button
      cy.get('table tbody tr').first().find('button').contains('Remove').click();
      cy.wait('@removeItem');

      // Check item was removed
      cy.get('table tbody tr').should('have.length', initialCount - 1);
    });
  });

  it('should calculate and display correct subtotal', () => {
    cy.contains('Cart').click();
    cy.wait('@getCart');

    // Calculate expected total
    let expectedTotal = 0;

    cy.get('table tbody tr').each(($row) => {
      const price = parseFloat($row.find('td').eq(1).text().replace('$', ''));
      const quantity = parseInt($row.find('td').eq(2).find('input').val());
      expectedTotal += price * quantity;
    }).then(() => {
      // Check displayed total
      cy.get('[data-testid="cart-total"]')
        .invoke('text')
        .should('include', `$${expectedTotal.toFixed(2)}`);
    });
  });

  it('should proceed to checkout', () => {
    cy.contains('Cart').click();
    cy.wait('@getCart');

    cy.contains('button', 'Proceed to Checkout').click();
    cy.url().should('include', '/payment');
  });
});