package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"update"},
	Short:   "Update outdated apps from the App Store",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()
	upgradeCmd.Flags().Bool("accurate", false, "Use accurate, slower update detection")
	upgradeCmd.Flags().Bool("bundle", false, "Process all app IDs as bundle IDs")
	upgradeCmd.Flags().Bool("check-min-os", true, "Check if macOS can install the latest app version")
	upgradeCmd.Flags().Bool("force", false, "Force reinstall")
	upgradeCmd.Flags().Bool("inaccurate", false, "Use inaccurate, faster logic")
	upgradeCmd.Flags().Bool("no-check-min-os", false, "Skip the minimum macOS version check")
	upgradeCmd.Flags().Bool("verbose", false, "Warn about app IDs unknown to the App Store")
	rootCmd.AddCommand(upgradeCmd)
}
