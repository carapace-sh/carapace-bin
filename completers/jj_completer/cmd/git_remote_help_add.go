package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_remote_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_help_addCmd).Standalone()

	git_remote_helpCmd.AddCommand(git_remote_help_addCmd)
}
