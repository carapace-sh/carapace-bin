#!/bin/sh

scriptdir=$(dirname $(readlink -f $0))

curl "https://img.shields.io/badge/macros-$($scriptdir/../cmd/carapace/carapace --macro | wc -l)-orange"
