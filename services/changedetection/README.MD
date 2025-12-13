# Changedetection.io

A self-hosted web page change detection and notification service. Monitor websites for content changes and receive alerts when pages are updated.

🔗 https://changedetection.io/

------------------------------------------------------------------------

## 🚀 Purpose

- **Monitor websites** for content changes without manual checking
- **Detect updates** to specific page elements (text, price, availability, etc.)
- **Automated alerts** via email, Telegram, Discord, and other services
- **Self-hosted alternative** to commercial monitoring services

------------------------------------------------------------------------

## 🧩 Docker Compose

This setup includes two services:
- **changedetection.io**: Main web interface for managing monitored pages
- **playwright-chrome**: Headless browser for JavaScript-heavy websites

```yaml
services:
  changedetection:
    image: ghcr.io/dgtlmoon/changedetection.io
    container_name: changedetection
    hostname: changedetection
    environment:
      - PLAYWRIGHT_DRIVER_URL=ws://playwright-chrome:3000
    volumes:
      - changedetection-data:/datastore
    ports:
      - 5000:5000
    restart: unless-stopped
    depends_on:
      playwright-chrome:
        condition: service_started

  playwright-chrome:
    hostname: playwright-chrome
    image: dgtlmoon/sockpuppetbrowser:latest
    cap_add:
      - SYS_ADMIN
    restart: unless-stopped
    environment:
      - SCREEN_WIDTH=1920
      - SCREEN_HEIGHT=1024
      - SCREEN_DEPTH=16
      - MAX_CONCURRENT_CHROME_PROCESSES=10

volumes:
  changedetection-data:
```

### Configuration Notes

- **PLAYWRIGHT_DRIVER_URL**: Points to the Chrome browser container for JavaScript rendering
- **PORT**: Changedetection runs on `5000` (access at `http://localhost:5000`)
- **SCREEN_WIDTH/HEIGHT**: Virtual screen size for rendered pages (adjust based on needs)
- **MAX_CONCURRENT_CHROME_PROCESSES**: Limit concurrent browser instances to avoid resource exhaustion

------------------------------------------------------------------------

## 🚀 Getting Started

### Start the Service

```bash
docker compose up -d
```

### Access the Web Interface

1. Open your browser to `http://localhost:5000`
2. The interface loads with a default "Add a new watch"

------------------------------------------------------------------------

## 📋 Adding a Watch (Monitor)

### Steps

1. Click **Add a new watch** on the home page
2. Enter the **URL** you want to monitor
3. (Optional) Configure:
   - **Watch type**: Whole page, specific CSS element, or JSON path
   - **Fetch method**: Simple fetch or Playwright (for JavaScript)
   - **Ignore patterns**: Skip text that changes (dates, tracking IDs, etc.)
4. Click **Add watch**
5. The service will check the page and save a baseline

### Check Frequency

By default, pages are checked periodically. Configure in the watch settings:
- **Edit watch** → Set **Check every** to your desired interval (minutes)

------------------------------------------------------------------------

## 🔔 Notifications

Changedetection can alert you when a page changes via multiple channels.

### Configure Notifications

1. Go to **Settings** (gear icon, top right)
2. Click **Notifications**
3. Add a notification service:
   - **Email**: SMTP configuration
   - **Telegram**: Bot token and chat ID
   - **Discord**: Webhook URL
   - **Other**: JSON POST webhook

### Enable Notifications for a Watch

1. **Edit** a watch
2. Under **Trigger**, select your notification service
3. Save

------------------------------------------------------------------------

## 🎯 Advanced Features

### CSS Selector Watches

Monitor a specific element instead of the entire page:

1. **Edit watch** → **Watch type** → Select **CSS Selector**
2. Enter the CSS selector (e.g., `.price`, `#status`)
3. Only changes to that element trigger notifications

### Ignore Patterns

Skip common non-content changes (dates, tracking pixels, ads):

1. **Edit watch** → **Ignore text patterns**
2. Add regex or plain text to ignore
3. Example: `\d{4}-\d{2}-\d{2}` to ignore dates

### Fetch Method

- **Fetch**: Default, fast, lightweight
- **Playwright**: For JavaScript-heavy sites that need rendering

------------------------------------------------------------------------

## 🔧 Troubleshooting

### Service won't start

- Check Docker daemon is running: `docker ps`
- Verify `playwright-chrome` container is healthy: `docker logs playwright-chrome`
- Ensure port `5000` is not in use: `lsof -i :5000`

### Watch not detecting changes

- Verify the page is actually changing (manual check)
- Try switching **Fetch method** from "Fetch" to "Playwright" (slower but handles JS)
- Check **Ignore text patterns** — you might be ignoring the change

### Browser container crashing

- Increase `MAX_CONCURRENT_CHROME_PROCESSES` if monitoring many pages
- Monitor memory usage: `docker stats playwright-chrome`
- Reduce concurrent watches or increase available memory

### Notifications not sending

- Test notification in **Settings → Notifications** before using
- Verify API keys/tokens are correct
- Check service logs: `docker logs changedetection`

------------------------------------------------------------------------

## 📊 Monitoring Tips

- **Price tracking**: Monitor e-commerce sites for price drops
- **Status pages**: Watch for service outages and updates
- **Content updates**: Track blogs, news sites, or documentation changes
- **Availability**: Monitor product availability without constant manual checking

------------------------------------------------------------------------

## 📚 Resources

- [Official Documentation](https://changedetection.io/)
- [GitHub Repository](https://github.com/dgtlmoon/changedetection.io)
- [Discord Community](https://discord.gg/MjWfVKvj8Y)
