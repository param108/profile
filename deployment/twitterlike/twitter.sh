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
            # moved to restart_twitter.sh
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

        if [ -e "/home/cicd/twitterlike_new" ]
        then
            mv /home/cicd/twitterlike_new .
        fi

        if [ -e twitterlike_new ]
        then
            # source the newly downloaded env file
            source twitterlike_new/.env.local
        else
            source twitterlike/.env.local
        fi

        if [ -e "twitterlike/PID" ]
        then
	        pkill -9 -P `cat "twitterlike/PID"`
	        kill -9 `cat "twitterlike/PID"`
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
