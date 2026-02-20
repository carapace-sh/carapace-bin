package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_shortcut_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a shortcut from an environment to your machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_shortcut_addCmd).Standalone()

	global_help_shortcutCmd.AddCommand(global_help_shortcut_addCmd)
}
