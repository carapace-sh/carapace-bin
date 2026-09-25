package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_git_pre_commitCmd = &cobra.Command{
	Use:   "git-pre-commit",
	Short: "Generate a git pre-commit hook",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_git_pre_commitCmd).Standalone()

	generateCmd.AddCommand(generate_git_pre_commitCmd)
}
