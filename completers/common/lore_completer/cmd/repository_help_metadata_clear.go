package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_metadata_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata from the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_metadata_clearCmd).Standalone()

	repository_help_metadataCmd.AddCommand(repository_help_metadata_clearCmd)
}
