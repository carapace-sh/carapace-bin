package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_abortCmd = &cobra.Command{
	Use:   "abort",
	Short: "Abort a revert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_abortCmd).Standalone()

	revision_revert_abortCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_revertCmd.AddCommand(revision_revert_abortCmd)
}
