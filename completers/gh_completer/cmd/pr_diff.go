package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_diffCmd = &cobra.Command{
	Use:   "diff [<number> | <url> | <branch>]",
	Short: "View changes in a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	pr_diffCmd.Flags().String("color", "auto", "Use color in diff output: {always|never|auto}")
	pr_diffCmd.Flags().Bool("patch", false, "Display diff in patch format")
	prCmd.AddCommand(pr_diffCmd)

	carapace.Gen(pr_diffCmd).PositionalCompletion(
		action.ActionPullRequests(pr_diffCmd, action.PullRequestOpts{Open: true}),
	)

	carapace.Gen(pr_diffCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("always", "never", "auto"),
	})
}
