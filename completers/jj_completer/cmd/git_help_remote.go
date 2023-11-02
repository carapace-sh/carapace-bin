package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage Git remotes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_remoteCmd).Standalone()

	git_helpCmd.AddCommand(git_help_remoteCmd)
}
