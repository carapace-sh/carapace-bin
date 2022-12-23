package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var issue_developCmd = &cobra.Command{
	Use:     "develop [flags] {<number> | <url>}",
	Short:   "Manage linked branches for an issue",
	GroupID: "targeted",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_developCmd).Standalone()

	issue_developCmd.Flags().StringP("base", "b", "", "Name of the base branch you want to make your new branch from")
	issue_developCmd.Flags().BoolP("checkout", "c", false, "Checkout the branch after creating it")
	issue_developCmd.Flags().StringP("issue-repo", "i", "", "Name or URL of the issue's repository")
	issue_developCmd.Flags().BoolP("list", "l", false, "List linked branches for the issue")
	issue_developCmd.Flags().StringP("name", "n", "", "Name of the branch to create")
	issueCmd.AddCommand(issue_developCmd)

	carapace.Gen(issue_developCmd).FlagCompletion(carapace.ActionMap{
		"base":       action.ActionBranches(issue_developCmd),
		"issue-repo": gh.ActionOwnerRepositories(gh.HostOpts{}),
		"name":       action.ActionBranches(issue_developCmd),
	})

	carapace.Gen(issue_developCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := issue_developCmd.Flag("issue-repo"); flag.Changed {
				cmd := &cobra.Command{}
				cmd.Flags().String("repo", flag.Value.String(), "fake-repo")
				cmd.Flag("repo").Changed = true
				return action.ActionIssues(cmd, action.IssueOpts{Open: true})
			}
			return action.ActionIssues(issue_developCmd, action.IssueOpts{Open: true})
		}),
	)
}
