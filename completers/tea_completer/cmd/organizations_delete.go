package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var organizations_deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete users Organizations",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_deleteCmd).Standalone()

	organizations_deleteCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	organizations_deleteCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	organizationsCmd.AddCommand(organizations_deleteCmd)

	// TODO completion
	carapace.Gen(organizations_deleteCmd).FlagCompletion(carapace.ActionMap{
		"remote": git.ActionRemotes(),
	})
}
