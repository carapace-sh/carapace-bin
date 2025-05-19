package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var issues_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an issue on repository",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issues_createCmd).Standalone()

	issues_createCmd.Flags().StringP("assignees", "a", "", "Comma-separated list of usernames to assign")
	issues_createCmd.Flags().StringP("deadline", "D", "", "Deadline timestamp to assign")
	issues_createCmd.Flags().StringP("description", "d", "", "")
	issues_createCmd.Flags().StringP("labels", "L", "", "Comma-separated list of labels to assign")
	issues_createCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	issues_createCmd.Flags().StringP("milestone", "m", "", "Milestone to assign")
	issues_createCmd.Flags().StringP("referenced-version", "v", "", "commit-hash or tag name to assign")
	issues_createCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	issues_createCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	issues_createCmd.Flags().StringP("title", "t", "", "")
	issuesCmd.AddCommand(issues_createCmd)

	// TODO completion
	carapace.Gen(issues_createCmd).FlagCompletion(carapace.ActionMap{
		"labels": action.ActionLabels(issues_createCmd).UniqueList(","),
		"login":  tea.ActionLogins(),
		"remote": git.ActionRemotes(),
	})
}
