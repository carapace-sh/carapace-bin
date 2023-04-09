package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/tea_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var pulls_approveCmd = &cobra.Command{
	Use:     "approve",
	Short:   "Approve a pull request",
	Aliases: []string{"lgtm", "a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_approveCmd).Standalone()

	pulls_approveCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_approveCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pulls_approveCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_approveCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pullsCmd.AddCommand(pulls_approveCmd)

	// TODO completion
	carapace.Gen(pulls_approveCmd).FlagCompletion(carapace.ActionMap{
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(pulls_approveCmd).PositionalCompletion(
		action.ActionPullrequests(pulls_approveCmd, true),
	)
}
