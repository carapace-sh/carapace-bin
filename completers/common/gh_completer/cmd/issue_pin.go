package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_pinCmd = &cobra.Command{
	Use:     "pin {<number> | <url>}",
	Short:   "Pin a issue",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_pinCmd).Standalone()

	issueCmd.AddCommand(issue_pinCmd)

	carapace.Gen(issue_pinCmd).PositionalCompletion(
		action.ActionIssues(issue_pinCmd, action.IssueOpts{Open: true, Closed: true}),
	)
}
