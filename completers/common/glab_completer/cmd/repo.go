package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:     "repo <command> [flags]",
	Short:   "Work with GitLab repositories and projects.",
	Aliases: []string{"project"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repoCmd).Standalone()

	rootCmd.AddCommand(repoCmd)
}
