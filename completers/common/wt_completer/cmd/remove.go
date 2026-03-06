package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wt"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove worktree; delete branch if merged",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().BoolP("force", "f", false, "Force worktree removal")
	removeCmd.Flags().BoolP("force-delete", "D", false, "Delete unmerged branches")
	removeCmd.Flags().Bool("foreground", false, "Run removal in foreground (block until complete)")
	removeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	removeCmd.Flags().Bool("no-delete-branch", false, "Keep branch after removal")
	removeCmd.Flags().Bool("no-verify", false, "Skip hooks")
	removeCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		wt.ActionWorktrees().FilterArgs(),
	)
}
