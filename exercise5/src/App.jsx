import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { Container, Navbar, Nav, Badge } from 'react-bootstrap';
import { CartProvider, useCart } from './context/CartContext';
import Products from './components/Products';
import Cart from './components/Cart';
import Payment from './components/Payment';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

// Navigation component with cart badge
const Navigation = () => {
  const { itemCount } = useCart();

  return (
    <Navbar bg="dark" variant="dark" expand="lg" className="mb-4">
      <Container>
        <Navbar.Brand as={Link} to="/">E-Shop</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="ms-auto">
            <Nav.Link as={Link} to="/">Products</Nav.Link>
            <Nav.Link as={Link} to="/cart">
              Cart {itemCount > 0 && <Badge bg="danger" data-testid="cart-count">{itemCount}</Badge>}
            </Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
};

function App() {
  return (
    <Router>
      <CartProvider>
        <div className="App">
          <Navigation />
          <Routes>
            <Route path="/" element={<Products />} />
            <Route path="/cart" element={<Cart />} />
            <Route path="/payment" element={<Payment />} />
          </Routes>
        </div>
      </CartProvider>
    </Router>
  );
}

export default App;
