package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Branch metadata operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_metadataCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_metadataCmd)
}
