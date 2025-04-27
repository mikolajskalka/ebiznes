#!/bin/bash

# Check if ngrok is installed
if ! command -v ngrok &> /dev/null; then
    echo "Error: ngrok is not installed. Please install it first."
    echo "Visit https://ngrok.com/download to download and install ngrok."
    exit 1
fi

# Check if container is running
if ! docker ps | grep -q e-shop-api; then
    echo "Starting e-shop-api container..."
    docker-compose up -d
fi

echo "Starting ngrok tunnel to the Play Framework application..."
echo "The public URL will be displayed below:"

# Start ngrok and forward to port 9000
ngrok http 9000