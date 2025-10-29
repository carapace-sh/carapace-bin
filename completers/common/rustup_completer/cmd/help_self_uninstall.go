package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_self_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_self_uninstallCmd).Standalone()

	help_selfCmd.AddCommand(help_self_uninstallCmd)
}
