package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_metadata_help_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata for a staged revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_metadata_help_clearCmd).Standalone()

	revision_metadata_helpCmd.AddCommand(revision_metadata_help_clearCmd)
}
