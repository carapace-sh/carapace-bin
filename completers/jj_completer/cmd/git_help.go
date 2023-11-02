package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_helpCmd).Standalone()

	gitCmd.AddCommand(git_helpCmd)

	carapace.Gen(git_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(gitCmd),
	)
}
