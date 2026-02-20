package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls environments from the global environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_uninstallCmd).Standalone()

	global_helpCmd.AddCommand(global_help_uninstallCmd)
}
