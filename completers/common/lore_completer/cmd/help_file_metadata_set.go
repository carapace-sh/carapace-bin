package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_metadata_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on for a staged file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_metadata_setCmd).Standalone()

	help_file_metadataCmd.AddCommand(help_file_metadata_setCmd)
}
