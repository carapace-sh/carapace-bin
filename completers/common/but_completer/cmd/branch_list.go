package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the branches in the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_listCmd).Standalone()

	branch_listCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_listCmd.Flags().BoolP("local", "l", false, "Show only local branches")
	branchCmd.AddCommand(branch_listCmd)
}
