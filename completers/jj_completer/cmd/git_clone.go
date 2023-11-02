package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Create a new repo backed by a clone of a Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_cloneCmd).Standalone()

	git_cloneCmd.Flags().Bool("colocate", false, "Whether or not to colocate the Jujutsu repo with the git repo")
	git_cloneCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gitCmd.AddCommand(git_cloneCmd)
}
