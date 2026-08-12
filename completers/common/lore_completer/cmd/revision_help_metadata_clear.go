package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_metadata_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata for a staged revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_metadata_clearCmd).Standalone()

	revision_help_metadataCmd.AddCommand(revision_help_metadata_clearCmd)
}
