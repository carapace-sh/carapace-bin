package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_metadata_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata from the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_metadata_clearCmd).Standalone()

	help_repository_metadataCmd.AddCommand(help_repository_metadata_clearCmd)
}
