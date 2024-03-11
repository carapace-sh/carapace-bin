package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparse_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparse_helpCmd).Standalone()

	sparseCmd.AddCommand(sparse_helpCmd)

	carapace.Gen(sparse_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(sparseCmd),
	)
}
