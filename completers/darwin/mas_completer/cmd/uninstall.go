package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall apps installed from the App Store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()
	uninstallCmd.Flags().Bool("all", false, "Uninstall all App Store apps")
	uninstallCmd.Flags().Bool("bundle", false, "Process all app IDs as bundle IDs")
	uninstallCmd.Flags().Bool("dry-run", false, "Perform dry run (show what would be uninstalled)")
	rootCmd.AddCommand(uninstallCmd)
}
