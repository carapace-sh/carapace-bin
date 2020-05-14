#!/bin/bash

IFS=$'\n'
export FILTER_BRANCH_SQUELCH_WARNING=1

for subcommand in $(git help -a -g | grep '^  '); do
    name=$(echo $subcommand | awk '{print $1}')
    if [[ "$name" == 'cvsserver' ]]; then
      echo "skipping: $subcommand" > /dev/stderr
      continue
    fi
    echo "generating: $subcommand" > /dev/stderr

    cmd=${name//-/_}
    description=$(echo $subcommand | cut -c25-)
    # convert flags like `--[no-]validate` to two lines with each variant 
    #s/^( +--)\[([^]]+)\](.*)/\1\2\3\t\1\3/
    #flags=$(git $name -h 2>&1 | grep '^    -' | sort | ./flags.sed | sed "s/^cmd/\t${cmd}Cmd/")
    git $name -h 2>&1 | tr '\n' '\t' | sed -r 's/\t                          /    /g' | tr '\t' '\n' | sed -r 's/^( +--)\[([^]]+)\](.*)/\1\2\3\t\1\3/' |tr '\t' '\n' | carapace -n "${cmd}" -p "rootCmd" -s "${description}" > "cmd/${cmd}_generated.go"
done
