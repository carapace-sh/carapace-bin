package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Manage metadata of a given revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_metadataCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_metadataCmd)
}
