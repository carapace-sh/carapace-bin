package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_remote_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Git remotes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_remote_listCmd).Standalone()

	git_help_remoteCmd.AddCommand(git_help_remote_listCmd)
}
