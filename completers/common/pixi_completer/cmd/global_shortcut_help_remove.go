package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_shortcut_help_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove shortcuts from your machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_shortcut_help_removeCmd).Standalone()

	global_shortcut_helpCmd.AddCommand(global_shortcut_help_removeCmd)
}
