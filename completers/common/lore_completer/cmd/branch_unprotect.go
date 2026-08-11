package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_unprotectCmd = &cobra.Command{
	Use:   "unprotect",
	Short: "Remove push protection from a branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_unprotectCmd).Standalone()

	branch_unprotectCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_unprotectCmd)
}
