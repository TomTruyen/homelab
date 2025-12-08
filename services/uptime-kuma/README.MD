# Uptime Kuma

A self-hosted monitoring tool that tracks the uptime and performance of websites, servers, and services with customizable alerts and notifications.

🔗 https://github.com/louislam/uptime-kuma

------------------------------------------------------------------------

## 🚀 Setup

### Docker Compose

```yaml
services:
  uptime-kuma:
    image: louislam/uptime-kuma:latest
    container_name: uptime-kuma
    restart: always
    ports:
      - "3001:3001"
    volumes:
      - uptime-kuma-data:/app/data

volumes:
  uptime-kuma-data:
```

### Initial Configuration

1. Access Uptime Kuma at `https://<server-ip>:3001`
2. Create your admin account

------------------------------------------------------------------------

## 📱 Telegram Notifications

To receive alerts via Telegram, configure the Telegram bot integration:

### Steps

1. Navigate to **Settings → Notifications**
2. Click **Add Notification** and select **Telegram**
3. Fill in your credentials:
   - **Telegram Bot Token**: Your bot token from BotFather
   - **Chat ID**: Your Telegram chat ID
4. Configure notification settings:
   - Enable **Protect Forwarding/Saving**
   - Enable **Default enabled**
   - Enable **Apply on all existing monitors**
5. Test the notification to verify it works

------------------------------------------------------------------------

## 📊 Adding Monitors

### Create a New Monitor

1. Click **Add New Monitor**
2. Configure the monitor:
   - **URL**: Website or service endpoint to monitor
   - **Heartbeat Interval**: `300` (seconds between checks)
   - **Retries**: `3` (retry attempts before marking as down)
   - **Notification**: Select **Telegram** for alerts
3. Save and the monitor will start tracking

### Monitor Settings

| Setting | Recommended Value | Description |
|---------|------------------|-------------|
| Heartbeat Interval | 300 | Check interval in seconds (5 minutes) |
| Retries | 3 | Number of retry attempts |
| Notification | Telegram | Send alerts to Telegram when down |

------------------------------------------------------------------------

## 🌐 Monitored Services

Current list of services monitored:

- https://tomtruyen.dev
- https://traefik.tomtruyen.dev
- https://mailsift.tomtruyen.dev
- https://droidlaunch.tomtruyen.dev
- https://docs.droidlaunch.tomtruyen.dev
- https://thijswoodworking.be