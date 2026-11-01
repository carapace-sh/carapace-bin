package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/flutter_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a Flutter app on an attached device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	installCmd.Flags().String("device-connection", "", "Discover devices based on connection type.")
	installCmd.Flags().StringP("device-id", "d", "", "Target device id or name (prefixes allowed).")
	installCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach. Longer timeouts may be necessary for networked devices.")
	installCmd.Flags().String("device-user", "", "Identifier number for a user or work profile on Android only. Run \"adb shell pm list users\" for available identifiers.")
	installCmd.Flags().String("flavor", "", "Build a custom app flavor as defined by platform-specific build setup. Supports the use of product flavors in Android Gradle scripts, and the use of custom Xcode schemes. Overrides the value of the \"default-flavor\" entry in the flutter pubspec.")
	installCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	installCmd.Flags().Bool("jit-release", false, "Build a JIT release version of your app.")
	installCmd.Flags().Bool("no-uninstall-only", false, "Uninstall the app if already on the device. Skip install.")
	installCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	installCmd.Flags().Bool("release", false, "Build a release version of your app.")
	installCmd.Flags().Bool("uninstall-only", false, "Uninstall the app if already on the device. Skip install.")
	installCmd.Flags().String("use-application-binary", "", "Specify a pre-built application binary to use when running. For Android applications, this must be the path to an APK. For iOS applications, the path to an IPA. Other device types do not yet support prebuilt application binaries.")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"device-connection": carapace.ActionValuesDescribed(
			"attached", "Only searches for devices connected by USB or built-in (such as simulators/emulators, MacOS/Windows, Chrome)",
			"both", "Searches for both attached and wireless devices.",
			"wireless", "Only searches for devices connected wirelessly. Discovering wireless devices may take longer.",
		),
		"device-id":              action.ActionDevices(),
		"device-user":            adb.ActionUsers(),
		"use-application-binary": carapace.ActionFiles(".apk", ".ipa"),
	})
}
