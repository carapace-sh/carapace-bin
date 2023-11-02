package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_remote_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_remote_addCmd).Standalone()

	git_help_remoteCmd.AddCommand(git_help_remote_addCmd)
}
