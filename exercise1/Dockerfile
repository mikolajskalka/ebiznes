FROM ubuntu:24.04

# Set noninteractive installation
ENV DEBIAN_FRONTEND=noninteractive

# Update and install basic utilities
RUN apt-get update && apt-get install -y \
    wget \
    curl \
    software-properties-common \
    unzip \
    zip \
    && rm -rf /var/lib/apt/lists/*

# Install Java 8
RUN apt-get update && \
    apt-get install -y openjdk-8-jdk && \
    rm -rf /var/lib/apt/lists/*

# Install Python 3.10
RUN apt-get update && \
    apt-get install -y python3.10 python3-pip && \
    rm -rf /var/lib/apt/lists/*

# Install Kotlin using SDKMAN
RUN curl -s "https://get.sdkman.io" | bash && \
    bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && \
    sdk install kotlin && \
    echo 'source $HOME/.sdkman/bin/sdkman-init.sh' >> $HOME/.bashrc"

# Add Kotlin to PATH
ENV PATH=$PATH:/root/.sdkman/candidates/kotlin/current/bin


# Install Gradle
RUN bash -c "source $HOME/.sdkman/bin/sdkman-init.sh && \
    sdk install gradle 4.10.3" 

# Add Gradle to PATH
ENV PATH=$PATH:/root/.sdkman/candidates/gradle/current/bin


# Verify installations and Java home
RUN java -version && \
    kotlin -version && \
    python3 --version && \
    gradle --version 

# Set working directory
WORKDIR /app

# Copy the project files
COPY . /app

# Set the entry point to run the application with Gradle
ENTRYPOINT ["gradle", "run"]
