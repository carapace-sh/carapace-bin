package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert a revision from the currently synced revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_revertCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_revertCmd)
}
