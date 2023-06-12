# Dev setup

## /etc/hosts

Add the following lines to your `/etc/hosts` file.

```
127.0.0.1       localtwitter.com
127.0.0.1       localapi.com
```

## Api

In the `/api` directory run `make dev`. This will do the following
1. create the required `.env` file by copying the file `dev.env`
2. insert the dev user `param108` into the db 
3. and start the server.

## TwitterLike

In the `/twitterlike` director run `yarn run dev`.

## How it works

Now go to `http://localtwitter.com:3000` on your browser.

The file `.env.development` specifies the port number (3000) and the backend URL (`http://localapi.com:8000`).

On the backend the file `.env` created by the `make dev` command specifies the frontend as (`http://localtwitter.com:3000`)

The backend is programmed such that if the environment variable `ENV` is `dev`, the system always logs in the user `param108` and returns a token. This avoids the need to ask twitter or any other service provider for authentication.

Now you can do everything as the user `param108`
