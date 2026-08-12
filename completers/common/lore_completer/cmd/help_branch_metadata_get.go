package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_metadata_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get metadata from the branch (omit key to list all)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_metadata_getCmd).Standalone()

	help_branch_metadataCmd.AddCommand(help_branch_metadata_getCmd)
}
