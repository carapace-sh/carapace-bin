package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_commentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Create a new issue comment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_commentCmd).Standalone()
	issue_commentCmd.Flags().StringP("body", "b", "", "Supply a body. Will prompt for one otherwise.")
	issue_commentCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	issue_commentCmd.Flags().BoolP("editor", "e", false, "Add body using editor")
	issue_commentCmd.Flags().BoolP("web", "w", false, "Add body in browser")
	issueCmd.AddCommand(issue_commentCmd)

	carapace.Gen(issue_commentCmd).FlagCompletion(carapace.ActionMap{
		"body-file": carapace.ActionFiles(),
	})
	carapace.Gen(issue_commentCmd).PositionalCompletion(
		action.ActionIssues(issue_commentCmd, action.IssueOpts{Open: true}),
	)
}
