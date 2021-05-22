package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_cloneCmd = &cobra.Command{
	Use:   "clone <repository> [<directory>] [-- <gitflags>...]",
	Short: "Clone a repository locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	repoCmd.AddCommand(repo_cloneCmd)

	carapace.Gen(repo_cloneCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_cloneCmd),
		carapace.ActionDirectories(),
	)
}
