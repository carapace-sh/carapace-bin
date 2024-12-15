#!/bin/bash
set -e pipefail

cd "$(dirname "$(readlink -f "$0")")/../book/html/release_notes/"
ls *.html \
 | sed 's/.html$//' \
 | xargs -I'{}' \
         sed --in-place \
             '/<head>/a <meta property="og:image" content="https://carapace-sh.github.io/carapace-bin/release_notes/{}/banner.png" />' \
             '{}.html'
