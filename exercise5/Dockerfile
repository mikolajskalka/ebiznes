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

# Install all dependencies with --ignore-scripts for security
RUN npm install --ignore-scripts && npm install --ignore-scripts react-bootstrap bootstrap

# Copy only necessary files for the application
COPY --chown=nodeuser:nodegroup vite.config.js index.html cypress.config.js eslint.config.js ./
COPY --chown=nodeuser:nodegroup public/ ./public/
COPY --chown=nodeuser:nodegroup src/ ./src/
# Explicitly exclude test files, logs and other non-essential files

# Set read-only permissions for all copied files to ensure security
RUN find /app -type f -exec chmod 444 {} \; && \
    find /app -type d -exec chmod 555 {} \;

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]