package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff two revisions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_diffCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_diffCmd)
}
