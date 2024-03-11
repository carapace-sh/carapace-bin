package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var times_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset tracked time on an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(times_resetCmd).Standalone()

	times_resetCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	times_resetCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	times_resetCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	timesCmd.AddCommand(times_resetCmd)

	// TODO completion
	carapace.Gen(times_resetCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"remote": git.ActionRemotes(),
	})
}
