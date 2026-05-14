package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ptyProxy_help_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ptyProxy_help_helpCmd).Standalone()

	ptyProxy_helpCmd.AddCommand(ptyProxy_help_helpCmd)
}
