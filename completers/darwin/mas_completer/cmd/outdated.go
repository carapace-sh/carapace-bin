package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "List pending app updates from the App Store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()
	outdatedCmd.Flags().Bool("accurate", false, "Use accurate, slower update detection")
	outdatedCmd.Flags().Bool("bundle", false, "Process all app IDs as bundle IDs")
	outdatedCmd.Flags().Bool("check-min-os", true, "Check if macOS can install the latest app version")
	outdatedCmd.Flags().Bool("inaccurate", false, "Use inaccurate, faster logic")
	outdatedCmd.Flags().Bool("json", false, "Output JSON")
	outdatedCmd.Flags().Bool("no-check-min-os", false, "Skip the minimum macOS version check")
	outdatedCmd.Flags().Bool("verbose", false, "Warn about app IDs unknown to the App Store")
	rootCmd.AddCommand(outdatedCmd)
}
