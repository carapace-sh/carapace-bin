package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_commentCmd = &cobra.Command{
	Use:     "comment {<number> | <url>}",
	Short:   "Add a comment to an issue",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_commentCmd).Standalone()

	issue_commentCmd.Flags().StringP("body", "b", "", "The comment body `text`")
	issue_commentCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	issue_commentCmd.Flags().Bool("create-if-none", false, "Create a new comment if no comments are found. Can be used only with --edit-last")
	issue_commentCmd.Flags().Bool("delete-last", false, "Delete the last comment of the current user")
	issue_commentCmd.Flags().Bool("edit-last", false, "Edit the last comment of the current user")
	issue_commentCmd.Flags().BoolP("editor", "e", false, "Skip prompts and open the text editor to write the body in")
	issue_commentCmd.Flags().BoolP("web", "w", false, "Open the web browser to write the comment")
	issue_commentCmd.Flags().Bool("yes", false, "Skip the delete confirmation prompt when --delete-last is provided")
	issueCmd.AddCommand(issue_commentCmd)

	carapace.Gen(issue_commentCmd).FlagCompletion(carapace.ActionMap{
		"body":      action.ActionBody(issue_commentCmd),
		"body-file": carapace.ActionFiles(),
	})

	carapace.Gen(issue_commentCmd).PositionalCompletion(
		action.ActionIssues(issue_commentCmd, action.IssueOpts{Open: true, Closed: true}),
	)
}
