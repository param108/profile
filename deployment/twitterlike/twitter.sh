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
	rm -rf twitterlike_new
	rm -rf /tmp/twitterlike
	rm  -f /tmp/twitter.tgz

	/usr/bin/gh auth login --with-token < ${GH_CONFIG_PATH}

	if [ $? -ne 0 ]
	then
		echo "failed login"
		exit 1
	fi

	# download artifact
	IMAGE_ID=`/usr/bin/gh run -R param108/profile list -w twitterlike --json conclusion,databaseId,workflowDatabaseId -L 1 -q 'select(.[].conclusion = "success")' | jq .[0].databaseId`

	echo "DOWNLOADING......"

	/usr/bin/gh run -R param108/profile download ${IMAGE_ID} -n twitter.tgz -D /tmp/

	DOWNLOAD_RC=$?

	echo "Download RC - ${DOWNLOAD_RC}"

	if [ ${DOWNLOAD_RC} -ne 0 ]
	then
		echo "failed download"
		exit 1
	fi
	echo "DOWNLOAD Done"

	# extract package
	pushd /tmp/
	tar -zxf /tmp/twitter.tgz
	popd

	mv /tmp/twitterlike  twitterlike_new
	if [ -e twitterlike/PID ]
	then
		cp twitterlike/PID twitterlike_new/
	fi
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


case "${COMMAND}" in
    start)

        download_latest_image

	echo "Downloaded latest image"

        # source the newly downloaded env file
        source twitterlike_new/.env.local

        if [ -e "twitterlike_new/PID" ]
        then
	    pkill -9 -P `cat "twitterlike_new/PID"`
	    kill -9 `cat "twitterlike_new/PID"`
            sleep 5
        fi

	echo "killed old twitter"
	echo "PORT: ${PORT}"

        curl -s "localhost:${PORT}" > /dev/null
        if [ $? -eq 0 ]
        then
            echo "Failed to shutdown server"
            exit_abnormal
        fi

	echo "checked killed old twitter"
        # update to the latest image only if UPDATE is not set
        # or it is "true"
        if [ "${UPDATE}x" = "x"  -o "${UPDATE}x" = "truex" ]
        then

            if [ -e "twitterlike_new" ]
            then
		rm -rf twitterlike
                mv twitterlike_new twitterlike
            fi
        fi

	echo "installed new twitter"
        cd twitterlike
        # create the pid file
        echo $$ > PID

	echo "HOME $HOME"
	source /root/nvm_init.sh
	cat /root/nvm_init.sh
	nvm use 16.20.0
	echo "starting new twitter....."
        yarn run start -p ${PORT}
    ;;
    stop)

        cd twitterlike
        source .env.local

        if [ -e "PID" ]
        then
	    pkill -9 -P `cat "PID"`
        fi
    ;;
esac
