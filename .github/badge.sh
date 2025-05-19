#!/bin/sh

scriptdir=$(dirname $(readlink -f $0))

curl "https://img.shields.io/badge/completers-$(ls $scriptdir/../completers/* | wc -l)-orange"
