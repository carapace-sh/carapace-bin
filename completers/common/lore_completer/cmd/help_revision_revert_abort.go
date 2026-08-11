package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_revert_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a revert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_revert_abortCmd).Standalone()

	help_revision_revertCmd.AddCommand(help_revision_revert_abortCmd)
}
