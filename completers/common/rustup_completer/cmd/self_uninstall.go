package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var self_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_uninstallCmd).Standalone()

	self_uninstallCmd.Flags().BoolP("help", "h", false, "Print help")
	self_uninstallCmd.Flags().Bool("no-modify-path", false, "Do not clean up the `PATH` environment variable")
	self_uninstallCmd.Flags().BoolS("y", "y", false, "Disable confirmation prompt")
	selfCmd.AddCommand(self_uninstallCmd)
}
