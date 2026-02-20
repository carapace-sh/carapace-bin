package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_shortcut_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a shortcut from an environment to your machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_shortcut_help_addCmd).Standalone()

	global_shortcut_helpCmd.AddCommand(global_shortcut_help_addCmd)
}
