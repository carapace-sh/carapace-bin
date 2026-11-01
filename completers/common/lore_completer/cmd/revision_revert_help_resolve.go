package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_help_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve conflicts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_help_resolveCmd).Standalone()

	revision_revert_helpCmd.AddCommand(revision_revert_help_resolveCmd)
}
