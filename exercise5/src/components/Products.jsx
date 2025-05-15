import { useState, useEffect } from 'react';
import { productService, categoryService } from '../services/api';
import { useCart } from '../context/CartContext';
import { Card, Button, Container, Row, Col, Form } from 'react-bootstrap';

const Products = () => {
    const [products, setProducts] = useState([]);
    const [categories, setCategories] = useState([]);
    const [selectedCategory, setSelectedCategory] = useState('');
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const { addToCart, loading: cartLoading } = useCart();

    useEffect(() => {
        // Fetch products and categories on component mount
        const fetchData = async () => {
            try {
                setLoading(true);

                // Fetch all products
                const productsResponse = await productService.getAll();
                setProducts(productsResponse.data);

                // Fetch all categories
                const categoriesResponse = await categoryService.getAll();
                setCategories(categoriesResponse.data);

                setError(null);
            } catch (err) {
                console.error('Error fetching products or categories:', err);
                setError('Failed to load products. Please try again later.');
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    // Handle category selection
    const handleCategoryChange = async (e) => {
        const categoryId = e.target.value;
        setSelectedCategory(categoryId);

        try {
            setLoading(true);

            if (categoryId) {
                // Fetch products by category
                const response = await productService.getByCategory(categoryId);
                setProducts(response.data);
            } else {
                // Fetch all products if no category selected
                const response = await productService.getAll();
                setProducts(response.data);
            }

            setError(null);
        } catch (err) {
            console.error('Error fetching products by category:', err);
            setError('Failed to filter products. Please try again later.');
        } finally {
            setLoading(false);
        }
    };

    // Add product to cart
    const handleAddToCart = (product) => {
        addToCart(product, 1);
    };

    if (loading) {
        return <div className="text-center p-5">Loading products...</div>;
    }

    if (error) {
        return <div className="text-center p-5 text-danger">{error}</div>;
    }

    return (
        <Container fluid="md">
            <h1 className="my-4 text-center">Products</h1>

            {/* Category filter */}
            <Form.Group className="mb-4">
                <Form.Label>Filter by Category</Form.Label>
                <Form.Select
                    value={selectedCategory}
                    onChange={handleCategoryChange}
                >
                    <option value="">All Categories</option>
                    {categories.map(category => (
                        <option key={category.ID} value={category.ID}>
                            {category.name}
                        </option>
                    ))}
                </Form.Select>
            </Form.Group>

            {/* Product listing */}
            <Row className="g-4">
                {products.length === 0 ? (
                    <Col xs={12}>
                        <p className="text-center">No products found.</p>
                    </Col>
                ) : (
                    products.map(product => (
                        <Col key={product.ID} xs={12} sm={6} md={4} className="mb-4">
                            <Card className="h-100">
                                <Card.Body className="d-flex flex-column">
                                    <Card.Title className="mb-3">{product.name}</Card.Title>
                                    <Card.Text>{product.description}</Card.Text>
                                    <div className="mt-auto">
                                        <Card.Text className="fw-bold">${product.price.toFixed(2)}</Card.Text>
                                        <Card.Text className="text-muted">
                                            {product.quantity > 0
                                                ? `In stock: ${product.quantity}`
                                                : 'Out of stock'}
                                        </Card.Text>
                                        <Button
                                            variant="primary"
                                            onClick={() => handleAddToCart(product)}
                                            disabled={cartLoading || product.quantity <= 0}
                                            className="w-100"
                                        >
                                            {cartLoading ? 'Adding...' : 'Add to Cart'}
                                        </Button>
                                    </div>
                                </Card.Body>
                            </Card>
                        </Col>
                    ))
                )}
            </Row>
        </Container>
    );
};

export default Products;