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

	git_initCmd.Flags().Bool("colocate", false, "Specifies that the `jj` repo should also be a valid `git` repo")
	git_initCmd.Flags().String("git-repo", "", "Specifies a path to an existing git repository to be used to back the new `jj` repo")
	git_initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gitCmd.AddCommand(git_initCmd)

	carapace.Gen(git_initCmd).FlagCompletion(carapace.ActionMap{
		"git-repo": carapace.ActionDirectories(),
	})

	carapace.Gen(git_initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
