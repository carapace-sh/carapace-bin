package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_remote_help_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_help_renameCmd).Standalone()

	git_remote_helpCmd.AddCommand(git_remote_help_renameCmd)
}
