import { defineConfig } from "cypress";

export default defineConfig({
  e2e: {
    baseUrl: "http://localhost:5173",
    setupNodeEvents(on, config) {
      // implement node event listeners here
    },
  },

  env: {
    apiUrl: "http://localhost:8080",
  },

  viewportWidth: 1280,
  viewportHeight: 720,
  video: true,
  screenshotOnRunFailure: true,

  // Configuration for Browserstack
  projectId: "exercise5",

  retries: {
    runMode: 1,
    openMode: 0,
  },

  component: {
    devServer: {
      framework: "react",
      bundler: "vite",
    },
  },
});
