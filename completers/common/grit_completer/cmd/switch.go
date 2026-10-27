package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:     "switch",
	Short:   "Switch to another branch, remote branch, or detached commit",
	GroupID: "branches",
	Aliases: []string{"checkout", "co"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchCmd).Standalone()

	switchCmd.Flags().BoolP("create", "c", false, "Create the branch first, then switch to it")
	switchCmd.Flags().Bool("detach", false, "Check out a commit in detached HEAD mode")
	switchCmd.Flags().Bool("dry-run", false, "Run overwrite checks without updating refs, the index, or the worktree")
	switchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(switchCmd)

	carapace.Gen(switchCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionRemoteBranchNames(""),
			git.ActionRefs(git.RefOption{LocalBranches: true}),
		).ToA(),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
