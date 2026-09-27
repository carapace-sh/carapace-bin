package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Print shell completion code to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_completionCmd).Standalone()

	helpCmd.AddCommand(help_completionCmd)
}
