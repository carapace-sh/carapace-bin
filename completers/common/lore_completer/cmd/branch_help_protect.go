package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_protectCmd = &cobra.Command{
	Use:   "protect",
	Short: "Protect a branch from direct pushes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_protectCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_protectCmd)
}
