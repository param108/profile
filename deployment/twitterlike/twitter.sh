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


	/usr/bin/gh run -R param108/profile download ${IMAGE_ID} -n twitter.tgz -D /tmp/

	if [ $? -ne 0 ]
	then
		echo "failed download"
		exit 1
	fi

	# extract package
	pushd /tmp/
	tar -zxvf /tmp/twitter.tgz
	popd

	mv /tmp/twitterlike  twitterlike_new
    cp twitterlike/PID twitterlike_new/
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

        if [ -e "twitterlike_new/PID" ]
        then
            kill -15 `cat "twitterlike_new/PID"`
            sleep 5
        fi

        curl "localhost:${PORT}"
        if [ $? -eq 0 ]
        then
            echo "Failed to shutdown server"
            exit_abnormal
        fi

        # source the newly downloaded env file
        source twitterlike_new/.env

        # update to the latest image only if UPDATE is not set
        # or it is "true"
        if [ "${UPDATE}x" = "x"  -o "${UPDATE}x" = "truex" ]
        then

            if [ -e "twitterlike_new" ]
            then
                mv twitterlike_new twitterlike
            fi
        fi

        cd twitterlike
        # create the pid file
        echo $$ > PID

        yarn run start
    ;;
    stop)

        cd twitterlike
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
