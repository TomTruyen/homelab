# Glance Agent

The Glance Agent enables your main Glance dashboard to securely access data from remote servers in your homelab. This allows you to aggregate system information and service data from multiple machines into a single unified dashboard.

## What It Does

The agent runs on secondary servers and exposes their system information through an API that your main Glance dashboard can query. This is useful for:

- **System Stats**: CPU, memory, disk usage, and network information from remote machines
- **Service Monitoring**: Status and metrics from services running on different servers
- **Distributed Infrastructure**: View health and performance of your entire homelab from one dashboard

## Setup

### 1. On Your Secondary Server

Clone the repository on the remote server:

```bash
git clone <repository-url>
cd dashboard/agent
```

### 2. Configure the Agent

If needed, configure any environment variables in a `.env` file:

```bash
cp .env.example .env
# Edit .env with your settings
```

### 3. Start the Agent

```bash
docker compose up -d
```

The agent will start and listen for connections from your main Glance dashboard.

### 4. Configure Your Main Dashboard

On your primary Glance dashboard server, configure widgets to use the remote agent. For example, to fetch system stats from a remote server:

```yaml
- type: server-stats
  servers:
    - type: remote
      name: Raspberry PI 3B+
      url: http://192.168.0.151:27973
```

Replace `192.168.0.XXX` with the IP address of your secondary server and `8080` with the agent port.

## Multiple Agents

You can run this agent on multiple secondary servers. Each instance should be configured independently with its own Docker environment, and each should be accessible from your primary dashboard at a different IP or port.