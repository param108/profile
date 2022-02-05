# Security

## CICD

1. CICD only has ssh access to the user `cicd` who can only run `/usr/bin/restart_server.sh` which can only call `sudo systemctl tribist stop` and then `sudo systemctl tribist start`
2. These commands call `server.sh` with appropriate arguments.

## Database secrets

1. The database is a local instance and is not available on the public ip address.
2. the password is passed in the env file via the secret `API_ENV_CONFIG`
   - This config is available to collaborators but the only ssh key available is for `cicd` above.
   - As a result, even if someone gets the ssh key for cicd they can't ssh into the system and hence cannot access the DB. So noone except the maintainer can have direct access to the DB.

