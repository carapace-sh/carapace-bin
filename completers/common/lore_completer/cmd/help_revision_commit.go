package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit the staged state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_commitCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_commitCmd)
}
