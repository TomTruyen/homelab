#!/bin/bash

git config --global --add safe.directory /repo

REPO_URL="git@github.com:TomTruyen/homelab.git"
REPO_DIR="/repo"
DOCKER_COMPOSE_DIR="/repo/dashboard"

cd "$REPO_DIR" || exit 1

# Clone if missing
if [ ! -d ".git" ]; then
    echo "[GlanceFileSync] Cloning repo..."
    git clone "$REPO_URL" . || exit 1
    exit 0
fi

# Get list of changed files BEFORE pull
echo "[GlanceFileSync] Checking for incoming changes..."
git fetch origin main

CHANGED_FILES=$(git diff --name-only HEAD origin/main)

if [ -z "$CHANGED_FILES" ]; then
    echo "[GlanceFileSync] No changes detected."
    exit 0
fi

echo "[GlanceFileSync] Incoming changes:"
echo "$CHANGED_FILES"

# Now pull updates
git pull --rebase --autostash

SHOULD_RELOAD_ALL=false
SHOULD_RELOAD_ANILIST=false

# --- Trigger rules ---
while IFS= read -r FILE; do
    # 1. Full stack reload if .env.example changed
    if [[ "$FILE" == "dashboard/.env.example" ]]; then
        SHOULD_RELOAD_ALL=true
    fi

    # 2. Only restart api/anilist
    if [[ "$FILE" == "dashboard/api/anilist/"* ]]; then
        SHOULD_RELOAD_ANILIST=true
    fi
done <<< "$CHANGED_FILES"

# --- Execute actions ---
if [ "$SHOULD_RELOAD_ALL" = true ]; then
    echo "[GlanceFileSync] Reloading ALL services..."
    cd "$DOCKER_COMPOSE_DIR"
    docker compose down
    docker compose up -d --force-recreate
    exit 0
fi

if [ "$SHOULD_RELOAD_ANILIST" = true ]; then
    echo "[GlanceFileSync] Recreating anilist container..."
    cd "$DOCKER_COMPOSE_DIR"
    docker compose up -d --force-recreate --build anilist-api
    exit 0
fi
echo "[GlanceFileSync] Changes pulled, no service restarts required."
