package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List and filter pull requests in this repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	pr_listCmd.Flags().BoolP("web", "w", false, "Open the browser to list the pull requests")
	pr_listCmd.Flags().IntP("limit", "L", 30, "Maximum number of items to fetch")
	pr_listCmd.Flags().StringP("assignee", "a", "", "Filter by assignee")
	pr_listCmd.Flags().StringP("author", "A", "", "Filter by author")
	pr_listCmd.Flags().StringP("base", "B", "", "Filter by base branch")
	pr_listCmd.Flags().BoolP("draft", "d", false, "Filter by draft state")
	pr_listCmd.Flags().StringP("head", "H", "", "Filter by head branch")
	pr_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	pr_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	pr_listCmd.Flags().StringSliceP("label", "l", nil, "Filter by labels")
	pr_listCmd.Flags().StringP("search", "S", "", "Search pull requests with `query`")
	pr_listCmd.Flags().StringP("state", "s", "open", "Filter by state: {open|closed|merged|all}")
	pr_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	prCmd.AddCommand(pr_listCmd)

	carapace.Gen(pr_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee": action.ActionAssignableUsers(pr_listCmd),
		"author":   action.ActionUsers(pr_listCmd, action.UserOpts{Users: true}),
		"base":     action.ActionBranches(pr_listCmd),
		"head":     action.ActionBranches(pr_listCmd),
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPullRequestFields().Invoke(c).Filter(c.Parts).ToA()
		}),
		"label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(pr_listCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"state": carapace.ActionValues("open", "closed", "merged", "all"),
	})
}
