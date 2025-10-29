package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var self_help_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_help_uninstallCmd).Standalone()

	self_helpCmd.AddCommand(self_help_uninstallCmd)
}
