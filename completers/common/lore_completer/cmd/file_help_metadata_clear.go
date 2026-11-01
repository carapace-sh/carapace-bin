package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_metadata_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata for a staged file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_metadata_clearCmd).Standalone()

	file_help_metadataCmd.AddCommand(file_help_metadata_clearCmd)
}
