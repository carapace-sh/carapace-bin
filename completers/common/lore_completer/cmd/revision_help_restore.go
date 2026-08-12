package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore current revision as latest revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_restoreCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_restoreCmd)
}
