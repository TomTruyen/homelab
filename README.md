# 🏠 Homelab

A personal homelab setup built around Docker and Docker Compose. This repo contains the configs and tooling I use to run, monitor, and maintain self-hosted services from a single dashboard.

---

## 📦 What’s in here

### Dashboard (Glance)

A single, unified dashboard that pulls data from all services.

* Glance dashboard configuration
* Custom APIs (e.g. AniList)
* Glance Agent for remote servers
* Homelab Sync (auto-pulls repo updates and restarts services)

📁 `dashboard/`

---

### Services

Each service is self-contained and can be deployed independently.

* **Pi-hole** – Network-wide DNS ad blocking
* **Portainer** – Docker management UI
* **Uptime Kuma** – Uptime & health monitoring
* **OpenMediaVault** – NAS & storage management
* **PriceBuddy** – Amazon price tracking
* **Watchtower** – Automatic Docker image updates
* **GitHub Backup Script** – Backup GitHub repos to NAS

📁 `services/`

---

## 🚀 Quick Start

### Requirements

* Linux-based system (server, VM, Raspberry Pi)
* Docker + Docker Compose
* Git

### Clone the repo

```bash
git clone git@github.com:TomTruyen/homelab.git
cd homelab
```

### Start the dashboard

```bash
cd dashboard
cp .env.example .env
# edit .env
docker compose up -d --build
```

### Start a service

```bash
cd services/<service-name>
docker compose up -d
```

---

## 🎯 Key Features

* **Single dashboard** for all services
* **Modular setup** – run only what you need
* **Auto-updates** via Homelab Sync + Watchtower
* **Remote monitoring** using agents
* **Easy recovery** – everything is version-controlled

---

## 🔐 Security & Access

### External Access

All externally exposed services are secured via **Cloudflare Tunnel**:

* No port forwarding
* TLS encryption
* Zero Trust (email-based authentication)

### Local Access (examples)

* Dashboard: `http://<server-ip>:8080`
* Pi-hole: `http://<server-ip>:8080/admin`
* Portainer: `http://<server-ip>:9000`
* Uptime Kuma: `http://<server-ip>:3001`
* PriceBuddy: `http://<server-ip>:5000`

---

## 🔄 Automatic Updates

**Homelab Sync** automatically:

* Watches the GitHub repo
* Pulls changes
* Restarts only affected services

Updating the homelab is as simple as pushing to GitHub.

---

## 📚 Documentation

Every component has its own README with setup and configuration details:

* Dashboard & agents → `dashboard/`
* Individual services → `services/<service>/`

---

## 📝 License

Personal project. Feel free to use it as inspiration for your own homelab.

---

## 🤝 Contributing

Not actively seeking contributions, but you’re welcome to fork or reuse anything here.
