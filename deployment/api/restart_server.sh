#!/bin/bash -x

systemctl stop tribist
sleep 5
systemctl start tribist
