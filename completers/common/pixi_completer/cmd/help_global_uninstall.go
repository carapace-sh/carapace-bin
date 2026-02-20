package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls environments from the global environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_uninstallCmd).Standalone()

	help_globalCmd.AddCommand(help_global_uninstallCmd)
}
