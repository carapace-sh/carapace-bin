package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_git_remote_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_git_remote_renameCmd).Standalone()

	help_git_remoteCmd.AddCommand(help_git_remote_renameCmd)
}
