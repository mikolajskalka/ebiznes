FROM ubuntu:24.04

# Set noninteractive installation
ENV DEBIAN_FRONTEND=noninteractive

# Update and install basic utilities
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    software-properties-common \
    unzip \
    wget \
    zip \
    && rm -rf /var/lib/apt/lists/*

# Install Java 8 and Python 3.10
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        openjdk-8-jdk \
        python3.10 \
        python3-pip \
    && rm -rf /var/lib/apt/lists/*

# Install Kotlin and Gradle using SDKMAN
RUN curl -s "https://get.sdkman.io" | bash && \
    bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && \
    sdk install kotlin && \
    sdk install gradle 4.10.3 && \
    echo 'source $HOME/.sdkman/bin/sdkman-init.sh' >> $HOME/.bashrc"

# Add Kotlin and Gradle to PATH
ENV PATH=$PATH:/root/.sdkman/candidates/kotlin/current/bin:/root/.sdkman/candidates/gradle/current/bin


# Verify installations and Java home
RUN java -version && \
    kotlin -version && \
    python3 --version && \
    gradle --version 

# Set working directory
WORKDIR /app

# Create a non-root user
RUN groupadd -r appuser && useradd -r -g appuser -m appuser

# Copy only necessary project files
COPY build.gradle /app/
COPY src/ /app/src/

# Set proper permissions
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Set the entry point to run the application with Gradle
ENTRYPOINT ["gradle", "run"]
