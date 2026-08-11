package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_listCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_listCmd)
}
