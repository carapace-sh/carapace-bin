package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktreeCmd = &cobra.Command{
	Use:     "worktree",
	Short:   "Manage worktrees",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "branching and committing",
}

func init() {
	carapace.Gen(worktreeCmd).Standalone()

	worktreeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(worktreeCmd)
}
