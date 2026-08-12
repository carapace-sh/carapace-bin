package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_metadata_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from the repository (omit key to list all)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_metadata_getCmd).Standalone()

	repository_help_metadataCmd.AddCommand(repository_help_metadata_getCmd)
}
