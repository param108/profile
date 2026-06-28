#!/bin/bash
# Script to update or add environment variables in a file without exposing secrets
# Usage: ./update_env.sh <env_file> <KEY> <VALUE>
#
# Examples:
#   ./update_env.sh env.prod HOST profile-api.pritiatma.in
#   ./update_env.sh env.prod FRONTEND_URL https://profile.pritiatma.in

set -e

ENV_FILE="$1"
KEY="$2"
VALUE="$3"

if [ -z "$ENV_FILE" ] || [ -z "$KEY" ] || [ -z "$VALUE" ]; then
    echo "Usage: $0 <env_file> <KEY> <VALUE>"
    exit 1
fi

if [ ! -f "$ENV_FILE" ]; then
    echo "Error: File '$ENV_FILE' not found"
    exit 1
fi

# Check if key exists in file (without showing the value)
if grep -q "^${KEY}=" "$ENV_FILE"; then
    # Key exists - update it using sed
    # Use a delimiter that won't appear in URLs (|)
    sed -i "s|^${KEY}=.*|${KEY}=${VALUE}|" "$ENV_FILE"
    echo "Updated: ${KEY}"
else
    # Key doesn't exist - append it
    echo "${KEY}=${VALUE}" >> "$ENV_FILE"
    echo "Added: ${KEY}"
fi
