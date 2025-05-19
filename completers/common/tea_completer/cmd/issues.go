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

var issuesCmd = &cobra.Command{
	Use:     "issues",
	Short:   "List, create and update issues",
	GroupID: "ENTITIES",
	Aliases: []string{"issue", "i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issuesCmd).Standalone()

	issuesCmd.Flags().StringP("assignee", "a", "", "")
	issuesCmd.Flags().StringP("author", "A", "", "")
	issuesCmd.Flags().Bool("comments", false, "Whether to display comments (will prompt if not provided & run interactively)")
	issuesCmd.Flags().StringP("fields", "f", "", "Comma-separated list of fields to print. Available values:")
	issuesCmd.Flags().StringP("from", "F", "", "Filter by activity after this date")
	issuesCmd.Flags().StringP("keyword", "k", "", "Filter by search string")
	issuesCmd.Flags().StringP("kind", "K", "", "Whether to return `issues`, `pulls`, or `all` (you can use this to apply advanced search filters to PRs)")
	issuesCmd.Flags().StringP("labels", "L", "", "Comma-separated list of labels to match issues against.")
	issuesCmd.Flags().String("limit", "", "specify limit of items per page")
	issuesCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	issuesCmd.Flags().StringP("mentions", "M", "", "")
	issuesCmd.Flags().StringP("milestones", "m", "", "Comma-separated list of milestones to match issues against.")
	issuesCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	issuesCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	issuesCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	issuesCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	issuesCmd.Flags().String("state", "", "Filter by state (all|open|closed)")
	issuesCmd.Flags().StringP("until", "u", "", "Filter by activity before this date")
	rootCmd.AddCommand(issuesCmd)

	// TODO completion
	carapace.Gen(issuesCmd).FlagCompletion(carapace.ActionMap{
		"fields": tea.ActionIssueFields().UniqueList(","),
		"from":   time.ActionDate(),
		"kind":   carapace.ActionValues("issues", "pulls", "all"),
		"labels": action.ActionLabels(issuesCmd).UniqueList(","),
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"state":  carapace.ActionValues("all", "open", "closed").StyleF(style.ForKeyword),
		"until":  time.ActionDate(),
	})
}
