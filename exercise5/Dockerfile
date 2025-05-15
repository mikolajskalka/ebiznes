FROM node:23-alpine AS build

WORKDIR /app

# Copy package files first for better layer caching
COPY package*.json ./

# Install all dependencies, including react-bootstrap
RUN npm install && npm install react-bootstrap bootstrap

# Copy the rest of the application code
COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]