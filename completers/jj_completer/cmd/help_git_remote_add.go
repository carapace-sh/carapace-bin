package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_git_remote_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_git_remote_addCmd).Standalone()

	help_git_remoteCmd.AddCommand(help_git_remote_addCmd)
}
