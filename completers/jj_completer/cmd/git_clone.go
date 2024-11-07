package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var git_cloneCmd = &cobra.Command{
	Use:   "clone [OPTIONS] <SOURCE> [DESTINATION]",
	Short: "Create a new repo backed by a clone of a Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_cloneCmd).Standalone()

	git_cloneCmd.Flags().Bool("colocate", false, "Whether or not to colocate the Jujutsu repo with the git repo")
	git_cloneCmd.Flags().Int("depth", 0, "Create a shallow clone of the given depth")
	git_cloneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_cloneCmd.Flags().String("remote", "origin", "Name of the newly created remote")
	gitCmd.AddCommand(git_cloneCmd)

	carapace.Gen(git_cloneCmd).PositionalCompletion(
		git.ActionRepositorySearch(git.SearchOpts{}.Default()),
		carapace.ActionDirectories(),
	)
}
