package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_git_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Create a new repo backed by a clone of a Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_git_cloneCmd).Standalone()

	help_gitCmd.AddCommand(help_git_cloneCmd)
}
