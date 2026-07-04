package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simctl",
	Short: "control the iOS Simulator",
	Long:  "https://keith.github.io/xcode-manpages/simctl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"create", "Create a new device",
			"delete", "Delete a device or set of devices",
			"erase", "Erase a device's contents",
			"boot", "Boot a device",
			"shutdown", "Shut down a device",
			"list", "List available devices, devicetypes, runtimes, or pairs",
			"install", "Install an application on a device",
			"uninstall", "Uninstall an application from a device",
			"launch", "Launch an application on a device",
			"terminate", "Terminate an application on a device",
			"spawn", "Spawn a process on a device",
			"add_media", "Add media to a device",
			"io", "Manage I/O for a device",
			"diagnose", "Collect diagnostic information",
			"override", "Override simulator preferences",
			"get_app_container", "Get the path of an app's container",
			"unpair", "Unpair a device",
			"pair", "Pair a watch with a phone",
			"push", "Send a push notification to a device",
			"status_bar", "Toggle status bar overrides",
			"keychain", "Manage the keychain on a device",
			"ui", "Manage the UI of a device",
			"logic_test", "Run a logic test",
			"copy_device_container", "Copy an app's container to the host",
		),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		simctl.ActionDevices(),
	)
}
