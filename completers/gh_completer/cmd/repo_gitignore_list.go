package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_gitignore_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List available repository gitignore templates",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_gitignore_listCmd).Standalone()

	repo_gitignoreCmd.AddCommand(repo_gitignore_listCmd)
}
