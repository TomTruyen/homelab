# 🏠 Homelab

A comprehensive collection of self-hosted services and configurations for a personal homelab infrastructure. This repository contains everything needed to set up and manage your own home automation, media, storage, and monitoring systems.

------------------------------------------------------------------------

## 📂 Repository Structure

### Dashboard
A unified dashboard powered by [Glance](https://github.com/glanceapp/glance) that aggregates data from all your services.

- **[Dashboard README](./dashboard/README.MD)** - Main dashboard setup and configuration
- **[Custom APIs](./dashboard/api/README.MD)** - Custom API services for dashboard integration
  - **[AniList API](./dashboard/api/anilist/README.MD)** - Anime tracking and episode information
- **[Glance Agent](./dashboard/agent/README.MD)** - Remote server data collection agent
- **[Homelab Sync](./dashboard/homelab-sync/README.MD)** - Automatic config sync and service restart

### Services
Individual service configurations and management.

- **[Pi-hole](./services/pi-hole/README.MD)** - Network-wide DNS adblocking
- **[Portainer](./services/portainer/README.MD)** - Docker container management UI
- **[Uptime Kuma](./services/uptime-kuma/README.MD)** - Service uptime monitoring and alerts
- **[OpenMediaVault](./services/openmediavault/README.MD)** - Network-attached storage (NAS) management

------------------------------------------------------------------------

## 🚀 Quick Start

### Prerequisites

- Docker and Docker Compose installed
- Linux-based system (Raspberry Pi, Ubuntu, Debian, etc.)
- Git

### Clone Repository

```bash
git clone git@github.com:TomTruyen/homelab.git
cd homelab
```

### Deploy Dashboard

```bash
cd dashboard
cp .env.example .env
# Edit .env with your configuration
docker compose up -d --build --force-recreate
```

### Deploy Services

Each service can be deployed independently:

```bash
# Pi-hole (DNS Adblocking)
cd services/pi-hole
docker compose up -d

# Portainer (Container Management)
cd services/portainer
docker compose up -d

# Uptime Kuma (Monitoring)
cd services/uptime-kuma
docker compose up -d

# Watchtower
cd services/watchtower
docker compose up -d

# PriceBuddy
cd services/pricebuddy
docker compose up -d
```

------------------------------------------------------------------------

## 🎯 Key Features

### Dashboard
- **Unified Interface**: Single dashboard to view all services and data
- **Customizable Widgets**: Configure what information to display
- **Remote Data Collection**: Aggregate data from multiple servers via agents
- **Auto-sync**: Automatically pull updates and restart services

### Monitoring & Uptime
- **Network Monitoring**: Pi-hole provides DNS-level ad blocking and network stats
- **Uptime Tracking**: Uptime Kuma monitors all services with Telegram alerts
- **Health Checks**: Continuous monitoring of critical services

### Storage & Media
- **NAS Management**: OpenMediaVault provides easy storage management
- **File Sharing**: NFS and SMB/CIFS support for network file access
- **RAID Support**: Redundant storage configuration

### Container Management
- **Portainer UI**: Web interface for Docker management
- **Container Monitoring**: View logs, stats, and health of all containers

------------------------------------------------------------------------

## 🔐 Security

### External Access (Cloudflare Tunnel)

All services are secured behind a Cloudflare Tunnel with:
- **Zero Trust Access**: Email-based authentication
- **TLS Encryption**: All traffic is encrypted
- **No port forwarding needed**: Safe exposure of internal services

See [Dashboard README](./dashboard/README.MD#-accessing-your-homelab-from-outside-cloudflare-tunnel) for Cloudflare Tunnel setup.

### Local Access

Services are accessible on your local network at:
- Dashboard: `http://\<server-ip\>:8080`
- Pi-hole: `http://\<server-ip\>:8080/admin`
- Portainer: `http://\<server-ip\>:9000`
- Uptime Kuma: `http://\<server-ip\>:3001`
- PriceBuddy: `http://\<server-ip\>:5000`

------------------------------------------------------------------------

## 📚 Documentation

Each service and component has its own comprehensive README:

| Service | Purpose | Documentation |
|---------|---------|---|
| **Dashboard** | Unified control center | [README](./dashboard/README.MD) |
| **Pi-hole** | DNS adblocking & network stats | [README](./services/pi-hole/README.MD) |
| **Portainer** | Container management | [README](./services/portainer/README.MD) |
| **Uptime Kuma** | Service monitoring | [README](./services/uptime-kuma/README.MD) |
| **OpenMediaVault** | NAS & storage | [README](./services/openmediavault/README.MD) |
| **Watchtower** | Automatic Docker image updates | [README](./services/watchtower/README.MD) |
| **PriceBuddy** | Price Tracker (for Amazon) | [README](./services/pricebuddy/README.MD) |
| **AniList API** | Anime tracking | [README](./dashboard/api/anilist/README.MD) |
| **Glance Agent** | Remote data collection | [README](./dashboard/agent/README.MD) |
| **Homelab Sync** | Auto-update service | [README](./dashboard/homelab-sync/README.MD) |

------------------------------------------------------------------------

## 🔄 Automatic Updates

The **Homelab Sync** service automatically:
- Monitors your GitHub repository for changes
- Pulls updates to your local configuration
- Restarts only affected services

This means you can update your homelab configuration by simply pushing changes to GitHub!

See [Homelab Sync README](./dashboard/homelab-sync/README.MD) for setup.

------------------------------------------------------------------------

## 📝 License

This repository is for personal use. Modify and use as needed for your own homelab.

------------------------------------------------------------------------

## 🤝 Contributing

This is a personal project, but feel free to use it as a reference for your own homelab setup!