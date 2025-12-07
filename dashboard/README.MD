# Homelab Dashboard (via Glance)

Powered by the Glance dashboard\
🔗 https://github.com/glanceapp/glance

------------------------------------------------------------------------

## 🚀 Setup

Rename the environment file:

``` bash
mv .env.example .env
```

Start the container:

``` bash
docker compose up -d
```

------------------------------------------------------------------------

## 🌐 Accessing Your Homelab from Outside (Cloudflare Tunnel)

To securely expose your Homelab dashboard to the internet, you'll use a
**Cloudflare Tunnel**.

------------------------------------------------------------------------

### 1. Prepare Your Domain on Cloudflare

1.  Go to https://www.cloudflare.com/ and log in.
2.  If the domain is **not yet linked** to Cloudflare:
    1.  Open **Domains**.
    2.  Click **Onboard a domain**.
    3.  Enter your domain (e.g., `tomtruyen.dev`).
    4.  Choose the **Automatic** setup option.
    5.  When Cloudflare shows DNS records, **uncheck Proxy Status** for
        all records --- your Cloudflare Tunnel will proxy only the
        services you explicitly add.
    6.  Update your nameservers on your registrar (e.g., Namecheap):
        -   Set Nameservers to **Custom DNS**
        -   Enter the **NS** values provided by Cloudflare.

------------------------------------------------------------------------

### 2. Create a Cloudflare Tunnel

1.  Navigate to **Zero Trust** in the sidebar.

2.  Go to **Networks → Connectors → Create a Tunnel**.

3.  Select **Cloudflared**.

4.  Name the tunnel (e.g., `Homelab Tunnel`).

5.  Choose **Docker** as the environment and copy the generated command.
    It should look like:

    ``` bash
    docker run cloudflare/cloudflared:latest tunnel --no-autoupdate run --token <TOKEN FROM CLOUDFLARE>
    ```

------------------------------------------------------------------------

### 3. Add an Application to the Tunnel

1.  Choose a **subdomain** to expose the service.
2.  Configure the service:
    -   **Type**: `HTTP` (unless your service uses self-signed TLS →
        then choose HTTPS)
    -   **URL**: internal IP + port (e.g., `192.168.0.150:8080`)
3.  Under **Additional Application Settings**:
    -   Go to **TLS**
    -   Enable **No TLS Verify**

------------------------------------------------------------------------

### 4. Secure Access with Cloudflare Access

Right now the service is public --- let's lock it down so only you can
access it.

1.  Go to **Access → Applications**.
2.  Click **Add an application**.
3.  Choose **Self-hosted**.
4.  Configure the application:
    -   **Name**: `Tom's Homelab`
    -   **Public hostname**: `homelab.tomtruyen.dev`
5.  Add a policy:
    -   **Policy name**: `PersonalAccessOnly`
    -   **Action**: `Allow`
    -   **Rule type**: **Emails**
    -   Add your allowed email addresses (personal, work, etc.)

------------------------------------------------------------------------

### 5. (Optional) Add More Identity Providers

If you want login via GitHub, Google, etc.:

1.  Go to **Zero Trust → Settings → Authentication → Identity
    Providers**.
2.  Choose a provider (e.g., GitHub).
3.  Follow the setup steps.
