FROM node:23-alpine AS build

WORKDIR /app

# Create a non-root user
RUN addgroup -S nodegroup && adduser -S -G nodegroup nodeuser

# Copy package files first for better layer caching
COPY package*.json ./

# Set proper ownership for npm cache directory to avoid permission issues
RUN mkdir -p /home/nodeuser/.npm && \
    chown -R nodeuser:nodegroup /home/nodeuser/.npm && \
    chown -R nodeuser:nodegroup /app

# Switch to non-root user for package installation
USER nodeuser

# Install all dependencies, including react-bootstrap
RUN npm install && npm install react-bootstrap bootstrap

# Copy only necessary files for the application
COPY --chown=nodeuser:nodegroup vite.config.js index.html cypress.config.js eslint.config.js ./
COPY --chown=nodeuser:nodegroup public/ ./public/
COPY --chown=nodeuser:nodegroup src/ ./src/
# Explicitly exclude test files, logs and other non-essential files

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]