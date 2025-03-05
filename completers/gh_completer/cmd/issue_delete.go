package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_deleteCmd = &cobra.Command{
	Use:     "delete {<number> | <url>}",
	Short:   "Delete issue",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_deleteCmd).Standalone()

	issue_deleteCmd.Flags().Bool("confirm", false, "Confirm deletion without prompting")
	issue_deleteCmd.Flags().Bool("yes", false, "Confirm deletion without prompting")
	issue_deleteCmd.Flag("confirm").Hidden = true
	issueCmd.AddCommand(issue_deleteCmd)

	carapace.Gen(issue_deleteCmd).PositionalCompletion(
		action.ActionIssues(issue_deleteCmd, action.IssueOpts{Open: true, Closed: true}),
	)
}
