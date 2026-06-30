package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfUninstallCmd = &cobra.Command{
	Use:   "self-uninstall",
	Short: "Uninstall swiftly itself",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfUninstallCmd).Standalone()

	selfUninstallCmd.Flags().BoolP("help", "h", false, "Show help information")

	rootCmd.AddCommand(selfUninstallCmd)
}
