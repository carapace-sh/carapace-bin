package cmd

import (
	"github.com/rsteube/carapace-bin/completers/grep_completer/cmd"
)

/**
Description for go:generate
	Use: "fgrep",
	Short: "print lines that match patterns",
*/

func Execute() error {
	return cmd.ExecuteOverride("fgrep")
}
