package cmd

import (
	"github.com/rsteube/carapace-bin/completers/code_completer/cmd"
)

/**
Description for go:generate
	Short: "Visual Studio Code Insiders",
*/

func Execute() error {
	return cmd.ExecuteInsiders()
}
