package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_shortcut_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove shortcuts from your machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_shortcut_removeCmd).Standalone()

	global_help_shortcutCmd.AddCommand(global_help_shortcut_removeCmd)
}
