package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Add a branch to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_applyCmd).Standalone()

	branch_applyCmd.Flags().BoolP("branch", "b", false, "Whether it's a branch that we're applying")
	branch_applyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branchCmd.AddCommand(branch_applyCmd)
}
