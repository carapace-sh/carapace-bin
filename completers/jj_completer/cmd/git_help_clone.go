package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Create a new repo backed by a clone of a Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_cloneCmd).Standalone()

	git_helpCmd.AddCommand(git_help_cloneCmd)
}
