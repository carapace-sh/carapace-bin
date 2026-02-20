package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_shortcut_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove shortcuts from your machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_shortcut_removeCmd).Standalone()

	help_global_shortcutCmd.AddCommand(help_global_shortcut_removeCmd)
}
