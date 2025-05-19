package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var pulls_reopenCmd = &cobra.Command{
	Use:     "reopen",
	Short:   "Change state of one or more pull requests to 'open'",
	Aliases: []string{"open"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_reopenCmd).Standalone()

	pulls_reopenCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_reopenCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pulls_reopenCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_reopenCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pullsCmd.AddCommand(pulls_reopenCmd)

	// TODO completion
	carapace.Gen(pulls_reopenCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(pulls_reopenCmd).PositionalCompletion(
		action.ActionPullrequests(pulls_reopenCmd, false),
	)
}
