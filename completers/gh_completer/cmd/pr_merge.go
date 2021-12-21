package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_mergeCmd).Standalone()
	pr_mergeCmd.Flags().Bool("admin", false, "Use administrator privileges to merge a pull request that does not meet requirements")
	pr_mergeCmd.Flags().Bool("auto", false, "Automatically merge only after necessary requirements are met")
	pr_mergeCmd.Flags().StringP("body", "b", "", "Body `text` for the merge commit")
	pr_mergeCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	pr_mergeCmd.Flags().BoolP("delete-branch", "d", false, "Delete the local and remote branch after merge")
	pr_mergeCmd.Flags().Bool("disable-auto", false, "Disable auto-merge for this pull request")
	pr_mergeCmd.Flags().BoolP("merge", "m", false, "Merge the commits with the base branch")
	pr_mergeCmd.Flags().BoolP("rebase", "r", false, "Rebase the commits onto the base branch")
	pr_mergeCmd.Flags().BoolP("squash", "s", false, "Squash the commits into one commit and merge it into the base branch")
	pr_mergeCmd.Flags().StringP("subject", "t", "", "Subject `text` for the merge commit")
	prCmd.AddCommand(pr_mergeCmd)

	carapace.Gen(pr_mergeCmd).FlagCompletion(carapace.ActionMap{
		"body-file": carapace.ActionFiles(),
	})

	carapace.Gen(pr_mergeCmd).PositionalCompletion(
		action.ActionPullRequests(pr_mergeCmd, action.PullRequestOpts{Open: true}),
	)
}
