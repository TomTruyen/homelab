# Homelab Sync

Automated synchronization service that keeps your homelab dashboard in sync with your GitHub repository. When changes are pushed to your repo, this service automatically pulls them and restarts only the affected services.

## What It Does

The Homelab Sync service runs as a Docker container and:

1. **Monitors your repository** for changes from GitHub
2. **Automatically pulls updates** from your remote repository
3. **Intelligently restarts services** based on what changed:
   - Full stack reload if environment variables change
   - Individual service restart if only specific service code changes
   - No restart if only documentation changes

This eliminates manual SSH sessions to restart services after pushing updates to your homelab configuration.

## Features

- **Automatic Git Sync**: Periodically fetches and pulls changes from your GitHub repository
- **Smart Service Restart**: Only restarts containers that need updating
- **Safe Git Operations**: Handles merge conflicts gracefully with `--autostash`
- **Partial Reloads**: Restarts only affected services (e.g., just the AniList API) for faster updates

## Trigger Rules

The service automatically restarts containers when:

| File Changed | Action |
|--------------|--------|
| `dashboard/.env.example` | Reloads **all services** (full stack restart) |
| `dashboard/api/anilist/*` | Restarts only the **AniList API container** |
| Other files | No automatic restart |

You can add additional trigger rules in `sync.sh` by modifying the conditions.

## Setup

### 1. SSH Key Configuration

The service needs SSH access to your GitHub repository. Ensure your SSH key is properly configured:

```bash
# Copy your SSH key into the container (adjust path as needed)
# This is typically handled via Docker volume mounts in docker-compose.yml
```

### 2. Docker Compose

The service runs as a Docker container. Ensure your `docker-compose.yml` includes:

```yaml
services:
  homelab-sync:
    build: ./homelab-sync
    container_name: glance-homelab-sync
    user: root
    volumes:
      - .././:/repo # Mount the homelab repo so that the container can access it
      - ~/.ssh:/root/.ssh:ro # Mount SSH keys for git access
      - /var/run/docker.sock:/var/run/docker.sock # Access Docker socket to restart services
    restart: unless-stopped
```

### 3. Start the Service

```bash
cd dashboard && docker compose up -d homelab-sync
```

## How It Works

1. **Initial Setup**: If the repo isn't cloned, the service clones it from GitHub
2. **Periodic Check**: Every few minutes (configurable), the service:
   - Fetches latest changes from the remote repository
   - Compares local state with remote state
   - Identifies which files changed
3. **Smart Restart**: Based on trigger rules, restarts only necessary services
4. **Logging**: All sync operations are logged with `[GlanceFileSync]` prefix

## Configuration

Edit `sync.sh` to customize:

- **Repository URL**: Change the `REPO_URL` variable
- **Sync Interval**: Set via environment variable or cron schedule
- **Trigger Rules**: Add new conditions for different file paths
- **Restart Commands**: Modify docker compose commands as needed

## Troubleshooting

### Service won't sync
- Verify SSH key is accessible inside the container
- Check Docker socket is mounted correctly
- Review logs: `docker logs <homelab-sync-container-name>`

### Services not restarting
- Check that file paths in trigger rules match your repository structure
- Ensure docker compose is accessible from the container
- Verify the container has permission to run docker commands

### Manual Sync
If needed, manually trigger a sync:

```bash
docker exec <homelab-sync-container-name> /app/sync.sh
```
