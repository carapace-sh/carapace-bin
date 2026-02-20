package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_shortcutCmd = &cobra.Command{
	Use:   "shortcut",
	Short: "Interact with the shortcuts on your machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_shortcutCmd).Standalone()

	globalCmd.AddCommand(global_shortcutCmd)
}
