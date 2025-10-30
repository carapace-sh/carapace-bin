package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_unapplyCmd = &cobra.Command{
	Use:   "unapply",
	Short: "Unapply a branch from the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_unapplyCmd).Standalone()

	branch_unapplyCmd.Flags().BoolP("force", "f", false, "Force unapply without confirmation")
	branch_unapplyCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_unapplyCmd)
}
