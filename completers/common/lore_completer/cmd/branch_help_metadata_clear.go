package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_metadata_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata from the branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_metadata_clearCmd).Standalone()

	branch_help_metadataCmd.AddCommand(branch_help_metadata_clearCmd)
}
