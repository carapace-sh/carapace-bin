package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/spf13/cobra"
)

var issue_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List issues in a repository",
	GroupID: "General commands",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_listCmd).Standalone()

	issue_listCmd.Flags().String("app", "", "Filter by GitHub App author")
	issue_listCmd.Flags().StringP("assignee", "a", "", "Filter by assignee")
	issue_listCmd.Flags().StringP("author", "A", "", "Filter by author")
	issue_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	issue_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	issue_listCmd.Flags().StringSliceP("label", "l", nil, "Filter by label")
	issue_listCmd.Flags().StringP("limit", "L", "", "Maximum number of issues to fetch")
	issue_listCmd.Flags().String("mention", "", "Filter by mention")
	issue_listCmd.Flags().StringP("milestone", "m", "", "Filter by milestone number or title")
	issue_listCmd.Flags().StringP("search", "S", "", "Search issues with `query`")
	issue_listCmd.Flags().StringP("state", "s", "", "Filter by state: {open|closed|all}")
	issue_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	issue_listCmd.Flags().BoolP("web", "w", false, "List issues in the web browser")
	issueCmd.AddCommand(issue_listCmd)

	// TODO app completion
	carapace.Gen(issue_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee":  action.ActionAssignableUsers(issue_listCmd),
		"author":    gh.ActionUsers(gh.HostOpts{}),
		"json":      action.ActionIssueFields().UniqueList(","),
		"label":     action.ActionLabels(issue_listCmd).UniqueList(","),
		"mention":   action.ActionMentionableUsers(issue_listCmd),
		"milestone": action.ActionMilestones(issue_listCmd),
		"state":     carapace.ActionValues("open", "closed", "all").StyleF(styles.Gh.ForState),
	})
}
