# Twitterlike server Deployment

We are choosing to avoid the overhead of docker for now.
We will use systemd to keep the executable running if it crashes.

# Prerequisites

1. github cli installed on server
2. nvm installed on server
3. node version v16.20.0 installed

# Overview

When a push is made to main in the twitterlike directory, an artifact is created called
`twitter.tgz` this has 
1. the complete code including `node_modules` after `yarn build`. 
2. `.env` file created from `TWTR_ENV_CONFIG` from cicd secrets.

The workflow then calls `restart_twitter.sh` on the server. This command basically does

```
systemctl stop twitter
sleep 5
systemctl start twitter
```

The first line just kills existing `twitter.sh` process and all children.
`systemctl start twitter` runs `/usr/bin/twitter.sh -c start -g gh.txt`. The `twitter.sh`
script is available in `/deployments/twitterlike` directory in this repo. 


## Start command

This command uses gh cli to download the latest `twitter.tgz` and untars it. After that
it kills the current `twitter.sh` process and then runs

`yarn run start -p 9090` 

to start the server.


## nvm command

We use nvm to manage our node version on the server. The current supported node version
is v16.20.0.
The code in twitter.sh calls `nvm use v16.20.0` so nvm needs to be setup properly.
In order to use nvm, the script sources `/root/nvm_init.sh`. This file is available
in `/deployments/twitterlike` directory of this repository. 


# Apache config for proxying

```
   ProxyPreserveHost On
   ProxyRequests Off
   ProxyPass / http://localhost:9090/
```

# Github

We use `TWTR_HOST` and `TWTR_PRIVATE_KEY` repository secrets to control this command.
In the authorized_keys file we setup forced commands for corresponding public key to only
call `restart_twitter.sh`
