# OpenMediaVault

A Linux-based operating system designed for network-attached storage (NAS), providing a user-friendly web interface for managing file sharing, storage, and services.

🔗 https://www.openmediavault.org/

------------------------------------------------------------------------

## ⚠️ Important Notes

- **Be careful when upgrading major versions** - failures can break SSH and other services
- **RAID Configuration**: Always plug the USB docking station into the same port after setup

------------------------------------------------------------------------

## 🚀 Setup

### Installation on Raspberry Pi

Use this script to install OpenMediaVault:

```bash
curl -sSL https://github.com/OpenMediaVault-Plugin-Developers/installScript/raw/master/install | sudo bash
sudo rm openmediavault-xxx.deb
```

### First-Time Setup

After installation, run the firstaid utility to configure networking:

```bash
sudo omv-firstaid
```

**Important**: Configure with **static IPv4** (disable DHCP and IPv6)

------------------------------------------------------------------------

## 🔧 Configuration

### Change Workbench (Web UI) Port

By default, OMV uses port 80. To change it to another port (e.g., 8000) for use with other services:

**Via Dashboard:**
1. Navigate to **System → Workbench**
2. Change the port number
3. Apply changes

**Via Command Line:**
```bash
sudo omv-firstaid
```

### Mount Storage Drives

1. Go to **Storage → File Systems**
2. Click the **Play icon (▶️)** next to your drive
3. Click **Save**
4. Apply changes when prompted

### Create Shared Folders

1. Navigate to **Storage → Shared Folders**
2. Click **➕** to create a new shared folder
3. Configure:
   - **Name**: `NAS` (or your preferred name)
   - **File system**: Select your mounted drive
   - **Relative Path**: `/`
   - **Permissions**: Default
4. Click **Save**

### Configure NFS (Network File System)

1. Navigate to **Services → NFS**
2. Go to **Settings** and enable NFS
3. Go to **Shares**:
   - Click **➕** to add a new share
   - **Shared folder**: Select your shared folder
   - **Client**: `192.168.1.0/24` (adjust to your network)
   - **Permission**: Read/Write
4. Click **Save**

### Configure SMB/CIFS (Windows Sharing)

1. Navigate to **Services → SMB/CIFS**
2. Go to **Settings** and enable SMB
3. Go to **Shares**:
   - Click **➕** to add a new share
   - **Shared folder**: Select your shared folder
   - Leave other settings as default
4. Click **Save**

### Manage Users

1. Navigate to **Users → Users**
2. Find or create a user account
3. Click **Edit (✏️)** to configure password and permissions
4. Click **Save**

### Apply Changes

After making any configuration changes, click **Apply** when prompted to save and apply all changes.

------------------------------------------------------------------------

## 💾 RAID Configuration

### Create RAID 1 (Mirroring)

Before setting up RAID, **back up your data or use new drives**.

```bash
# List available drives
ls -l /dev/s*

# Create RAID 1 array
mdadm --create /dev/md/name /dev/sda1 /dev/sdb1 --level=1 --raid-devices=2
```

**Wait for sync to complete** - Monitor progress on the OMV dashboard at `http://192.168.0.150`

**Important**: After setup, always plug the USB docking station into the **same USB port** (typically USB 3.0, bottom port closest to ethernet).

------------------------------------------------------------------------

## 🔗 Create Symlinks

To access NAS data from alternative paths:

```bash
# Create symlink from /srv to /mnt/ExternalHDD
sudo ln -s /srv/<Generated_Disk_Name>/ /mnt/ExternalHDD
```

Note: The trailing slash is important for correct functionality.

------------------------------------------------------------------------

## 📖 Default Credentials

| Item | Value |
|------|-------|
| **URL** | `http://<server-ip>` |
| **Username** | `admin` |
| **Password** | `openmediavault` |

**Change these credentials immediately after first login.**

------------------------------------------------------------------------

## 📚 Resources

- [Official Documentation](https://docs.openmediavault.org/)
- [OpenMediaVault GitHub](https://github.com/openmediavault/openmediavault)
- [Community Forum](https://forum.openmediavault.org/)