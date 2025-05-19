package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_gitignoreCmd = &cobra.Command{
	Use:     "gitignore <command>",
	Short:   "List and view available repository gitignore templates",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_gitignoreCmd).Standalone()

	repoCmd.AddCommand(repo_gitignoreCmd)
}
