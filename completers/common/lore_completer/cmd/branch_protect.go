package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branch_protectCmd = &cobra.Command{
	Use:   "protect",
	Short: "Protect a branch from direct pushes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_protectCmd).Standalone()

	branch_protectCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_protectCmd)

	carapace.Gen(branch_protectCmd).PositionalCompletion(
		action.ActionBranches(branch_protectCmd),
	)
}
