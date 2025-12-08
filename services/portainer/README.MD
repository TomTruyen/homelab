# Portainer

A lightweight, open-source container management UI that helps you manage Docker containers, images, networks, and volumes with an intuitive web interface.

🔗 https://www.portainer.io/

------------------------------------------------------------------------

## 🚀 Setup

### Docker Compose

```yaml
services:
  portainer:
    image: portainer/portainer-ce:linux-arm
    container_name: portainer
    ports:
      - "9000:9000"
    restart: always
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - portainer_data:/data

volumes:
  portainer_data:
```

### Initial Access

1. Access Portainer at `http://<server-ip>:9000`
2. Create your admin account
3. Select "Docker" as the environment to manage

------------------------------------------------------------------------

## 📋 Features

- **Container Management**: Create, run, stop, and remove containers
- **Image Management**: Pull, push, and manage Docker images
- **Volume Management**: Create and manage persistent storage
- **Network Management**: Create and configure networks
- **Stack Deployment**: Deploy multi-container applications using docker-compose
- **User Access Control**: Manage user accounts and permissions
- **Templates**: Quick deployment of popular applications

------------------------------------------------------------------------

## 🔐 Security Best Practices

### Change Default Password

After initial setup, change the admin password:

1. Go to **Settings → Users**
2. Click on your admin account
3. Update your password

### Enable HTTPS

For production deployments, enable HTTPS:

1. Go to **Settings → SSL**
2. Upload your SSL certificate and key
3. Enable SSL

### Restrict Access

- Use a reverse proxy (like Traefik) to control external access
- Consider disabling port `9000` and only allowing `9443` (HTTPS)
- Use Portainer's user and team features to control access

------------------------------------------------------------------------

## 📚 Common Tasks

### Deploy a Container

1. Go to **Containers**
2. Click **Add container**
3. Select an image and configure settings
4. Click **Deploy**

### Deploy a Stack (Docker Compose)

1. Go to **Stacks**
2. Click **Add stack**
3. Upload a docker-compose.yml file or paste the content
4. Deploy

### Manage Multiple Hosts

1. Go to **Environments**
2. Click **Add environment**
3. Choose connection method (Docker Socket, TCP, etc.)
4. Follow the setup instructions

------------------------------------------------------------------------

## 🔧 Useful Settings

| Setting | Purpose |
|---------|---------|
| **Edge Agent** | Connect remote Docker hosts to Portainer |
| **Registries** | Add private Docker registries |
| **Webhooks** | Trigger stack updates via webhooks |
| **Backup** | Backup Portainer configuration |
| **User Management** | Control access and permissions |

------------------------------------------------------------------------

## 📖 Documentation

- [Official Documentation](https://docs.portainer.io/)
- [GitHub Repository](https://github.com/portainer/portainer)
- [Community Forums](https://www.portainer.io/community/)
