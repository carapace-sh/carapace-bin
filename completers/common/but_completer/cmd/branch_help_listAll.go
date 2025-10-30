package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_listAllCmd = &cobra.Command{
	Use:   "list-all",
	Short: "List all branches that can be relevant",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_listAllCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_listAllCmd)
}
