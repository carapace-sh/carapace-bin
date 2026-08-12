package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_find_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Find revision by metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_find_metadataCmd).Standalone()

	revision_help_findCmd.AddCommand(revision_help_find_metadataCmd)
}
