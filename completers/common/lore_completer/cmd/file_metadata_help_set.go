package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_metadata_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on for a staged file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_metadata_help_setCmd).Standalone()

	file_metadata_helpCmd.AddCommand(file_metadata_help_setCmd)
}
