package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var times_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a single tracked time on an issue",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(times_deleteCmd).Standalone()

	times_deleteCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	times_deleteCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	times_deleteCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	timesCmd.AddCommand(times_deleteCmd)

	// TODO completion
	carapace.Gen(times_deleteCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"remote": git.ActionRemotes(),
	})
}
