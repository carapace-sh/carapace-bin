package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var setKeychainSettingsCmd = &cobra.Command{
	Use:   "set-keychain-settings",
	Short: "Set settings for a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setKeychainSettingsCmd).Standalone()
	setKeychainSettingsCmd.Flags().BoolP("lock-after-timeout", "u", false, "Lock keychain after timeout interval")
	setKeychainSettingsCmd.Flags().BoolP("lock-on-sleep", "l", false, "Lock keychain when the system sleeps")
	setKeychainSettingsCmd.Flags().StringP("timeout", "t", "", "Timeout interval in seconds (omit for no timeout)")
	rootCmd.AddCommand(setKeychainSettingsCmd)
	carapace.Gen(setKeychainSettingsCmd).PositionalAnyCompletion(security.ActionKeychains())
}
