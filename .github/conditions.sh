#!/bin/sh
set -e

echo "# Conditions"
echo

carapace --condition \
  | sed -r 's_(https://[^ ]+)_[\1](\1)_' \
  | sed 's_^\([^ ]\+\) \+\(.*\)_- [\1]\(https://pkg.go.dev/github.com/carapace-sh/carapace-bin/pkg/conditions/#Condition\1) \2_' \
  | sed -e ':loop' -e 's_\(carapace-bin/pkg/condition/[^#]*\)[.]_\1/_' -e 't loop'
