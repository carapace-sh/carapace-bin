package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_unapplyCmd = &cobra.Command{
	Use:   "unapply",
	Short: "Remove a branch from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_unapplyCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_unapplyCmd)
}
