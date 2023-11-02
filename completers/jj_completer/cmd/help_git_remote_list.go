package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_git_remote_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Git remotes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_git_remote_listCmd).Standalone()

	help_git_remoteCmd.AddCommand(help_git_remote_listCmd)
}
