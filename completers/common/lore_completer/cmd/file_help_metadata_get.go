package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_metadata_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_metadata_getCmd).Standalone()

	file_help_metadataCmd.AddCommand(file_help_metadata_getCmd)
}
