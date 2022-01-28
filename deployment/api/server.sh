#!/bin/bash

usage() {                                 # Function: Print a help message.
  echo "Usage: $0 [ -e ENV ] [ -p PATH ] -c [restart|start|stop]" 1>&2
}

exit_abnormal() {                         # Function: Exit with error.
  usage
  exit 1
}

download_latest_image() {
	/usr/bin/gh auth login --with-token < ${GH_CONFIG_PATH}

	if [ $? -ne 0 ]
	then
		echo "failed login"
		exit 1
	fi

	IMAGE_ID=`/usr/bin/gh run -R param108/profile list -w api_deploy --json conclusion,databaseId,workflowDatabaseId -L 1 -q 'select(.[].conclusion = "success")' | jq .[0].databaseId`

	rm /tmp/server

	/usr/bin/gh run -R param108/profile download ${IMAGE_ID} -n server -D /tmp/

	if [ $? -ne 0 ]
	then
		echo "failed download"
		exit 1
	fi

	mv /tmp/server  server_new

}

while getopts "c:g:" options; do
    case "${options}" in
        c)
            COMMAND=${OPTARG}
            ;;
        g)
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
    reload)
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

        download_latest_image

        if [ -e "server_new" ]
        then
            mv "server_new" "server"
           chmod 500 server
        fi

        ./server
    ;;
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

        download_latest_image

        if [ -e "server_new" ]
        then
            mv "server_new" "server"
            chmod 500 server
        fi

        ./server
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
