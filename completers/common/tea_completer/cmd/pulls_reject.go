package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var pulls_rejectCmd = &cobra.Command{
	Use:   "reject",
	Short: "Request changes to a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_rejectCmd).Standalone()

	pulls_rejectCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_rejectCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pulls_rejectCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_rejectCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pullsCmd.AddCommand(pulls_rejectCmd)

	// TODO completion
	carapace.Gen(pulls_rejectCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(pulls_rejectCmd).PositionalCompletion(
		action.ActionPullrequests(pulls_rejectCmd, true),
	)
}
