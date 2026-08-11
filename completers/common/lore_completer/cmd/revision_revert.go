package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert a revision from the currently synced revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revertCmd).Standalone()

	revision_revertCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_revertCmd.Flags().String("message", "", "Change the message for committing when no conflicts arise from the revert")
	revision_revertCmd.Flags().Bool("no-commit", false, "Disable auto commits even if no conflicts arise from the revert")
	revisionCmd.AddCommand(revision_revertCmd)
}
