package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_metadata_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from the repository (omit key to list all)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_metadata_getCmd).Standalone()

	help_repository_metadataCmd.AddCommand(help_repository_metadata_getCmd)
}
