#!/bin/sh
set -e

echo "# Macros"
echo "On top of the base [Macros](https://carapace-sh.github.io/carapace-spec/carapace-spec/macros.html) defined in [carapace-spec](https://carapace-sh.github.io/carapace-spec),"
echo "carapace-bin also provides the following macros:"
echo

carapace --macro \
  | sed -r 's_(https://[^ ]+)_[\1](\1)_' \
  | sed 's_^\([^ ]*\.\)\([^. ]\+\) \+\(.*\)_- [\1\2]\(https://pkg.go.dev/github.com/carapace-sh/carapace-bin/pkg/actions/\1#Action\2) \3_' \
  | sed -e ':loop' -e 's_\(carapace-bin/pkg/actions/[^#]*\)[.]_\1/_' -e 't loop' \
  | sed 's_https://pkg.go.dev/github.com/carapace-sh/carapace-bin/pkg/actions/bridge_https://pkg.go.dev/github.com/carapace-sh/carapace-bridge/pkg/actions/bridge_'
