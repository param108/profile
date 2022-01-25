#!/bin/bash

usage() {                                 # Function: Print a help message.
  echo "Usage: $0 [ -e ENV ] [ -p PATH ] -c [restart|start|stop]" 1>&2
}

exit_abnormal() {                         # Function: Exit with error.
  usage
  exit 1
}

while getopts "e:p:c:" options; do
    case "${options}" in
        e)
            ENV=${OPTARG}
            ;;
        p)
            EXECPATH=${OPTARG}
            ;;
        c)
            COMMAND=${OPTARG}
            ;;
        :)
            echo "Error: -${OPTARG} requires an argument."
            exit_abnormal                       # Exit abnormally.
            ;;
    esac
done

source ${ENV}

case "${COMMAND}" in
    restart)
        if [ -e "${EXECPATH}/PID" ]
        then
            kill -15 `cat "${EXECPATH}/PID"`
            sleep 5
        fi

        curl "localhost:${PORT}"
        if [ $? -eq 0 ]
        then
            echo "Failed to shutdown server"
            exit_abnormal
        fi

        if [ -e "${EXECPATH}/server_new" ]
        then
            mv "${EXECPATH}/server_new" "${EXECPATH}/server"
        fi

        cd ${EXECPATH}
        ./server
    ;;
    start)
        if [ -e "${EXECPATH}/PID" ]
        then
            kill -15 `cat "${EXECPATH}/PID"`
            sleep 5
        fi

        curl "localhost:${PORT}"
        if [ $? -eq 0 ]
        then
            echo "Failed to shutdown server"
            exit_abnormal
        fi

        if [ -e "${EXECPATH}/server_new" ]
        then
            mv "${EXECPATH}/server_new" "${EXECPATH}/server"
        fi

        cd ${EXECPATH}
        ./server
    ;;
    stop)
        if [ -e "${EXECPATH}/PID" ]
        then
            kill -15 `cat "${EXECPATH}/PID"`
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
