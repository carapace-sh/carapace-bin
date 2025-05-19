package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_unlockCmd = &cobra.Command{
	Use:     "unlock {<number> | <url>}",
	Short:   "Unlock issue conversation",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_unlockCmd).Standalone()

	issueCmd.AddCommand(issue_unlockCmd)

	carapace.Gen(issue_unlockCmd).PositionalCompletion(
		action.ActionIssues(issue_lockCmd, action.IssueOpts{Open: true, Closed: true}), // TODO just complete locked issues
	)
}
