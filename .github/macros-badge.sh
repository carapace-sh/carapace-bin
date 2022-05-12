#!/bin/sh

scriptdir=$(dirname $(readlink -f $0))

curl "https://img.shields.io/badge/macros-$($scriptdir/../cmd/carapace/carapace --macros | wc -l)-orange"
