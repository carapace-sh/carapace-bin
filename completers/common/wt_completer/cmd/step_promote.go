package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var step_promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "[experimental] Put a branch into the main worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_promoteCmd).Standalone()

	step_promoteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stepCmd.AddCommand(step_promoteCmd)

	carapace.Gen(step_promoteCmd).PositionalCompletion(
		wt.ActionBranches(),
	)
}
