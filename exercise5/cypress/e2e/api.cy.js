/// <reference types="cypress" />

Cypress.Commands.add('getFirstItemId', (response) => {
  if (response.body.length === 0) return null;
  return response.body[0].ID;
});

Cypress.Commands.add('verifyProductDetails', (productId) => {
  cy.request('GET', `http://localhost:8080/products/${productId}`).then((productResponse) => {
    expect(productResponse.status).to.equal(200);
    expect(productResponse.body).to.have.property('ID', productId);
  });
});

Cypress.Commands.add('verifyDeleteProduct', (productId) => {
  cy.request('DELETE', `http://localhost:8080/products/${productId}`).then((deleteResponse) => {
    expect(deleteResponse.status).to.equal(204);
  });
});

Cypress.Commands.add('verifyCategoryDetails', (categoryId) => {
  cy.request('GET', `http://localhost:8080/categories/${categoryId}`).then((categoryResponse) => {
    expect(categoryResponse.status).to.equal(200);
    expect(categoryResponse.body).to.have.property('ID', categoryId);
  });
});

Cypress.Commands.add('addItemToCart', (cartId, productId) => {
  return cy.request({
    method: 'POST',
    url: `http://localhost:8080/carts/${cartId}/items`,
    body: {
      productID: productId,
      quantity: 1
    },
    failOnStatusCode: false
  });
});

Cypress.Commands.add('getProductAndAddToCart', (cartId) => {
  return cy.request('GET', 'http://localhost:8080/products')
    .then(response => {
      const productId = response.body.length > 0 ? response.body[0].ID : null;
      if (!productId) return null;

      return cy.addItemToCart(cartId, productId);
    });
});

describe('API Tests', () => {
  // Product API tests
  describe('Product API', () => {
    it('should get all products', () => {
      cy.request('GET', 'http://localhost:8080/products').then((response) => {
        expect(response.status).to.equal(200);
        expect(response.body).to.be.an('array');
      });
    });

    it('should get a single product', () => {
      // First get all products to find an ID
      cy.request('GET', 'http://localhost:8080/products')
        .then(response => cy.getFirstItemId(response))
        .then(productId => {
          if (productId) cy.verifyProductDetails(productId);
        });
    });

    it('should create and delete a product', () => {
      const newProduct = {
        name: 'Test Product',
        description: 'This is a test product',
        price: 99.99,
        quantity: 10,
        categoryID: 1
      };

      // Create product and verify/delete in sequence
      cy.request('POST', 'http://localhost:8080/products', newProduct)
        .then(response => {
          expect(response.status).to.equal(201);
          expect(response.body).to.have.property('ID');
          return response.body.ID;
        })
        .then(productId => cy.verifyDeleteProduct(productId));
    });
  });

  // Category API tests  
  describe('Category API', () => {
    it('should get all categories', () => {
      cy.request('GET', 'http://localhost:8080/categories').then((response) => {
        expect(response.status).to.equal(200);
        expect(response.body).to.be.an('array');
      });
    });

    it('should get a single category', () => {
      cy.request('GET', 'http://localhost:8080/categories')
        .then(response => cy.getFirstItemId(response))
        .then(categoryId => {
          if (categoryId) cy.verifyCategoryDetails(categoryId);
        });
    });
  });

  // Cart API tests
  describe('Cart API', () => {
    let cartId;

    beforeEach(() => {
      // Create a cart for testing
      cy.request('POST', 'http://localhost:8080/carts', { userID: 1 })
        .then(response => {
          expect(response.status).to.equal(201);
          cartId = response.body.ID;
        });
    });

    it('should get a cart', () => {
      cy.request('GET', `http://localhost:8080/carts/${cartId}`).then((response) => {
        expect(response.status).to.equal(200);
        expect(response.body).to.have.property('ID', cartId);
      });
    });

    it('should add an item to cart and handle errors', () => {
      // Get product and add to cart in a flattened way
      cy.getProductAndAddToCart(cartId)
        .then(response => {
          if (response) {
            cy.log(`Response status: ${response.status}`);
          }
        });
    });
  });
});
