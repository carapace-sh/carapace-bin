package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Peel path-limited changes out of a commit into a new parent and replay descendants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_splitCmd).Standalone()

	history_helpCmd.AddCommand(history_help_splitCmd)
}
