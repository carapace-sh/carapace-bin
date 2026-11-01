package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore current revision as latest revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_restoreCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_restoreCmd)
}
