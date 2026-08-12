package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_find_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Find revision by metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_find_metadataCmd).Standalone()

	help_revision_findCmd.AddCommand(help_revision_find_metadataCmd)
}
