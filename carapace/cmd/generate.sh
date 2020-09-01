#!/bin/bash

completers=$(cd ../completers && ls -d * | sed 's/_completer//')

echo "package cmd

import ("

for completer in $completers; do
    echo "	${completer//-/_} \"github.com/rsteube/carapace-bin/completers/${completer}_completer/cmd\""
done

echo ")

var completers = []string{"

for completer in $completers; do
  echo "	\"$completer\","
done

echo "}

func executeCompleter(completer string) {
	switch completer {"


for completer in $completers; do
    echo -e "	case \"${completer}\":\n		${completer//-/_}.Execute()"
done

echo "	}
}
"
