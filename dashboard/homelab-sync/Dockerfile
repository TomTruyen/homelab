# Use Docker CLI image (includes docker and docker compose plugin)
FROM docker:24.0.5-cli

# Install required packages
RUN apk add --no-cache git bash curl

# Prepare SSH for GitHub
RUN mkdir -p /root/.ssh && \
    ssh-keyscan github.com >> /root/.ssh/known_hosts

# Set working directory where the repo will be mounted
WORKDIR /repo

# Add sync script
COPY sync.sh /sync.sh
RUN chmod +x /sync.sh

# Run update script every minute
CMD ["/bin/bash", "-c", "while true; do /sync.sh; sleep 60; done"]
