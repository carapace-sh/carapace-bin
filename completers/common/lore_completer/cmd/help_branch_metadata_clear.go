package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_metadata_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear metadata from the branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_metadata_clearCmd).Standalone()

	help_branch_metadataCmd.AddCommand(help_branch_metadata_clearCmd)
}
