package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var step_relocateCmd = &cobra.Command{
	Use:   "relocate",
	Short: "[experimental] Move worktrees to expected paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_relocateCmd).Standalone()

	step_relocateCmd.Flags().Bool("clobber", false, "Backup non-worktree paths at target locations")
	step_relocateCmd.Flags().Bool("commit", false, "Commit uncommitted changes before relocating")
	step_relocateCmd.Flags().Bool("dry-run", false, "Show what would be moved")
	step_relocateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	stepCmd.AddCommand(step_relocateCmd)

	carapace.Gen(step_relocateCmd).PositionalAnyCompletion(
		wt.ActionBranches().FilterArgs(),
	)
}
