#!/bin/sh
set -e

curl "https://img.shields.io/badge/macros-$(carapace --macro | wc -l)-orange"
