package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Provide details about given branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_detailsCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_detailsCmd)
}
