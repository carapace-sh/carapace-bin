package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_setDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Make the named branch the default so all worktree or index changes are associated with it automatically",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_setDefaultCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_setDefaultCmd)
}
