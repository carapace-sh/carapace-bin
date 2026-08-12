package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_findCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_findCmd)
}
