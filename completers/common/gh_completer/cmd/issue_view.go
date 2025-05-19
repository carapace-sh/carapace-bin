package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_viewCmd = &cobra.Command{
	Use:     "view {<number> | <url>}",
	Short:   "View an issue",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_viewCmd).Standalone()

	issue_viewCmd.Flags().BoolP("comments", "c", false, "View issue comments")
	issue_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	issue_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	issue_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	issue_viewCmd.Flags().BoolP("web", "w", false, "Open an issue in the browser")
	issueCmd.AddCommand(issue_viewCmd)

	carapace.Gen(issue_viewCmd).FlagCompletion(carapace.ActionMap{
		"json": action.ActionIssueFields().UniqueList(","),
	})

	carapace.Gen(issue_viewCmd).PositionalCompletion(
		action.ActionIssues(issue_viewCmd, action.IssueOpts{Open: true}),
	)
}
