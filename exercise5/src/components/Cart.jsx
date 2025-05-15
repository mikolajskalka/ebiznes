import { useNavigate } from 'react-router-dom';
import { useCart } from '../context/CartContext';
import { Container, Table, Button, Alert } from 'react-bootstrap';

const Cart = () => {
    const { items, totalPrice, loading, removeFromCart, updateQuantity } = useCart();
    const navigate = useNavigate();

    const handleRemoveItem = (itemId) => {
        removeFromCart(itemId);
    };

    const updateCartItemQuantity = (itemId, newQuantity) => {
        if (newQuantity > 0) {
            updateQuantity(itemId, newQuantity);
        }
    };

    const handleCheckout = () => {
        navigate('/payment');
    };

    if (loading) {
        return <div className="text-center p-5">Loading cart...</div>;
    }

    return (
        <Container>
            <h1 className="my-4">Your Cart</h1>

            {items.length === 0 ? (
                <Alert variant="info">
                    Your cart is empty. <Button variant="link" onClick={() => navigate('/')}>Browse products</Button>
                </Alert>
            ) : (
                <>
                    <Table striped bordered hover>
                        <thead>
                            <tr>
                                <th>Product</th>
                                <th>Price</th>
                                <th>Quantity</th>
                                <th>Subtotal</th>
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {items.map((item, index) => (
                                <tr key={item.ID || index}>
                                    <td>{item.product ? item.product.name : `Product #${item.product_id}`}</td>
                                    <td>${item.price ? item.price.toFixed(2) : '0.00'}</td>
                                    <td>
                                        <div className="d-flex align-items-center">
                                            <Button
                                                variant="outline-secondary"
                                                size="sm"
                                                onClick={() => updateCartItemQuantity(item.ID, item.quantity - 1)}
                                                disabled={item.quantity <= 1}
                                            >
                                                -
                                            </Button>
                                            <input
                                                type="number"
                                                className="form-control mx-2"
                                                style={{ width: '60px' }}
                                                value={item.quantity}
                                                readOnly
                                            />
                                            <Button
                                                variant="outline-secondary"
                                                size="sm"
                                                onClick={() => updateCartItemQuantity(item.ID, item.quantity + 1)}
                                            >
                                                +
                                            </Button>
                                        </div>
                                    </td>
                                    <td>${(item.price * item.quantity).toFixed(2)}</td>
                                    <td>
                                        <Button
                                            variant="danger"
                                            size="sm"
                                            onClick={() => handleRemoveItem(item.ID)}
                                        >
                                            Remove
                                        </Button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                        <tfoot>
                            <tr>
                                <td colSpan="4" className="text-end fw-bold">Total:</td>
                                <td colSpan="2" className="fw-bold" data-testid="cart-total">${totalPrice.toFixed(2)}</td>
                            </tr>
                        </tfoot>
                    </Table>

                    <div className="d-flex justify-content-between my-4">
                        <Button variant="secondary" onClick={() => navigate('/')}>
                            Continue Shopping
                        </Button>
                        <Button variant="success" onClick={handleCheckout}>
                            Proceed to Checkout
                        </Button>
                    </div>
                </>
            )}
        </Container>
    );
};

export default Cart;