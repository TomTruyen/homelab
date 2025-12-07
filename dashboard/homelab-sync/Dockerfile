FROM alpine:latest

RUN apk add --no-cache git bash curl

# Directory where repo will be cloned
WORKDIR /repo

# Add sync script
COPY sync.sh /sync.sh
RUN chmod +x /sync.sh

# Run update script every minute using a loop (simpler than cron in Docker)
CMD ["/bin/bash", "-c", "while true; do /sync.sh; sleep 60; done"]
