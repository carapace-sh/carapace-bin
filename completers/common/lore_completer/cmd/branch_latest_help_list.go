package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_latest_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_latest_help_listCmd).Standalone()

	branch_latest_helpCmd.AddCommand(branch_latest_help_listCmd)
}
