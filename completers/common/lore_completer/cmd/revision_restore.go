package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore current revision as latest revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_restoreCmd).Standalone()

	revision_restoreCmd.Flags().BoolP("help", "h", false, "Print help")
	revisionCmd.AddCommand(revision_restoreCmd)
}
