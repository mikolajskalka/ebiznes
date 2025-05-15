import axios from 'axios';

const API_URL = 'http://localhost:8080';

// Create axios instance with CORS headers
const apiClient = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        // CORS headers
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
        'Access-Control-Allow-Headers': 'Origin, Content-Type, Accept',
    }
});

// Product API services
export const productService = {
    getAll: () => apiClient.get('/products'),
    getById: (id) => apiClient.get(`/products/${id}`),
    create: (product) => apiClient.post('/products', product),
    update: (id, product) => apiClient.put(`/products/${id}`, product),
    delete: (id) => apiClient.delete(`/products/${id}`),
    getByCategory: (categoryId) => apiClient.get(`/products/category/${categoryId}`),
};

// Category API services
export const categoryService = {
    getAll: () => apiClient.get('/categories'),
    getById: (id) => apiClient.get(`/categories/${id}`),
    getAllWithProducts: () => apiClient.get('/categories/with-products'),
};

// Cart API services
export const cartService = {
    getAll: () => apiClient.get('/carts'),
    getById: (id) => apiClient.get(`/carts/${id}`),
    create: (cart) => apiClient.post('/carts', cart),
    addItem: (cartId, item) => apiClient.post(`/carts/${cartId}/items`, item),
    removeItem: (cartId, itemId) => apiClient.delete(`/carts/${cartId}/items/${itemId}`),
    getByUser: (userId) => apiClient.get(`/carts/user/${userId}`),
};

// Payment API service (mock since there's no actual payment endpoint in backend)
export const paymentService = {
    processPayment: (paymentData) => {
        // This would normally hit a real payment endpoint
        // For now, we'll just simulate a successful payment
        return new Promise((resolve) => {
            setTimeout(() => {
                resolve({
                    status: 'success',
                    message: 'Payment processed successfully',
                    data: paymentData
                });
            }, 800);
        });
    }
};