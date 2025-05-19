package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var pulls_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Deletes local & remote feature-branches for a closed pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_cleanCmd).Standalone()

	pulls_cleanCmd.Flags().Bool("ignore-sha", false, "Find the local branch by name instead of commit hash (less precise)")
	pulls_cleanCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_cleanCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pulls_cleanCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_cleanCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pullsCmd.AddCommand(pulls_cleanCmd)

	// TODO completion
	carapace.Gen(pulls_cleanCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(pulls_cleanCmd).PositionalCompletion(
		action.ActionPullrequests(pulls_cleanCmd, false),
	)
}
