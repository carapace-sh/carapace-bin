package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_help_remote_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_help_remote_renameCmd).Standalone()

	git_help_remoteCmd.AddCommand(git_help_remote_renameCmd)
}
