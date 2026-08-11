package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_metadata_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_metadata_getCmd).Standalone()

	revision_metadata_getCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_metadata_getCmd.Flags().String("revision", "", "Revision to get metadata for")
	revision_metadataCmd.AddCommand(revision_metadata_getCmd)
}
