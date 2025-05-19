package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var pulls_mergeCmd = &cobra.Command{
	Use:     "merge",
	Short:   "Merge a pull request",
	Aliases: []string{"m"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_mergeCmd).Standalone()

	pulls_mergeCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_mergeCmd.Flags().StringP("message", "m", "", "Merge commit message")
	pulls_mergeCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pulls_mergeCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_mergeCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pulls_mergeCmd.Flags().StringP("style", "s", "", "Kind of merge to perform: merge, rebase, squash, rebase-merge")
	pulls_mergeCmd.Flags().StringP("title", "t", "", "Merge commit title")
	pullsCmd.AddCommand(pulls_mergeCmd)

	// TODO completion
	carapace.Gen(pulls_mergeCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"style":  carapace.ActionValues("merge", "rebase", "squash", "rebase-merge"),
	})

	carapace.Gen(pulls_mergeCmd).PositionalCompletion(
		action.ActionPullrequests(pulls_mergeCmd, true),
	)
}
