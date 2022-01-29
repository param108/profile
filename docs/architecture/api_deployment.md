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
3. copy the github access token to `/home/tribist/api/gh.txt`
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
   This script is a wrapper for the command `systemctl stop tribist` and then `systemctl start tribist`
   The unit file will then run `/usr/bin/server.sh` with appropriate configurations
   `server.sh` is responsible to download the new server image if any.
   
# Details 

## apache config for proxying
```
    ProxyPreserveHost On
    ProxyRequests Off
    ProxyPass / http://localhost:8383/
```

## cicd user
    The user must have only the ability to restart the server.
    ```
        Cmnd_Alias TRIBIST_CMDS = /usr/bin/restart_server.sh
        
        cicd ALL=(root) NOPASSWD: TRIBIST_CMDS
    ```
    
    This allows `cicd` on `ALL` hosts to run `TRIBIST_CMDS`` without PASSWORD as `root`
    
## server.sh
    `$1` is one of `[stop|start|restart]`
    
    Working directory is `/home/tribist/api/`
    
    *start*
    
    1. download the latest release artifact from github.

        ```
        gh run -R param108/profile list -w api_deploy --json conclusion,databaseId,workflowDatabaseId -L 1 -q 'select(.[].conclusion = "success")' | jq .[0].databaseId
        ```

        use the `databaseId` to download artifact.

        ```
        gh run -R param108/profile download <databaseId from previous command> -n server -D /tmp/
        ```

    use the run id obtained above to extract the `server` artifact to `/tmp/server`
    
    2. `mv` it to `/home/tribist/api/server_new`
    
    3. check for `/home/tribist/api/server_new` and if it exists `mv` it to `/home/tribist/api/server`
       
    4. run `/home/tribist/api/server`
    
    *restart*

    1. download the latest release artifact from github.
    
    2. if it does exist `mv` it to `/home/tribist/api/server_new`
    
    3. stop the server by sending `SIG_TERM` signal to server using the PID file.

    3. check for `/home/tribist/api/server_new` and if it exists `mv` it to `/home/tribist/api/server`
       
    4. run `/home/tribist/api/server`

    *stop*
    
    1. stop the server by sending `SIG_TERM` signal to server using PID file.
    
