package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_shortcut_help_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_shortcut_help_helpCmd).Standalone()

	global_shortcut_helpCmd.AddCommand(global_shortcut_help_helpCmd)
}
