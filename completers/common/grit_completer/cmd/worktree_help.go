package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var worktree_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(worktree_helpCmd).Standalone()

	worktreeCmd.AddCommand(worktree_helpCmd)

	carapace.Gen(worktree_helpCmd).PositionalCompletion(
		carapace.ActionCommands(worktreeCmd),
	)

	carapace.Gen(worktree_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(worktreeCmd),
	)
}
