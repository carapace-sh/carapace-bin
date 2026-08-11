package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_metadata_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_metadata_getCmd).Standalone()

	help_file_metadataCmd.AddCommand(help_file_metadata_getCmd)
}
