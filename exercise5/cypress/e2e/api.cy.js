/// <reference types="cypress" />

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
      // First get all products to find an ID, then use an alias
      cy.request('GET', 'http://localhost:8080/products').as('productsRequest');
      cy.get('@productsRequest').then((response) => {
        if (response.body.length === 0) return;
        
        const productId = response.body[0].ID;
        // Store product ID as an alias
        cy.wrap(productId).as('productId');
      });
      
      // Use the stored ID in a separate command chain
      cy.get('@productId').then((productId) => {
        cy.request('GET', `http://localhost:8080/products/${productId}`).then((productResponse) => {
          expect(productResponse.status).to.equal(200);
          expect(productResponse.body).to.have.property('ID', productId);
        });
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

      // Create product and store the ID
      cy.request('POST', 'http://localhost:8080/products', newProduct).as('createProductRequest');
      cy.get('@createProductRequest').then((response) => {
        expect(response.status).to.equal(201);
        expect(response.body).to.have.property('ID');
        
        // Store product ID as an alias
        cy.wrap(response.body.ID).as('createdProductId');
      });
      
      // Delete product in a separate command chain
      cy.get('@createdProductId').then((productId) => {
        cy.request('DELETE', `http://localhost:8080/products/${productId}`).then((deleteResponse) => {
          expect(deleteResponse.status).to.equal(204);
        });
      });
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
      // First get all categories to find an ID
      cy.request('GET', 'http://localhost:8080/categories').as('categoriesRequest');
      cy.get('@categoriesRequest').then((response) => {
        if (response.body.length === 0) return;
        
        // Store category ID as an alias
        cy.wrap(response.body[0].ID).as('categoryId');
      });
      
      // Use the stored ID in a separate command chain
      cy.get('@categoryId').then((categoryId) => {
        cy.request('GET', `http://localhost:8080/categories/${categoryId}`).then((categoryResponse) => {
          expect(categoryResponse.status).to.equal(200);
          expect(categoryResponse.body).to.have.property('ID', categoryId);
        });
      });
    });
  });

  // Cart API tests
  describe('Cart API', () => {
    let cartId;

    beforeEach(() => {
      // Create a cart for testing
      cy.request('POST', 'http://localhost:8080/carts', { userID: 1 }).then((response) => {
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
      // Get products first and store the result
      cy.request('GET', 'http://localhost:8080/products').as('productsForCartRequest');
      cy.get('@productsForCartRequest').then((productsResponse) => {
        if (productsResponse.body.length === 0) return;
        
        // Store product ID as an alias
        cy.wrap(productsResponse.body[0].ID).as('productForCartId');
      });
      
      // Add item to cart using the stored product ID
      cy.get('@productForCartId').then((productId) => {
        cy.request({
          method: 'POST',
          url: `http://localhost:8080/carts/${cartId}/items`,
          body: {
            productID: productId,
            quantity: 1
          },
          failOnStatusCode: false
        }).then((response) => {
          cy.log(`Response status: ${response.status}`);
        });
      });
    });
  });
});
