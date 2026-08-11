package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Amend the latest commit's message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_amendCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_amendCmd)
}
