#!/bin/bash
# Place this in the home directory of the user that is running
# twitter.sh command. (Check the unit file)
# This initializes nvm in the shell.
# Assumes that nvm has been installed globally
# and the nvm has installed nodejs version 16.20.0 has been installed
#
# $HOME is not set if user is not specified.
# if you specify a user in your systemd init file then
# please change this.
HOME="/root"

export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
