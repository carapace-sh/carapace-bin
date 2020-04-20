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
    flags=$(git $name -h 2>&1 | tr '\n' '\t' | sed -r 's/\t                          /    /g' | tr '\t' '\n' | sed -r 's/^( +--)\[([^]]+)\](.*)/\1\2\3\t\1\3/' |tr '\t' '\n' | grep '^    -' | sort | ./flags.sed | sed "s/^cmd/\t${cmd}Cmd/")
    echo "\
package cmd

import (
	\"github.com/spf13/cobra\"
)

var ${cmd}Cmd = &cobra.Command{
	Use: \"${name}\",
	Short: \"${description}\",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
$flags
    rootCmd.AddCommand(${cmd}Cmd)
}" > "cmd/${cmd}_generated.go"
done
