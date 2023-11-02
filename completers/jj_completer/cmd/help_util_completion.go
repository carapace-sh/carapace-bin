package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_util_completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Print a command-line-completion script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_util_completionCmd).Standalone()

	help_utilCmd.AddCommand(help_util_completionCmd)
}
