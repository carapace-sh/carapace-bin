package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_help_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a revert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_help_abortCmd).Standalone()

	revision_revert_helpCmd.AddCommand(revision_revert_help_abortCmd)
}
