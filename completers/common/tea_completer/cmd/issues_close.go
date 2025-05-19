package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var issues_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Change state of one ore more issues to 'closed'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issues_closeCmd).Standalone()

	issues_closeCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	issues_closeCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	issues_closeCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	issues_closeCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	issuesCmd.AddCommand(issues_closeCmd)

	// TODO completion
	carapace.Gen(issues_closeCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(issues_closeCmd).PositionalCompletion(
		action.ActionIssues(issues_closeCmd, true),
	)
}
