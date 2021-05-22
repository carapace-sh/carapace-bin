package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_checksCmd = &cobra.Command{
	Use:   "checks [<number> | <url> | <branch>]",
	Short: "Show CI status for a single pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	pr_checksCmd.Flags().BoolP("web", "w", false, "Open the web browser to show details about checks")
	prCmd.AddCommand(pr_checksCmd)

	carapace.Gen(pr_checksCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			pullRequests := action.ActionPullRequests(pr_checksCmd, action.PullRequestOpts{Open: true}).Invoke(c)
			branches := action.ActionBranches(pr_checksCmd).Invoke(c)
			return pullRequests.Merge(branches).ToA()
		}),
	)
}
