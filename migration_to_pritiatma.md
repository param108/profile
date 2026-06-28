# Migration from tribist.com to pritiatma.in

This document describes the migration of the Profile application from `tribist.com` to `pritiatma.in`.

## New URL Structure

| Old URL | New URL |
|---------|---------|
| ui.tribist.com | profile.pritiatma.in |
| data.tribist.com | profile-api.pritiatma.in |

## Code Changes

### Backend (api/)

**1. `api/users/login/email/email_password.go:70`**

Changed hardcoded URL to use environment variable:
```go
// Before
loginURL := fmt.Sprintf("https://ui.tribist.com/login?key=%s", oneTime.ID)

// After
loginURL := fmt.Sprintf("%s/login?key=%s", os.Getenv("FRONTEND_URL"), oneTime.ID)
```

**2. Environment files updated with `FRONTEND_URL`:**
- `api/ci.env`
- `api/dev.env`
- `api/test.env`
- `api/env.sample`

### Frontend (twitterlike/)

**1. `twitterlike/.env.production`**
```bash
# Before
NEXT_PUBLIC_BE_URL=https://data.tribist.com
NEXT_PUBLIC_HOST=http://ui.tribist.com

# After
NEXT_PUBLIC_BE_URL=https://profile-api.pritiatma.in
NEXT_PUBLIC_HOST=https://profile.pritiatma.in
```

**2. `twitterlike/app/components/header.tsx`**

Changed hardcoded link to relative URL and used Next.js Link component:
```tsx
// Before
<a href="https://ui.tribist.com/user/param108/tweets">

// After
<Link href="/user/param108/tweets">
```

## Environment Variables

### Backend Production Environment

The following variables were updated in the production environment (`env.prod`):

| Variable | Old Value | New Value |
|----------|-----------|-----------|
| HOST | tribist.com | profile-api.pritiatma.in |
| FRONTEND_URL | (new) | https://profile.pritiatma.in |
| AUTH_REDIRECT_URL | https://ui.tribist.com/login-redirect | https://profile.pritiatma.in/login-redirect |
| ALLOWED_CLIENTS | https://ui.tribist.com | https://profile.pritiatma.in |

### Frontend Production Environment

Updated in `env.local.prod`:

| Variable | Old Value | New Value |
|----------|-----------|-----------|
| NEXT_PUBLIC_BE_URL | https://data.tribist.com | https://profile-api.pritiatma.in |
| NEXT_PUBLIC_HOST | https://ui.tribist.com | https://profile.pritiatma.in |

## GitHub Secrets Updated

| Secret | Description |
|--------|-------------|
| API_ENV_CONFIG | Backend production environment |
| TWTR_ENV_CONFIG | Frontend production environment |
| API_HOST | Changed from tribist.com to IP address (167.71.239.0) |
| TWTR_HOST | Changed from tribist.com to IP address (167.71.239.0) |

## Server Configuration

### DNS Records

Added A records for `pritiatma.in`:
```
profile.pritiatma.in      A    167.71.239.0
profile-api.pritiatma.in  A    167.71.239.0
```

### SSL Certificates

Generated Let's Encrypt certificates:
```bash
certbot certonly --apache -d profile.pritiatma.in -d profile-api.pritiatma.in
```

Certificates stored at:
- `/etc/letsencrypt/live/profile.pritiatma.in/fullchain.pem`
- `/etc/letsencrypt/live/profile.pritiatma.in/privkey.pem`

### Apache Configuration

Created two new Apache virtual host configurations in `/etc/apache2/sites-available/`:

**1. `profile.pritiatma.in-le-ssl.conf`** (Frontend)
```apache
<IfModule mod_ssl.c>
<VirtualHost *:443>
        ServerName profile.pritiatma.in
        ServerAdmin webmaster@localhost

        ProxyPreserveHost On
        ProxyRequests Off
        ProxyPass / http://localhost:9090/

        ErrorLog ${APACHE_LOG_DIR}/error.log
        CustomLog ${APACHE_LOG_DIR}/access.log combined

Include /etc/letsencrypt/options-ssl-apache.conf
ServerAlias profile.pritiatma.in
SSLCertificateFile /etc/letsencrypt/live/profile.pritiatma.in/fullchain.pem
SSLCertificateKeyFile /etc/letsencrypt/live/profile.pritiatma.in/privkey.pem
</VirtualHost>
</IfModule>
```

**2. `profile-api.pritiatma.in-le-ssl.conf`** (Backend API)
```apache
<IfModule mod_ssl.c>
<VirtualHost *:443>
        ServerName profile-api.pritiatma.in
        ServerAdmin webmaster@localhost

        ProxyPreserveHost On
        ProxyRequests Off
        ProxyPass / http://localhost:8383/

        ErrorLog ${APACHE_LOG_DIR}/error.log
        CustomLog ${APACHE_LOG_DIR}/access.log combined

Include /etc/letsencrypt/options-ssl-apache.conf
ServerAlias profile-api.pritiatma.in
SSLCertificateFile /etc/letsencrypt/live/profile.pritiatma.in/fullchain.pem
SSLCertificateKeyFile /etc/letsencrypt/live/profile.pritiatma.in/privkey.pem
</VirtualHost>
</IfModule>
```

Enable the sites:
```bash
a2ensite profile.pritiatma.in-le-ssl.conf
a2ensite profile-api.pritiatma.in-le-ssl.conf
systemctl reload apache2
```

### CORS Configuration

CORS is handled by the Go backend (not Apache) to avoid duplicate headers. The backend uses:
```go
originsOk := handlers.AllowedOrigins([]string{"*"})
```

## Service Configuration

The existing systemd services (`tribist.service` and `twitter.service`) continue to work as they reference local paths and ports:
- Backend: port 8383
- Frontend: port 9090

## Cloudflare

If using Cloudflare proxy:
1. Temporarily disable proxy (orange cloud → gray cloud) for SSL certificate generation
2. Re-enable proxy after certificates are issued
3. Cloudflare handles SSL termination at edge

## Rollback Procedure

If rollback is needed:
1. Re-enable old Apache configs for `tribist.com`
2. Revert GitHub secrets to old values
3. Redeploy via CI/CD or restart services manually

## Files Modified

### profile repository
- `api/users/login/email/email_password.go`
- `api/ci.env`
- `api/dev.env`
- `api/test.env`
- `api/env.sample`
- `twitterlike/.env.production`
- `twitterlike/app/components/header.tsx`

### pritiatma repository
- `deployment/profile.pritiatma.in-le-ssl.conf` (new)
- `deployment/profile-api.pritiatma.in-le-ssl.conf` (new)

### Server files
- `/etc/apache2/sites-available/profile.pritiatma.in-le-ssl.conf`
- `/etc/apache2/sites-available/profile-api.pritiatma.in-le-ssl.conf`
- `/home/tribist/api/.env`
- `/home/tribist/twitterlike/.env.local`

### Credentials (local)
- `~/credentials/tribist/env.prod`
- `~/credentials/tribist/env.local.prod`

## Migration Date

Completed: June 29, 2026
