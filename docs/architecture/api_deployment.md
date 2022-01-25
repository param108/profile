# API server Deployment

We are choosing to avoid the overhead of docker for now.
We will use systemd to keep the executable running if it crashes.

# Prerequisites

1. 
1. The server binary must listen for SIG_TERM and gracefully shutdown.

# Overview

## Installation

1. ssh into server and copy the `systemd` unit file to `/etc/systemd/system/tribist.service`
   (Not `user` as we want this service to always run.)
2. place the server in the appropriate directory with appropriate permissions
   `/home/tribist/api/server`
3. when run, the executable must create a file `PID` in the current directory
   with the process' current PID. 
   This will be used to send the process `SIG_TERM` when we wish to restart.
4. a shell script `server.sh` will be placed at `/home/tribist/bin/server.sh` which will
   be used to start and stop the server.
5. check for successful start will use the script `wait_for_it.sh` and will fail in 30 seconds.
   We will hit `http://localhost:8383/health` which will also check the DB connection.
   
   The unit file is responsible to restart the server in case of failure.
6. `sudo systemctl daemon-reload`
7. `sudo systemctl start tribist`

## Upgrade
1. scp `server` binary to `/home/tribist/api/server_new`
1. ssh into the server
2. execute `sudo systemctl restart tribist`
   
   `server.sh` is responsible to check for `server_new` existing and moving it to `server`
3. wait for `wait_for_it.sh` to return.
   
   if `wait_for_it.sh` fails print appropriate message and exit.
   
# Details 

## server.sh
    `$1` is one of `[stop|start|restart]`
    
    *start*
    
    1. check if `home/tribist/api/server_new` exists.
       
       if it does `mv` it to `/home/tribist/api/server`
       
    2. `sudo systemctl start tribist`
    
    *restart*

    1. check if `home/tribist/api/server_new` exists.
       
       if it does `mv` it to `/home/tribist/api/server`
       
    2. `sudo systemctl restart tribist`

    *stop*
    
    1. `sudo systemctl stop tribist`
    
