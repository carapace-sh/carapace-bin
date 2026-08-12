package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff two revisions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_diffCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_diffCmd)
}
