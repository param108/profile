#!/bin/bash

usage() {                                 # Function: Print a help message.
  echo "Usage: $0 [ -e ENV ] [ -p PATH ] -c [start|stop]" 1>&2
}

exit_abnormal() {                         # Function: Exit with error.
  usage
  exit 1
}

download_latest_image() {
	# clean up
	rm -rf /tmp/build
	rm  -f /tmp/server.tgz

	/usr/bin/gh auth login --with-token < ${GH_CONFIG_PATH}

	if [ $? -ne 0 ]
	then
		echo "failed login"
		exit 1
	fi

	# download artifact
	IMAGE_ID=`/usr/bin/gh run -R param108/profile list -w api_deploy --json conclusion,databaseId,workflowDatabaseId -L 1 -q 'select(.[].conclusion = "success")' | jq .[0].databaseId`

	rm /tmp/server

	/usr/bin/gh run -R param108/profile download ${IMAGE_ID} -n server.tgz -D /tmp/

	if [ $? -ne 0 ]
	then
		echo "failed download"
		exit 1
	fi

	# extract package
	pushd /tmp/
	tar -zxvf /tmp/server.tgz
	popd

	mv /tmp/build/server  server_new
	mv /tmp/build/env .env

    rm -rf db
    mv /tmp/build/db .
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
