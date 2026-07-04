package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userTrustSettingsEnableCmd = &cobra.Command{
	Use:   "user-trust-settings-enable",
	Short: "Display or manipulate user-level trust settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userTrustSettingsEnableCmd).Standalone()
	userTrustSettingsEnableCmd.Flags().BoolP("disable", "d", false, "Disable user-level Trust Settings")
	userTrustSettingsEnableCmd.Flags().BoolP("enable", "e", false, "Enable user-level Trust Settings")
	rootCmd.AddCommand(userTrustSettingsEnableCmd)
}
