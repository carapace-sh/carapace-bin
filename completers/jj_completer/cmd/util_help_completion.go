package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var util_help_completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Print a command-line-completion script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_help_completionCmd).Standalone()

	util_helpCmd.AddCommand(util_help_completionCmd)
}
