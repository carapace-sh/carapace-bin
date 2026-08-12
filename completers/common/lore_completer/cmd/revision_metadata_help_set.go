package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_metadata_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set metadata on for a staged revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_metadata_help_setCmd).Standalone()

	revision_metadata_helpCmd.AddCommand(revision_metadata_help_setCmd)
}
