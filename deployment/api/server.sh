#!/bin/bash

usage() {                                 # Function: Print a help message.
  echo "Usage: $0 [ -e ENV ] [ -p PATH ] -c [start|stop]" 1>&2
}

exit_abnormal() {                         # Function: Exit with error.
  usage
  exit 1
}

while getopts "c:g:" options; do
    case "${options}" in
        c)
            COMMAND=${OPTARG}
            ;;
        g)
            # NOT USED, moved to restart_server.sh
            GH_CONFIG_PATH=${OPTARG}
            ;;
        :)
            echo "Error: -${OPTARG} requires an argument."
            exit_abnormal                       # Exit abnormally.
            ;;
    esac
done

source .env

case "${COMMAND}" in
    start)
        if [ -e "PID" ]
        then
            kill -15 `cat "PID"`
            sleep 5
        fi

        curl "localhost:${PORT}"
        if [ $? -eq 0 ]
        then
            echo "Failed to shutdown server"
            exit_abnormal
        fi

        if [ -e "/home/cicd/.env" ]
        then
            mv /home/cicd/.env .
        fi

        if [ -e "/home/cicd/server_new" ]
        then
            mv /home/cicd/server_new .
        fi

        if [ -e "/home/cicd/db" ]
        then
            rm -rf db
            mv /home/cicd/db .
        fi

        # read the new downloaded env file
        source .env

        # update to the latest image only if UPDATE is not set
        # or it is "true"
        if [ "${UPDATE}x" = "x"  -o "${UPDATE}x" = "truex" ]
        then

            if [ -e "server_new" ]
            then
                mv "server_new" "server"
                chmod 500 server
            fi
        fi

        # run all migrations if MIGRATE is not set OR
        # it is set to true
        if [ "${MIGRATE}x" = "x"  -o "${MIGRATE}x" = "truex" ]
        then
            # run all the migrations
            ./server migrate --migrationsPath "db/migrations"
            if [ $? -ne 0 ]
            then
                echo "Failed migrations"
                exit 1
            fi
        fi

        ./server serve
    ;;
    stop)
        if [ -e "PID" ]
        then
            kill -15 `cat "PID"`
            sleep 5
        fi

        curl "localhost:${PORT}"
        if [ $? -eq 0 ]
        then
            echo "Failed to shutdown server"
            exit_abnormal
        fi
    ;;
esac
