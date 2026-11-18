package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_help_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_help_helpCmd).Standalone()

	threads_helpCmd.AddCommand(threads_help_helpCmd)
}
