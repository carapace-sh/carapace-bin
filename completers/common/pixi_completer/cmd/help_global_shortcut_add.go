package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_shortcut_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a shortcut from an environment to your machine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_shortcut_addCmd).Standalone()

	help_global_shortcutCmd.AddCommand(help_global_shortcut_addCmd)
}
