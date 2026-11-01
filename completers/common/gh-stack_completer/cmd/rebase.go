package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:     "rebase [branch]",
	Short:   "Rebase a stack of branches",
	GroupID: "remote",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebaseCmd).Standalone()

	rebaseCmd.Flags().Bool("abort", false, "Abort rebase and restore all branches")
	rebaseCmd.Flags().Bool("committer-date-is-author-date", false, "Set the committer date to the author date during rebase")
	rebaseCmd.Flags().Bool("continue", false, "Continue rebase after resolving conflicts")
	rebaseCmd.Flags().Bool("downstack", false, "Only rebase branches from trunk to current branch")
	rebaseCmd.Flags().Bool("no-trunk", false, "Skip trunk. Only rebase stack branches onto each other")
	rebaseCmd.Flags().Bool("preserve-dates", false, "Alias for --committer-date-is-author-date")
	rebaseCmd.Flags().String("remote", "", "Remote to fetch from (defaults to auto-detected remote)")
	rebaseCmd.Flags().Bool("upstack", false, "Only rebase branches from current branch to top")
	rootCmd.AddCommand(rebaseCmd)

	carapace.Gen(rebaseCmd).FlagCompletion(carapace.ActionMap{
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(rebaseCmd).PositionalCompletion(
		git.ActionLocalBranches(),
	)
}
