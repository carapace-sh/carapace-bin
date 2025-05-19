package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var pulls_checkoutCmd = &cobra.Command{
	Use:     "checkout",
	Short:   "Locally check out the given PR",
	Aliases: []string{"co"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_checkoutCmd).Standalone()

	pulls_checkoutCmd.Flags().BoolP("branch", "b", false, "Create a local branch if it doesn't exist yet")
	pulls_checkoutCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_checkoutCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pulls_checkoutCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_checkoutCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pullsCmd.AddCommand(pulls_checkoutCmd)

	// TODO completion
	carapace.Gen(pulls_checkoutCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})

	carapace.Gen(pulls_checkoutCmd).PositionalCompletion(
		action.ActionPullrequests(pulls_checkoutCmd, true),
	)
}
