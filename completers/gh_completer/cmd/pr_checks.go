package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_checksCmd = &cobra.Command{
	Use:     "checks [<number> | <url> | <branch>]",
	Short:   "Show CI status for a single pull request",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_checksCmd).Standalone()

	pr_checksCmd.Flags().Bool("fail-fast", false, "Exit watch mode on first check failure")
	pr_checksCmd.Flags().StringP("interval", "i", "", "Refresh interval in seconds when using `--watch` flag")
	pr_checksCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	pr_checksCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	pr_checksCmd.Flags().Bool("required", false, "Only show checks that are required")
	pr_checksCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	pr_checksCmd.Flags().Bool("watch", false, "Watch checks until they finish")
	pr_checksCmd.Flags().BoolP("web", "w", false, "Open the web browser to show details about checks")
	prCmd.AddCommand(pr_checksCmd)

	carapace.Gen(pr_checksCmd).FlagCompletion(carapace.ActionMap{
		"json": action.ActionPullRequestCheckFields().UniqueList(","),
	})

	carapace.Gen(pr_checksCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionPullRequests(pr_checksCmd, action.PullRequestOpts{Open: true}),
				action.ActionBranches(pr_checksCmd),
			).Invoke(c).Merge().ToA()
		}),
	)
}
