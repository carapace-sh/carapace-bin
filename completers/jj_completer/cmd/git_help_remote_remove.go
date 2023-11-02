package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_remote_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a Git remote and forget its branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_remote_removeCmd).Standalone()

	git_help_remoteCmd.AddCommand(git_help_remote_removeCmd)
}
