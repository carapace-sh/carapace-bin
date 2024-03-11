package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_remote_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_remote_helpCmd).Standalone()

	git_remoteCmd.AddCommand(git_remote_helpCmd)

	carapace.Gen(git_remote_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(git_remoteCmd),
	)
}
