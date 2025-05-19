package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_commentCmd = &cobra.Command{
	Use:     "comment [<number> | <url> | <branch>]",
	Short:   "Add a comment to a pull request",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_commentCmd).Standalone()

	pr_commentCmd.Flags().StringP("body", "b", "", "The comment body `text`")
	pr_commentCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	pr_commentCmd.Flags().Bool("create-if-none", false, "Create a new comment if no comments are found. Can be used only with --edit-last")
	pr_commentCmd.Flags().Bool("edit-last", false, "Edit the last comment of the same author")
	pr_commentCmd.Flags().BoolP("editor", "e", false, "Skip prompts and open the text editor to write the body in")
	pr_commentCmd.Flags().BoolP("web", "w", false, "Open the web browser to write the comment")
	prCmd.AddCommand(pr_commentCmd)

	carapace.Gen(pr_commentCmd).FlagCompletion(carapace.ActionMap{
		"body":      action.ActionBody(pr_commentCmd),
		"body-file": carapace.ActionFiles(),
	})

	carapace.Gen(pr_commentCmd).PositionalCompletion(
		action.ActionPullRequests(pr_commentCmd, action.PullRequestOpts{Open: true}),
	)
}
