package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Add a branch to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_applyCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_applyCmd)
}
