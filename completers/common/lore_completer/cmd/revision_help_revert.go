package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert a revision from the currently synced revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_revertCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_revertCmd)
}
