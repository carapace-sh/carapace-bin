package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_metadata_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on for a staged file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_metadata_setCmd).Standalone()

	file_help_metadataCmd.AddCommand(file_help_metadata_setCmd)
}
