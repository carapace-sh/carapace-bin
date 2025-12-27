#!/bin/bash
set -e

echo "# Completers"
echo

carapace --list | jq --raw-output  '.[][] | select(.group != "bridge" and .group != "user" and .group != "system") | "- [\(.name)\(if .variant then "/\(.variant)" else "" end)\(if .group != "common" then "@\(.group)" else "" end)](\(.url))"'
