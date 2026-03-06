package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var step_copyIgnoredCmd = &cobra.Command{
	Use:   "copy-ignored",
	Short: "Copy gitignored files to another worktree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_copyIgnoredCmd).Standalone()

	step_copyIgnoredCmd.Flags().Bool("dry-run", false, "Show what would be copied")
	step_copyIgnoredCmd.Flags().Bool("force", false, "Overwrite existing files in destination")
	step_copyIgnoredCmd.Flags().String("from", "", "Source worktree branch")
	step_copyIgnoredCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	step_copyIgnoredCmd.Flags().String("to", "", "Destination worktree branch")
	stepCmd.AddCommand(step_copyIgnoredCmd)

	carapace.Gen(step_copyIgnoredCmd).FlagCompletion(carapace.ActionMap{
		"from": wt.ActionWorktrees(),
		"to":   wt.ActionWorktrees(),
	})
}
