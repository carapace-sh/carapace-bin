package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Manage metadata of a given revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_metadataCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_metadataCmd)
}
