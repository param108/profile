# API server Deployment

We are choosing to avoid the overhead of docker for now.
We will use systemd to keep the executable running if it crashes.

# Prerequisites

1. github cli installed on server
2. The server binary must listen for `SIG_TERM` and gracefully shutdown.

# Overview

## Installation

1. ssh into server and copy the `systemd` unit file to `/etc/systemd/system/tribist.service`
   (Not `user` as we want this service to always run.)
2. copy `server.sh`  `/usr/bin/server.sh`
   - this will be triggered by the unit file to start|stop|restart the server
3. copy the github access token to `/home/cicd/gh.txt`
   - this should be readable only by root.
5. a shell script `restart_server.sh` will be placed at `/usr/bin/restart_server.sh`
   This will be called by the user `cicd` via sudo to trigger `systemctl restart tribist`.
   This is needed to avoid the user `cicd` to get more access than necessary via sudo.
5. check for successful start will use the script `wait_for_it.sh` and will fail in 30 seconds.
   We will hit `http://localhost:8383/health` which will also check the DB connection.
   
   The unit file is responsible to restart the server in case of failure.
6. `sudo systemctl daemon-reload`
7. `sudo systemctl start tribist`

## Upgrade
1. ssh into the server
2. execute `/usr/bin/restart_server.sh`
   This script is a wrapper to download the latest image and run the commands `systemctl stop tribist` and then `systemctl start tribist`
   The unit file will then run `/usr/bin/server.sh` with appropriate configurations
   `server.sh` is responsible to switch to the new server downloaded in restart_server.sh.
   
# Details 

## apache config for proxying
```
    ProxyPreserveHost On
    ProxyRequests Off
    ProxyPass / http://localhost:8383/
```

## artifact contents
    The artifact will be created by the `api_deploy` action on github.
    It will be called `server.tgz`
    It will consist of
    1. server <-- the api server
    2. env    <-- the environment file
    3. db     <-- directory of migrations for api

    The server file will be generated in `build` directory as in a normal build.

    The `env` file will be generated by the action.
    - extract the file contents from the github secret `API_ENV_CONFIG`
    - write it into `api/build/` directory

    The `db` directory will be copied from `api/db` directory

    In the action code, we can create a zipfile of `{server, env}` files.

## cicd user
    The user must have only the ability to restart the server.
    ```
        Cmnd_Alias TRIBIST_CMDS = /usr/bin/restart_server.sh
        
        cicd ALL=(root) NOPASSWD: TRIBIST_CMDS
    ```
    
    This allows `cicd` on `ALL` hosts to run `TRIBIST_CMDS` without PASSWORD as `root`
    
    Also `.ssh/authorized_keys` limits the cicd user to only run the `sudo /usr/bin/restart_server.sh` command
    like below -
    
    ```
    command="sudo /usr/bin/restart_server.sh" ssh-rsa AAAA...
    ```
## `restart_server.sh`

    1. download the latest release artifact from github.

        ```
        gh run -R param108/profile list -w api_deploy --json conclusion,databaseId,workflowDatabaseId -L 1 -q 'select(.[].conclusion = "success")' | jq .[0].databaseId
        ```

        use the `databaseId` to download artifact.

        ```
        gh run -R param108/profile download <databaseId from previous command> -n server.tgz -D /tmp/
        ```

	```
	cd /tmp/; rm -rf build; tar -zxvf /tmp/server.tgz
	```
    `server` is now present in `/tmp/build/server`.
    Unzip it and extract the `server` and `env` files.

    mv the server to `/home/cicd/server_new`
    mv env to `/home/cicd/.env`
    mv db to `/home/cicd/db`
    
## server.sh
    `$1` is one of `[stop|start]`
    
    First load the environment variables from `.env`. This will give us important variables needed
    by `server.sh`
    - `PORT`
    - `MIGRATE`
    - `UPDATE`
    
    Download of new server and server restart will only happen if environment variable `UPDATE` 
    is not set or if it is equal to `true`.
    
    Migrations will only run if `MIGRATE` is NOT SET or if it is equal to `true`.
    
    Working directory is `/home/tribist/api/`
    
    *start*
    
    Assumes the server has stopped running.
 
    Do steps 1 through 4 only if `UPDATE` is NOT set or `UPDATE` is `true`
    
    
    1. `mv` the `server` binary `/home/cicd/server_new` to `/home/tribist/api/server_new` if server_new exists

    2. `mv` the `env` file from `/home/cicd/.env` to `/home/tribist/api/.env` if .env exists
    
    3. `mv` the `db` directory from `/home/cicd/db` to `/home/tribist/api/db` if db exists
    
    4. check for `/home/tribist/api/server_new` and if it exists `mv` it to `/home/tribist/api/server`
    
    5. migrate (ONLY if `MIGRATE` is NOT set or `MIGRATE` is `true`) using the command

       `/home/tribist/api/server migrate`

    6. run `/home/tribist/api/server`
    
    *stop*
    
    1. stop the server by sending `SIG_TERM` signal to server using PID file.
    
    Note: If any error occurs we simply stop there.
