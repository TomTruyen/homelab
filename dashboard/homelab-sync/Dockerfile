FROM alpine:latest

RUN apk add --no-cache git bash curl openssh

RUN addgroup -g 1000 appuser \
    && adduser -D -u 1000 -G appuser appuser

# Directory where repo will be cloned
WORKDIR /repo

# Add sync script
COPY sync.sh /sync.sh
RUN chmod +x /sync.sh

USER appuser

# Run update script every minute using a loop (simpler than cron in Docker)
CMD ["/bin/bash", "-c", "while true; do /sync.sh; sleep 60; done"]
