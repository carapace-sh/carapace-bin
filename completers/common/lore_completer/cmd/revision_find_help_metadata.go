package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_find_help_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Find revision by metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_find_help_metadataCmd).Standalone()

	revision_find_helpCmd.AddCommand(revision_find_help_metadataCmd)
}
