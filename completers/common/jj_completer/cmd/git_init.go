package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_initCmd = &cobra.Command{
	Use:   "init [OPTIONS] [DESTINATION]",
	Short: "Create a new Git backed repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_initCmd).Standalone()

	git_initCmd.Flags().Bool("colocate", false, "Colocate the Jujutsu repo with the git repo")
	git_initCmd.Flags().String("git-repo", "", "Specifies a path to an **existing** git repository to be used as the backing git repo for the newly created `jj` repo")
	git_initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_initCmd.Flags().Bool("no-colocate", false, "Disable colocation of the Jujutsu repo with the git repo")
	gitCmd.AddCommand(git_initCmd)

	carapace.Gen(git_initCmd).FlagCompletion(carapace.ActionMap{
		"git-repo": carapace.ActionDirectories(),
	})

	carapace.Gen(git_initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
