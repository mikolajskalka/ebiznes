// ***********************************************************
// This support file is processed and loaded automatically before your test files.
// This is a great place to put global configuration and behavior that modifies Cypress.
// ***********************************************************

// Import commands.js using ES2015 syntax:
import './commands'

// Import global styles
import 'bootstrap/dist/css/bootstrap.min.css'
import '../../src/index.css'

// Alternatively you can use CommonJS syntax:
// require('./commands')

import { mount } from 'cypress/react'

Cypress.Commands.add('mount', mount)

// Example use:
// cy.mount(<MyComponent />)