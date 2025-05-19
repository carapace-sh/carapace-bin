package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_viewCmd = &cobra.Command{
	Use:     "view [<number> | <url> | <branch>]",
	Short:   "View a pull request",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_viewCmd).Standalone()

	pr_viewCmd.Flags().BoolP("comments", "c", false, "View pull request comments")
	pr_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	pr_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	pr_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	pr_viewCmd.Flags().BoolP("web", "w", false, "Open a pull request in the browser")
	prCmd.AddCommand(pr_viewCmd)

	carapace.Gen(pr_viewCmd).FlagCompletion(carapace.ActionMap{
		"json": action.ActionPullRequestFields().UniqueList(","),
	})

	carapace.Gen(pr_viewCmd).PositionalCompletion(
		action.ActionPullRequests(pr_viewCmd, action.PullRequestOpts{Open: true}),
	)
}
