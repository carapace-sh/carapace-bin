package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branch_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset local latest pointer for a branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_resetCmd).Standalone()

	branch_resetCmd.Flags().String("branch", "", "Branch to reset, or the current branch if not set")
	branch_resetCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_resetCmd)

	carapace.Gen(branch_resetCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(branch_resetCmd),
	})

	carapace.Gen(branch_resetCmd).PositionalCompletion(
		action.ActionRevisions(branch_resetCmd),
	)
}
