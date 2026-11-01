package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_metadata_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on for a staged revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_metadata_setCmd).Standalone()

	revision_help_metadataCmd.AddCommand(revision_help_metadata_setCmd)
}
