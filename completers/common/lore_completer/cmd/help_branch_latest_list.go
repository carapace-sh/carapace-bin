package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_latest_listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_latest_listCmd).Standalone()

	help_branch_latestCmd.AddCommand(help_branch_latest_listCmd)
}
