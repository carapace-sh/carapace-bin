package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_metadata_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_metadata_getCmd).Standalone()

	revision_help_metadataCmd.AddCommand(revision_help_metadata_getCmd)
}
