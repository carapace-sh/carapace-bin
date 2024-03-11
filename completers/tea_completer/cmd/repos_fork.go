package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var repos_forkCmd = &cobra.Command{
	Use:     "fork",
	Short:   "Fork an existing repository",
	Aliases: []string{"f"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repos_forkCmd).Standalone()

	repos_forkCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	repos_forkCmd.Flags().StringP("owner", "O", "", "name of fork's owner, defaults to current user")
	repos_forkCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	repos_forkCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	reposCmd.AddCommand(repos_forkCmd)

	// TODO completion
	carapace.Gen(repos_forkCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"remote": git.ActionRemotes(),
	})
}
