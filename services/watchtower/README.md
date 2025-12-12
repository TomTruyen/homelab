# Watchtower

Automated container update service that watches running Docker containers and automatically pulls and restarts them when a new image is available.

Watchtower (containrrr/watchtower) is a simple helper service to keep your homelab containers up-to-date with minimal manual intervention.

🔗 https://github.com/containrrr/watchtower

------------------------------------------------------------------------

## 🚀 Purpose

- Automatically checks for new images for containers running on the host
- Pulls updated images and recreates containers with the new image
- Helps keep images current without manually running docker-compose pull / up

## 🧩 Docker Compose (example)

This repository runs Watchtower with the Docker socket mounted so it can manage local containers:

```yaml
services:
  watchtower:
    image: containrrr/watchtower
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    restart: unless-stopped
    env_file: .env
    # Example: configure env in `services/watchtower/.env` or copy from `.env.example`
```

Notes:
- Mounting `/var/run/docker.sock` grants Watchtower control over Docker on the host; treat this privileged access carefully.
- Consider using the `--interval` flag to control how often Watchtower checks for updates (seconds).
- `--cleanup` removes old images after updates to free disk space.

## Environment File (`.env`)

This compose uses an environment file to keep credentials and configuration out of `docker-compose.yml`.

- Example file: `services/watchtower/.env.example`
- To use it, copy it to `.env` and edit values:

```bash
cp services/watchtower/.env.example services/watchtower/.env
# then edit services/watchtower/.env
```

The example contains common Watchtower variables such as `WATCHTOWER_SCHEDULE`, `WATCHTOWER_CLEANUP`, and optional SMTP notification settings. Leave SMTP values blank if you don't use email notifications.

## Security Considerations

- Watchtower requires access to the Docker socket — this is equivalent to root privileges for the Docker host. Run only on trusted hosts.
- Use the `--label-enable` option and mark only specific containers for auto-updates if you want tighter control.
- Review updated images before enabling automatic restarts on critical services.

## Troubleshooting

- Watchtower not updating:
  - Confirm it has access to `/var/run/docker.sock`.
  - Check logs: `docker logs watchtower`.
  - Ensure images are tagged (latest vs specific tags); watchtower detects new digests for images.

- Container keeps restarting after update:
  - The new image might fail on startup — check the container logs for errors.

## Best Practices

- Use specific version tags for critical services and enable label-based updates for only those you trust to auto-update.
- Test updates in a staging environment before enabling auto-updates on production services.
- Combine with monitoring (e.g., Uptime Kuma) to alert if a service goes down after an update.

------------------------------------------------------------------------

If you want, I can add an example `docker-compose.override.yml` that limits watchtower to specific containers or configures scheduling. Would you like that?
