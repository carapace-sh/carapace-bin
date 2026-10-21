package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"configure"},
	Short:   "Configure Flutter settings",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().Bool("analytics", false, "Enable or disable reporting anonymously tool usage statistics and crash reports. (An alias for \"--enable-analytics\" and \"--disable-analytics\" top level flags.)")
	configCmd.Flags().String("android-sdk", "", "The Android SDK directory.")
	configCmd.Flags().String("android-studio-dir", "", "The Android Studio installation directory. If unset, flutter will search for valid installations at well-known locations.")
	configCmd.Flags().String("build-dir", "", "The relative path to override a projects build directory.")
	configCmd.Flags().Bool("clear-features", false, "Remove all configured features and restore them to the default values.")
	configCmd.Flags().Bool("clear-ios-signing-cert", false, "Clear the saved development certificate or provisioning profile choice used to sign apps for iOS device deployment.")
	configCmd.Flags().Bool("clear-ios-signing-settings", false, "Clear the saved development certificate or provisioning profile choice used to sign apps for iOS device deployment.")
	configCmd.Flags().Bool("cli-animations", false, "Enable or disable animations in the command line interface.")
	configCmd.Flags().Bool("enable-android", false, "Enable or disable Flutter for Android.")
	configCmd.Flags().Bool("enable-custom-devices", false, "Enable or disable early support for custom device types.")
	configCmd.Flags().Bool("enable-dart-data-assets", false, "Enable or disable Dart data assets building and bundling. This setting applies only to the master channel.")
	configCmd.Flags().Bool("enable-fuchsia", false, "Enable or disable Flutter for Fuchsia. This setting applies only to the master channel.")
	configCmd.Flags().Bool("enable-hcpp", false, "Enable or disable the HCPP platform view rendering mode. HCPP requires the Impeller rendering backend, and Android API 34 or above. Devices that do not support it fall back to each platform view's developer-chosen mode.")
	configCmd.Flags().Bool("enable-ios", false, "Enable or disable Flutter for iOS.")
	configCmd.Flags().Bool("enable-linux-desktop", false, "Enable or disable support for desktop on Linux.")
	configCmd.Flags().Bool("enable-lldb-debugging", false, "Enable or disable support for debugging with LLDB for physical iOS devices. If LLDB debugging is off, Xcode debugging is used instead. Only available for iOS 17 or newer devices. Requires Xcode 26 or greater.")
	configCmd.Flags().Bool("enable-macos-arm64-only", false, "Enable or disable building arm64 architecture only for non-debug macOS builds. If enabled, macOS release and profile builds generate Apple Silicon binaries instead of universal binaries. This feature is disabled by default, but will default to enabled in a future release, before Intel Mac support is eventually discontinued. See https://flutter.dev/go/macos-intel-deprecation for details.")
	configCmd.Flags().Bool("enable-macos-desktop", false, "Enable or disable support for desktop on macOS.")
	configCmd.Flags().Bool("enable-native-assets", false, "Enable or disable native assets compilation and bundling.")
	configCmd.Flags().Bool("enable-record-use", false, "Enable or disable record use experiment.")
	configCmd.Flags().Bool("enable-riscv64", false, "Enable or disable support for riscv64 architecture.")
	configCmd.Flags().Bool("enable-swift-package-manager", false, "Enable or disable support for Swift Package Manager for iOS and macOS.")
	configCmd.Flags().Bool("enable-tool-extensions", false, "Enable or disable support for tool extensions. This setting applies only to the master channel.")
	configCmd.Flags().Bool("enable-uiscene-migration", false, "Enable or disable support for migrating to UIScene lifecycle. If enabled, Flutter will migrate your app to iOS UIScene lifecycle if possible or otherwise instruct you to migrate manually.")
	configCmd.Flags().Bool("enable-web", false, "Enable or disable Flutter for web.")
	configCmd.Flags().Bool("enable-windowing", false, "Enable or disable support for windowing on macOS, Linux, and Windows. This setting applies only to the master channel.")
	configCmd.Flags().Bool("enable-windows-desktop", false, "Enable or disable support for desktop on Windows.")
	configCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	configCmd.Flags().String("jdk-dir", "", "The Java Development Kit (JDK) installation directory. If unset, flutter will search for one in the following order: 1) the JDK bundled with the latest installation of Android Studio, 2) the JDK found at the directory found in the JAVA_HOME environment variable, and 3) the directory containing the java binary found in the user's path.")
	configCmd.Flags().Bool("list", false, "List all settings and their current values.")
	configCmd.Flags().Bool("machine", false, "Outputs in a machine readable structured JSON format.")
	configCmd.Flags().Bool("no-analytics", false, "Enable or disable reporting anonymously tool usage statistics and crash reports. (An alias for \"--enable-analytics\" and \"--disable-analytics\" top level flags.)")
	configCmd.Flags().Bool("no-cli-animations", false, "Enable or disable animations in the command line interface.")
	configCmd.Flags().Bool("no-enable-android", false, "Enable or disable Flutter for Android.")
	configCmd.Flags().Bool("no-enable-custom-devices", false, "Enable or disable early support for custom device types.")
	configCmd.Flags().Bool("no-enable-dart-data-assets", false, "Enable or disable Dart data assets building and bundling. This setting applies only to the master channel.")
	configCmd.Flags().Bool("no-enable-fuchsia", false, "Enable or disable Flutter for Fuchsia. This setting applies only to the master channel.")
	configCmd.Flags().Bool("no-enable-hcpp", false, "Enable or disable the HCPP platform view rendering mode. HCPP requires the Impeller rendering backend, and Android API 34 or above. Devices that do not support it fall back to each platform view's developer-chosen mode.")
	configCmd.Flags().Bool("no-enable-ios", false, "Enable or disable Flutter for iOS.")
	configCmd.Flags().Bool("no-enable-linux-desktop", false, "Enable or disable support for desktop on Linux.")
	configCmd.Flags().Bool("no-enable-lldb-debugging", false, "Enable or disable support for debugging with LLDB for physical iOS devices. If LLDB debugging is off, Xcode debugging is used instead. Only available for iOS 17 or newer devices. Requires Xcode 26 or greater.")
	configCmd.Flags().Bool("no-enable-macos-arm64-only", false, "Enable or disable building arm64 architecture only for non-debug macOS builds. If enabled, macOS release and profile builds generate Apple Silicon binaries instead of universal binaries. This feature is disabled by default, but will default to enabled in a future release, before Intel Mac support is eventually discontinued. See https://flutter.dev/go/macos-intel-deprecation for details.")
	configCmd.Flags().Bool("no-enable-macos-desktop", false, "Enable or disable support for desktop on macOS.")
	configCmd.Flags().Bool("no-enable-native-assets", false, "Enable or disable native assets compilation and bundling.")
	configCmd.Flags().Bool("no-enable-record-use", false, "Enable or disable record use experiment.")
	configCmd.Flags().Bool("no-enable-riscv64", false, "Enable or disable support for riscv64 architecture.")
	configCmd.Flags().Bool("no-enable-swift-package-manager", false, "Enable or disable support for Swift Package Manager for iOS and macOS.")
	configCmd.Flags().Bool("no-enable-tool-extensions", false, "Enable or disable support for tool extensions. This setting applies only to the master channel.")
	configCmd.Flags().Bool("no-enable-uiscene-migration", false, "Enable or disable support for migrating to UIScene lifecycle. If enabled, Flutter will migrate your app to iOS UIScene lifecycle if possible or otherwise instruct you to migrate manually.")
	configCmd.Flags().Bool("no-enable-web", false, "Enable or disable Flutter for web.")
	configCmd.Flags().Bool("no-enable-windowing", false, "Enable or disable support for windowing on macOS, Linux, and Windows. This setting applies only to the master channel.")
	configCmd.Flags().Bool("no-enable-windows-desktop", false, "Enable or disable support for desktop on Windows.")
	configCmd.Flags().Bool("no-omit-legacy-version-file", false, "Enable or disable stops writing the legacy version file. If set, the file {FLUTTER_ROOT}/version is no longer written as part of the flutter tool execution; a newer file format has existed for some time in {FLUTTER_ROOT}/bin/cache/flutter.version.json.")
	configCmd.Flags().Bool("omit-legacy-version-file", false, "Enable or disable stops writing the legacy version file. If set, the file {FLUTTER_ROOT}/version is no longer written as part of the flutter tool execution; a newer file format has existed for some time in {FLUTTER_ROOT}/bin/cache/flutter.version.json.")
	configCmd.Flags().Bool("select-ios-signing-settings", false, "Complete prompt to select and save code signing settings used to sign apps for iOS device deployment.")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"android-sdk":        carapace.ActionDirectories(),
		"android-studio-dir": carapace.ActionDirectories(),
		"build-dir":          carapace.ActionDirectories(),
	})
}
