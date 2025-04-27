#!/bin/bash
# API Testing Script for E-Shop API
# --------------------------------
# This script contains commands to test all endpoints of the E-Shop API
# Make sure to run it after starting the API server (docker-compose up)

API_URL="http://localhost:9000/api"

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}E-Shop API Testing Script${NC}"
echo "============================"
echo "Testing endpoints on $API_URL"
echo ""

# Function to print section headers
section() {
  echo -e "\n${YELLOW}$1${NC}"
  echo "--------------------------------------"
}

test_endpoint() {
  local method=$1
  local endpoint=$2
  local data=$3
  local description=$4
  
  echo -e "${GREEN}$description${NC}"
  echo "curl -X $method $API_URL$endpoint $data"
  
  # Run the actual curl command
  if [ -n "$data" ]; then
    curl -X $method "$API_URL$endpoint" -H "Content-Type: application/json" -d "$data" -s | json_pp
  else
    curl -X $method "$API_URL$endpoint" -s | json_pp
  fi
  
  echo -e "\n"
}

# ------------------------------
section "CATEGORY ENDPOINTS"

test_endpoint "GET" "/categories" "" "Get all categories"

# Create a category
test_endpoint "POST" "/categories" '{"name": "Electronics", "description": "Electronic devices and gadgets"}' "Create Electronics category"

# Create another category
test_endpoint "POST" "/categories" '{"name": "Clothing", "description": "Fashion and apparel"}' "Create Clothing category"

# Get category by ID
test_endpoint "GET" "/categories/1" "" "Get category with ID 1"

# Update category
test_endpoint "PUT" "/categories/1" '{"name": "Electronics and Gadgets", "description": "Latest electronic devices and accessories"}' "Update category with ID 1"

# ------------------------------
section "PRODUCT ENDPOINTS"

test_endpoint "GET" "/products" "" "Get all products"

# Create a product in Electronics category
test_endpoint "POST" "/products" '{"name": "iPhone 14", "description": "Latest Apple smartphone", "price": 999.99, "categoryId": 1}' "Create iPhone product"

# Create another product
test_endpoint "POST" "/products" '{"name": "Samsung Galaxy S23", "description": "Feature-rich Android phone", "price": 899.99, "categoryId": 1}' "Create Samsung Galaxy product"

# Create a clothing product
test_endpoint "POST" "/products" '{"name": "Denim Jacket", "description": "Stylish blue denim jacket", "price": 79.99, "categoryId": 2}' "Create clothing product"

# Get product by ID
test_endpoint "GET" "/products/1" "" "Get product with ID 1"

# Update product
test_endpoint "PUT" "/products/1" '{"name": "iPhone 14 Pro", "description": "Latest Apple smartphone with enhanced camera", "price": 1099.99, "categoryId": 1}' "Update product with ID 1"

# ------------------------------
section "CART ENDPOINTS"

test_endpoint "GET" "/carts" "" "Get all carts"

# Create a new empty cart
test_endpoint "POST" "/carts" '{"items": []}' "Create new empty cart"

# Add item to cart
test_endpoint "POST" "/carts/1/items" '{"productId": 1, "quantity": 2}' "Add iPhone to cart"

# Add another item to cart
test_endpoint "POST" "/carts/1/items" '{"productId": 3, "quantity": 1}' "Add Denim Jacket to cart"

# Get cart by ID
test_endpoint "GET" "/carts/1" "" "Get cart with ID 1"

# Update cart item quantity by updating the entire cart
test_endpoint "PUT" "/carts/1" '{"items": [{"id": 1, "productId": 1, "quantity": 3}, {"id": 2, "productId": 3, "quantity": 1}]}' "Update cart - change iPhone quantity to 3"

# Remove item from cart
test_endpoint "DELETE" "/carts/1/items/2" "" "Remove Denim Jacket from cart"

# Get updated cart
test_endpoint "GET" "/carts/1" "" "Get updated cart"

# ------------------------------
section "DELETE OPERATIONS"

# Delete a product
test_endpoint "DELETE" "/products/3" "" "Delete product with ID 3 (Denim Jacket)"

# Delete a category
test_endpoint "DELETE" "/categories/2" "" "Delete category with ID 2 (Clothing)"

# Delete a cart
test_endpoint "DELETE" "/carts/1" "" "Delete cart with ID 1"

echo -e "${BLUE}Testing complete!${NC}"