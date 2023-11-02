package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_remote_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Git remotes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_help_listCmd).Standalone()

	git_remote_helpCmd.AddCommand(git_remote_help_listCmd)
}
