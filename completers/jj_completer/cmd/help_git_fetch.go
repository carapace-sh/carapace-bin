package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_git_fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch from a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_git_fetchCmd).Standalone()

	help_gitCmd.AddCommand(help_git_fetchCmd)
}
