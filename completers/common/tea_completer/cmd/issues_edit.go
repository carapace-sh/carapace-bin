package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var issues_editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit one or more issues",
	Aliases: []string{"e"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issues_editCmd).Standalone()

	issues_editCmd.Flags().StringP("add-assignees", "a", "", "Comma-separated list of usernames to assign")
	issues_editCmd.Flags().StringP("add-labels", "L", "", "Comma-separated list of labels to assign. Takes precedence over --remove-labels")
	issues_editCmd.Flags().StringP("deadline", "D", "", "Deadline timestamp to assign")
	issues_editCmd.Flags().StringP("description", "d", "", "")
	issues_editCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	issues_editCmd.Flags().StringP("milestone", "m", "", "Milestone to assign")
	issues_editCmd.Flags().StringP("referenced-version", "v", "", "commit-hash or tag name to assign")
	issues_editCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	issues_editCmd.Flags().String("remove-labels", "", "Comma-separated list of labels to remove")
	issues_editCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	issues_editCmd.Flags().StringP("title", "t", "", "")
	issuesCmd.AddCommand(issues_editCmd)

	// TODO completion
	carapace.Gen(issues_editCmd).FlagCompletion(carapace.ActionMap{
		"add-labels":    action.ActionLabels(issues_editCmd).UniqueList(","), // TODO assigned labels
		"login":         tea.ActionLogins(),
		"remote":        git.ActionRemotes(),
		"remove-labels": action.ActionLabels(issues_editCmd).UniqueList(","), // TODO assigned labels
	})

	carapace.Gen(issues_editCmd).PositionalCompletion(
		action.ActionIssues(issues_editCmd, true),
	)
}
