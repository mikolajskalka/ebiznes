# E-Business Course Solutions

This repository contains solutions for exercises completed as part of the E-Business university course.

## Exercise 1: Containerized Development Environment

The first exercise demonstrates a containerized development environment that includes:

- Ubuntu 24.04 base image
- Java 8 (OpenJDK)
- Python 3.10
- Kotlin (installed via SDKMAN)
- Gradle 4.10.3

The container runs a simple Java "Hello World" application.

### Running the Solution

To build and run the containerized application:

```bash
cd exercise1
docker compose up
```

### Docker Image

The Docker image for this solution is available on Docker Hub:
[mikolajskalka/java-hello-world-app:latest](https://hub.docker.com/repository/docker/mikolajskalka/java-hello-world-app/tags/latest/sha256-c5824510a94d5fdeedd1904e5ef0124b06fbc82af781cc287afa69949da041b3)

You can pull the image directly with:

```bash
docker pull mikolajskalka/java-hello-world-app:latest