package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:     "reset",
	Short:   "Move the branch tip and optionally reset the index or worktree",
	GroupID: "branches",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetCmd).Standalone()

	resetCmd.Flags().Bool("dry-run", false, "Preview the reset without updating refs, the index, or the worktree")
	resetCmd.Flags().Bool("hard", false, "Reset the index and worktree to match the target commit")
	resetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	resetCmd.Flags().Bool("keep", false, "Reset the tip and update differing paths while keeping unrelated local edits")
	resetCmd.Flags().StringP("limit", "n", "", "Maximum orphaned commits to list in `--json` output")
	resetCmd.Flags().Bool("mixed", false, "Reset the index to the target commit; keep the worktree (default)")
	resetCmd.Flags().Bool("soft", false, "Keep the index and worktree; move the branch tip only")
	rootCmd.AddCommand(resetCmd)

	carapace.Gen(resetCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
