#!/bin/bash -x
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
}

GH_CONFIG_PATH=/home/cicd/gh.txt

download_latest_image

systemctl stop twitter
sleep 5
systemctl start twitter
