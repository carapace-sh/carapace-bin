package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_metadata_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_metadata_getCmd).Standalone()

	file_metadata_getCmd.Flags().BoolP("help", "h", false, "Print help")
	file_metadata_getCmd.Flags().String("revision", "", "Revision to get metadata for")
	file_metadataCmd.AddCommand(file_metadata_getCmd)
}
