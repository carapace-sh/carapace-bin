#!/bin/bash

scriptdir=$(dirname $(readlink -f $0))

echo "# Completers"
echo

for file in $(ls -1 $scriptdir/../completers/*/*/cmd/root.go); do
  use=$(grep --max-count=1 "	Use:" $file | sed 's/[^"]\+"\(.*\)",/\1/' | awk '{print $1}')
  short=$(grep --max-count=1 "	Short:" $file | sed 's/[^"]\+"\(.*\)",/\1/')
  long=$(grep --max-count=1 "	Long:" $file | sed 's/[^"]\+"\(.*\)",/\1/')

  echo "- [$use]($long) $short"
done



