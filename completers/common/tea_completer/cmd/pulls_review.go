package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var pulls_reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Interactively review a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_reviewCmd).Standalone()

	pulls_reviewCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_reviewCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pulls_reviewCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_reviewCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pullsCmd.AddCommand(pulls_reviewCmd)

	// TODO completion
	carapace.Gen(pulls_reviewCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(pulls_reviewCmd).PositionalCompletion(
		action.ActionPullrequests(pulls_reviewCmd, true),
	)
}
