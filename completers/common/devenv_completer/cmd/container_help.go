package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_helpCmd).Standalone()

	containerCmd.AddCommand(container_helpCmd)

	carapace.Gen(container_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(containerCmd),
	)
}
