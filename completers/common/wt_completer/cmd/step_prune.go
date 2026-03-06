package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var step_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "[experimental] Remove worktrees merged into the default branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(step_pruneCmd).Standalone()

	step_pruneCmd.Flags().Bool("dry-run", false, "Show what would be removed")
	step_pruneCmd.Flags().Bool("foreground", false, "Run removal in foreground (block until complete)")
	step_pruneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	step_pruneCmd.Flags().String("min-age", "", "Skip worktrees younger than this")
	step_pruneCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	stepCmd.AddCommand(step_pruneCmd)
}
