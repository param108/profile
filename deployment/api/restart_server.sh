#!/bin/bash -x

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
    IMAGE_ID=`/usr/bin/gh run -R param108/profile list -w api_test --json conclusion,databaseId,workflowDatabaseId | jq '[.[]|select(.conclusion=="success")][0] | .databaseId'`

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

GH_CONFIG_PATH=/home/cicd/gh.txt

# wait 10 seconds for the image to uploaded on github
sleep 10

download_latest_image

systemctl stop tribist
sleep 5
systemctl start tribist
