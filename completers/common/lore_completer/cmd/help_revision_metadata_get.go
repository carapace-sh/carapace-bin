package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_metadata_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_metadata_getCmd).Standalone()

	help_revision_metadataCmd.AddCommand(help_revision_metadata_getCmd)
}
