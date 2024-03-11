package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:     "open",
	Short:   "Open something of the repository in web browser",
	GroupID: "HELPERS",
	Aliases: []string{"o"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openCmd).Standalone()

	openCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	openCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	openCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	rootCmd.AddCommand(openCmd)

	// TODO completion
	carapace.Gen(openCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"remote": git.ActionRemotes(),
	})
}
