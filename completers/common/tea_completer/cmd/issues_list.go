package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var issues_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List issues of the repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issues_listCmd).Standalone()

	issues_listCmd.Flags().StringP("assignee", "a", "", "")
	issues_listCmd.Flags().StringP("author", "A", "", "")
	issues_listCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	issues_listCmd.Flags().StringP("from", "F", "", "Filter by activity after this date")
	issues_listCmd.Flags().StringP("keyword", "k", "", "Filter by search string")
	issues_listCmd.Flags().StringP("kind", "K", "", "Whether to return `issues`, `pulls`, or `all` (you can use this to apply advanced search filters to PRs)")
	issues_listCmd.Flags().StringP("labels", "L", "", "Comma-separated list of labels to match issues against.")
	issues_listCmd.Flags().String("limit", "", "specify limit of items per page")
	issues_listCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	issues_listCmd.Flags().StringP("mentions", "M", "", "")
	issues_listCmd.Flags().StringP("milestones", "m", "", "Comma-separated list of milestones to match issues against.")
	issues_listCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	issues_listCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	issues_listCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	issues_listCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	issues_listCmd.Flags().String("state", "", "Filter by state (all|open|closed)")
	issues_listCmd.Flags().StringP("until", "u", "", "Filter by activity before this date")
	issuesCmd.AddCommand(issues_listCmd)

	// TODO completion
	carapace.Gen(issues_listCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionIssueFields().UniqueList(","),
		"from":   time.ActionDate(),
		"kind":   carapace.ActionValues("issues", "pulls", "all"),
		"labels": action.ActionLabels(issues_listCmd).UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"state":  carapace.ActionValues("all", "open", "closed").StyleF(style.ForKeyword),
		"until":  time.ActionDate(),
	})
}
