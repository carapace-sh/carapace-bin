package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var branch_help_forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget everything about a branch, including its local and remote targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_forgetCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_forgetCmd)
}
