package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_metadata_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata for a staged revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_metadata_clearCmd).Standalone()

	help_revision_metadataCmd.AddCommand(help_revision_metadata_clearCmd)
}
