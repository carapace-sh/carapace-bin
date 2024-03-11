package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var times_addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Track spent time on an issue",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(times_addCmd).Standalone()

	times_addCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	times_addCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	times_addCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	timesCmd.AddCommand(times_addCmd)

	// TODO completion
	carapace.Gen(times_addCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"remote": git.ActionRemotes(),
	})
}
