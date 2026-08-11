package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset local latest pointer for a branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_resetCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_resetCmd)
}
