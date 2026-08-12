package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Branch metadata operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_metadataCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_metadataCmd)
}
