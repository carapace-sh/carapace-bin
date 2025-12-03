#!/bin/sh
set -e

amount=$(carapace --list | jq '[ .[][] | select(.group != "bridge" and .group != "user" and .group != "system") ] | length')

curl "https://img.shields.io/badge/completers-${amount}-orange"
