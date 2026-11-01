package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Remove a commit and replay descendants onto its parent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_dropCmd).Standalone()

	history_helpCmd.AddCommand(history_help_dropCmd)
}
