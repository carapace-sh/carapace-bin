package cmd

import (
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch branches",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	switchCmd.Flags().BoolP("create", "c", false, "<branch>    create and switch to a new branch")
	switchCmd.Flags().BoolP("force-create", "C", false, "<branch>    create/reset and switch to a branch")
	switchCmd.Flags().String("conflict", "", "conflict style (merge or diff3)")
	switchCmd.Flags().BoolP("detach", "d", false, "detach HEAD at named commit")
	switchCmd.Flags().Bool("discard-changes", false, "throw away local modifications")
	switchCmd.Flags().BoolP("force", "f", false, "force checkout (throw away local modifications)")
	switchCmd.Flags().Bool("guess", false, "second guess 'git switch <no-such-branch>'")
	switchCmd.Flags().Bool("ignore-other-worktrees", false, "do not check if another worktree is holding the given ref")
	switchCmd.Flags().BoolP("merge", "m", false, "perform a 3-way merge with the new branch")
	switchCmd.Flags().String("orphan", "", "new unparented branch")
	switchCmd.Flags().Bool("overwrite-ignore", false, "update ignored files (default)")
	switchCmd.Flags().Bool("progress", false, "force progress reporting")
	switchCmd.Flags().BoolP("quiet", "q", false, "suppress progress reporting")
	switchCmd.Flags().String("recurse-submodules", "", "control recursive updating of submodules")
	switchCmd.Flags().BoolP("track", "t", false, "set upstream info for new branch")
	rootCmd.AddCommand(switchCmd)
}
