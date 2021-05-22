package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_commentCmd = &cobra.Command{
	Use:   "comment [<number> | <url> | <branch>]",
	Short: "Create a new pr comment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	pr_commentCmd.Flags().StringP("body", "b", "", "Supply a body. Will prompt for one otherwise.")
	pr_commentCmd.Flags().StringP("body-file", "F", "", "Read body text from `file`")
	pr_commentCmd.Flags().BoolP("editor", "e", false, "Add body using editor")
	pr_commentCmd.Flags().BoolP("web", "w", false, "Add body in browser")
	prCmd.AddCommand(pr_commentCmd)

	carapace.Gen(pr_commentCmd).FlagCompletion(carapace.ActionMap{
		"body-file": carapace.ActionFiles(),
	})

	carapace.Gen(pr_commentCmd).PositionalCompletion(
		action.ActionPullRequests(pr_commentCmd, action.PullRequestOpts{Open: true}),
	)
}
