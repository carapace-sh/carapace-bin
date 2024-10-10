package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bookmark_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_helpCmd).Standalone()

	bookmarkCmd.AddCommand(bookmark_helpCmd)

	carapace.Gen(bookmark_helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(bookmarkCmd),
	)
}
