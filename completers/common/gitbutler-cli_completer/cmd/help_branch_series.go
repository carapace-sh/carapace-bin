package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_seriesCmd = &cobra.Command{
	Use:   "series",
	Short: "Create a new series on top of the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_seriesCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_seriesCmd)
}
