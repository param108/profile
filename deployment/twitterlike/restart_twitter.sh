#!/bin/bash -x

systemctl stop twitter
sleep 5
systemctl start twitter
