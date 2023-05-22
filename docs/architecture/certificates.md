# Certificates

The current backend will be hosted on `data.tribist.com` because I had an
existing domain name.

We use a wildcard domain certificate for now.

The frontend will also be hosted here on `ui.tribist.com`. Frontend will be
a nextjs APP run using pm2.

# setup

[Source](https://www.digitalocean.com/community/tutorials/how-to-create-let-s-encrypt-wildcard-certificates-with-certbot)

## Install certbot

`sudo apt install certbot`

## Install the digitalocean certbot plugin

`sudo apt install python3-certbot-dns-digitalocean`

## setup ini file for creds

```~certbot-creds.ini
dns_digitalocean_token=<get a personal access token from digitalocean>
```

Proper permissions for the ini file
```
chmod 600 ~/certbot-creds.ini
```

```
sudo certbot certonly \
  --dns-digitalocean \
  --dns-digitalocean-credentials ~/certbot-creds.ini \
  -d '*.tribist.com'
```

Update apache virtual hosts now









                o
