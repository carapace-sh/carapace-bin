package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_git_remote_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a Git remote and forget its branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_git_remote_removeCmd).Standalone()

	help_git_remoteCmd.AddCommand(help_git_remote_removeCmd)
}
