package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_remote_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a Git remote and forget its branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_help_removeCmd).Standalone()

	git_remote_helpCmd.AddCommand(git_remote_help_removeCmd)
}
