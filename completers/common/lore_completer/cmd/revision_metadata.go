package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Manage metadata of a given revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_metadataCmd).Standalone()

	revision_metadataCmd.Flags().BoolP("help", "h", false, "Print help")
	revisionCmd.AddCommand(revision_metadataCmd)
}
