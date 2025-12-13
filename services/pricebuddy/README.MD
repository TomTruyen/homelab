# PriceBuddy

PriceBuddy is a small self-hosted price-tracking and alerting web application that scrapes and stores product prices so you can track them over time.

> Image used in this repo: `jez500/pricebuddy:latest`

------------------------------------------------------------------------

## 🚀 Purpose

- Track product prices across websites
- Store historical price data in a local MySQL database
- Provide a web UI to view tracked items and change history

------------------------------------------------------------------------

## 🧩 Docker Compose (what this repo runs)

The compose file defines two services:

- `pricebuddy` — the web application (exposes port `5000`)
- `db` — a MySQL 8.0 database used by PriceBuddy

Excerpt from `docker-compose.yml`:

```yaml
services:
  pricebuddy:
    image: jez500/pricebuddy:latest
    container_name: pricebuddy
    restart: unless-stopped
    depends_on:
      - db
    ports:
      - "5000:80"
    environment: # see .env.example

  db:
    image: mysql:8.0
    container_name: pricebuddy-db
    restart: unless-stopped
    environment: # see .env.example
    volumes:
      - ./data/mysql:/var/lib/mysql

volumes:
  pricebuddy_storage:
```

------------------------------------------------------------------------

## Environment Variables

The compose file currently defines environment variables inline. For better security and per-host configuration copy `services/pricebuddy/.env.example` to `services/pricebuddy/.env` and edit values there.

Key variables (from `.env.example`):

| Variable | Required | Description |
|----------|----------:|-------------|
| `APP_NAME` | Yes | Application name |
| `APP_ENV` | Yes | Environment (`production` / `development`) |
| `APP_DEBUG` | No | Debug mode (`true`/`false`) |
| `APP_URL` | Yes | Public URL of app (e.g. `http://yourhost:5000`) |
| `DB_CONNECTION` | Yes | Database driver (mysql) |
| `DB_HOST` | Yes | DB host/container name (db) |
| `DB_PORT` | Yes | DB port (3306) |
| `DB_DATABASE` | Yes | Database name |
| `DB_USERNAME` | Yes | DB user |
| `DB_PASSWORD` | Yes | DB password |
| `MYSQL_ROOT_PASSWORD` | Yes | MySQL root password (set in `db` service) |

------------------------------------------------------------------------

## Setup

1. Configure environment variables in `docker-compose.yml`

2. Start the stack:

```bash
cd services/pricebuddy
docker compose up -d
```

3. Wait for the database to initialize (give MySQL ~10–30s on first run).

4. Configure first user:
   ```bash
    docker exec -it pricebuddy php artisan tinker

    # Paste & modify values
    use App\Models\User;

    User::create([
        'name' => 'admin',
        'email' => 'admin@example.com',
        'password' => bcrypt('changeme123'),
    ]);

    # Finish
    exit
   ```

5. Visit the web UI at `http://<host-ip>:5000` (replace `<host-ip>` with your host address).

6. Go to settings a configure notifications

------------------------------------------------------------------------

## Database / Migrations

If the image exposes artisan/management commands you may need to run migrations on first run. If required, run the container shell or an exec command (example - will vary by image):

```bash
# Example: run migrations (only if the container provides such commands)
docker compose exec pricebuddy /bin/sh -c "php artisan migrate --force"
```

If the image does not include migration tooling, the database schema may be created automatically on first start.

------------------------------------------------------------------------

## Backups

- MySQL data is persisted to `services/pricebuddy/data/mysql` (see compose). Back up that directory or use a DB dump to preserve data.

Example DB dump:

```bash
docker exec -t pricebuddy-db mysqldump -u root -p${MYSQL_ROOT_PASSWORD} pricebuddy > pricebuddy_$(date +%F).sql
```

------------------------------------------------------------------------

## Security & Notes

- The current compose file contains example plain-text passwords. Move these values to `services/pricebuddy/.env` and do **not** commit secrets to Git.
- Use strong passwords for `DB_PASSWORD` and `MYSQL_ROOT_PASSWORD`.
- If exposing the app to the internet, place it behind a reverse proxy (Traefik) with TLS and authentication.

------------------------------------------------------------------------

## Troubleshooting

- App not reachable on port 5000:
  - Ensure the container is running: `docker compose ps`
  - Check logs: `docker compose logs -f pricebuddy`
  - Confirm the DB container is healthy: `docker compose logs -f db`

- Database connection errors:
  - Confirm env values for DB host/user/pass are correct
  - Ensure the database initialized correctly and user exists

------------------------------------------------------------------------

## Resources

- Image: https://github.com/jez500/pricebuddy (if available)
- Changedetection/price tracking: adapt as needed

If you want, I can also:

- Add `services/pricebuddy/.env.example` to the repository (I can create it for you)
- Move inline compose environment values to an `.env` and update `docker-compose.yml` to use `env_file`

Would you like me to do either of those?