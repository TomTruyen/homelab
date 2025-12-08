# Pi-hole

A network-wide DNS adblocking service that blocks ads, trackers, and malware at the DNS level for all devices on your network.

🔗 https://pi-hole.net/

------------------------------------------------------------------------

## 🚀 Setup

### Docker Compose

```yaml
services:
  pihole:
    container_name: pihole
    image: pihole/pihole:latest
    ports:
      - "53:53/tcp"
      - "53:53/udp"
      - "8080:80/tcp"
    environment:
      TZ: 'Europe/Brussels'
      PIHOLE_DNS_1: 8.8.8.8
      PIHOLE_DNS_2: 8.8.8.4
    volumes:
      - './etc-pihole:/etc/pihole'
      - './etc-dnsmasq.d:/etc/dnsmasq.d'
    restart: unless-stopped
    dns:
      - 127.0.0.1
```

### Initial Access

1. Open your browser and navigate to `http://<server-ip>:8080/admin`
2. Retrieve the admin password:

```bash
docker exec -it pihole pihole -a -p
```

3. Use the password to log in to the dashboard

------------------------------------------------------------------------

## 🛡️ Dashboard Configuration

### Add Ad-Blocking Lists

1. Navigate to **Adlists** in the dashboard
2. Add blocklists for ad blocking and tracking prevention
3. Recommended starting point: [StevenBlack's hosts file](https://github.com/StevenBlack/hosts)

### Configure DNS Servers

1. Go to **Settings → DNS**
2. Set your preferred DNS providers:
   - Upstream DNS Servers: Google (8.8.8.8), Cloudflare (1.1.1.1), etc.
3. Save and update

### Update Gravity

After adding new blocklists:

1. Navigate to **Tools → Update Gravity**
2. Click **Update** to download and apply the newly added lists
3. Regularly update Gravity to ensure the latest lists are applied

------------------------------------------------------------------------

## 🌐 Network DNS Configuration

Optionally configure your network devices to use Pi-hole as their primary DNS.

### Android

1. Go to **Settings → Network & Internet → Internet → [Select Wi-Fi SSID]**
2. Tap the settings icon (gear) → Edit
3. Change **IP Settings** to **Manual**
4. Configure:
   - **DNS 1**: IP address of Pi-hole
   - **DNS 2**: `8.8.8.8` (backup)
5. Save

### iOS

1. Go to **Settings → Wi-Fi → [Select your network]**
2. Tap the info (ℹ️) icon
3. Select **DNS → Configure DNS → Manual**
4. Configure:
   - **Primary DNS**: IP address of Pi-hole
   - **Secondary DNS**: `8.8.8.8` (backup)
5. Save

### Windows 11

1. Go to **Settings → Network & Internet → Ethernet/Wi-Fi**
2. Click **Edit** under DNS Server Assignment
3. Configure:
   - **Preferred DNS (IPv4)**: IP address of Pi-hole
   - **Alternative DNS**: `8.8.8.8`
   - Ensure **DNS over HTTPS** is **OFF**
4. Save

------------------------------------------------------------------------

## 🔧 Troubleshooting

### systemd-resolved Port Conflict

**Issue**: `systemd-resolved` uses port 53 by default, preventing Pi-hole from binding to it. Disabling it entirely breaks DNS for other services.

**Solution**: Reconfigure `systemd-resolved` to free port 53 while maintaining DNS access.

#### Steps

1. Edit the configuration file:

```bash
sudo nano /etc/systemd/resolved.conf
```

2. Set the following values in the `[Resolve]` section:

```ini
[Resolve]
# Upstream DNS servers (using Google as example)
DNS=8.8.8.8
FallbackDNS=8.8.4.4

# Disable DNS stub listener (frees port 53)
DNSStubListener=no
```

3. Restart systemd-resolved:

```bash
sudo systemctl restart systemd-resolved
```

**Result**: `systemd-resolved` now uses Google's DNS while freeing port 53 for Pi-hole. Other services on the host maintain DNS access through `systemd-resolved`.

------------------------------------------------------------------------

## 📊 Monitoring

- **Dashboard**: Monitor queries, blocked domains, and statistics at `http://<server-ip>:8080/admin`
- **Query Log**: View real-time DNS queries and blocks
- **Gravity Database**: Manage blocklists and whitelist/blacklist entries