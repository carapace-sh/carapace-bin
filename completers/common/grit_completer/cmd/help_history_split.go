package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Peel path-limited changes out of a commit into a new parent and replay descendants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_splitCmd).Standalone()

	help_historyCmd.AddCommand(help_history_splitCmd)
}
