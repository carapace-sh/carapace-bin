package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize to a given state of a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_syncCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_syncCmd)
}
