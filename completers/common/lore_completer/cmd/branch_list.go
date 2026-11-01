package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_listCmd).Standalone()

	branch_listCmd.Flags().Bool("archived", false, "Include archived local branches")
	branch_listCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_listCmd)
}
