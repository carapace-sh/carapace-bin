package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branch_unprotectCmd = &cobra.Command{
	Use:   "unprotect",
	Short: "Remove push protection from a branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branch_unprotectCmd).Standalone()

	help_branchCmd.AddCommand(help_branch_unprotectCmd)
}
