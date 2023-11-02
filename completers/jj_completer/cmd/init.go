package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [OPTIONS] [DESTINATION]",
	Short: "Create a new repo in the given directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("git", false, "Use the Git backend, creating a jj repo backed by a Git repo")
	initCmd.Flags().String("git-repo", "", "Path to a git repo the jj repo will be backed by")
	initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"git-repo": carapace.ActionDirectories(),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
