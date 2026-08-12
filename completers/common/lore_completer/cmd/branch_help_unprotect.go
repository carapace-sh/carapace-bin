package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_unprotectCmd = &cobra.Command{
	Use:   "unprotect",
	Short: "Remove push protection from a branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_unprotectCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_unprotectCmd)
}
