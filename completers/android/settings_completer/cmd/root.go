package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "settings",
	Short: "Settings provider commands",
	Long:  "https://cs.android.com/android/platform/superproject/main/+/main:frameworks/base/packages/SettingsProvider/src/com/android/providers/settings/SettingsService.java",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

}
