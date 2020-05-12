#!/bin/sh

exa --help | sed -r 's/^(.*)colo\[u\]r(.*)/\1color\2\n\1colour\2/' | carapace -n exa > cmd/root.go
