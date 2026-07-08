package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var remote_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_helpCmd).Standalone()

	remoteCmd.AddCommand(remote_helpCmd)

	carapace.Gen(remote_helpCmd).PositionalCompletion(
		carapace.ActionCommands(remoteCmd),
	)

	carapace.Gen(remote_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(remoteCmd),
	)
}
