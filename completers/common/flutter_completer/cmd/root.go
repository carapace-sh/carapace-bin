package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flutter",
	Short: "Manage your Flutter app development",
	Long:  "https://flutter.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("ci", false, "Enable a set of CI-specific test debug settings.")
	rootCmd.Flags().Bool("color", false, "Whether to use terminal colors (requires support for ANSI escape sequences).")
	rootCmd.Flags().String("debug-logs-dir", "", "Path to a directory where logs for debugging may be added.")
	rootCmd.Flags().StringP("device-id", "d", "", "Target device id or name (prefixes allowed).")
	rootCmd.Flags().Bool("disable-analytics", false, "Disable telemetry reporting each time a flutter or dart command runs, until it is re-enabled.")
	rootCmd.Flags().Bool("enable-analytics", false, "Enable telemetry reporting each time a flutter or dart command runs.")
	rootCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.Flags().String("local-engine", "", "Name of a build output within the engine out directory, if you are building Flutter locally. Use this to select a specific version of the engine if you have built multiple engine targets. This path is relative to \"--local-engine-src-path\" (see above).")
	rootCmd.Flags().String("local-engine-host", "", "The host operating system for which engine artifacts should be selected, if you are building Flutter locally. This is only used when \"--local-engine\" is also specified. By default, the host is determined automatically, but you may need to specify this if you are building on one platform (e.g. MacOS ARM64) but intend to run Flutter on another (e.g. Android).")
	rootCmd.Flags().String("local-engine-src-path", "", "Path to your engine src directory, if you are building Flutter locally. Defaults to $FLUTTER_ENGINE if set, otherwise defaults to the path given in your pubspec.yaml dependency_overrides for sky_engine, if any.")
	rootCmd.Flags().String("local-web-sdk", "", "Name of a build output within the engine out directory, if you are building Flutter locally. Use this to select a specific version of the web sdk if you have built multiple engine targets. This path is relative to \"--local-engine-src-path\" (see above).")
	rootCmd.Flags().Bool("machine", false, "When used with the \"--version\" flag, outputs the information using JSON.")
	rootCmd.Flags().Bool("no-color", false, "Whether to use terminal colors (requires support for ANSI escape sequences).")
	rootCmd.Flags().Bool("no-version-check", false, "Allow Flutter to check for updates when this command runs.")
	rootCmd.Flags().Bool("no-wrap", false, "Toggles output word wrapping, regardless of whether or not the output is a terminal.")
	rootCmd.Flags().String("packages", "", "Path to your \"package_config.json\" file.")
	rootCmd.Flags().Bool("prefixed-errors", false, "Causes lines sent to stderr to be prefixed with \"ERROR:\".")
	rootCmd.Flags().Bool("print-dtd", false, "Print the address of the Dart Tooling Daemon, if one is hosted by the Flutter CLI.")
	rootCmd.Flags().Bool("quiet", false, "Reduce the amount of output from some commands.")
	rootCmd.Flags().Bool("show-test-device", false, "List the special \"flutter-tester\" device in device listings. This headless device is used to test Flutter tooling.")
	rootCmd.Flags().Bool("show-web-server-device", false, "List the special \"web-server\" device in device listings.")
	rootCmd.Flags().Bool("suppress-analytics", false, "Suppress analytics reporting for the current CLI invocation.")
	rootCmd.Flags().BoolP("verbose", "v", false, "Noisy logging, including all shell commands executed. If used with \"--help\", shows hidden options. If used with \"flutter doctor\", shows additional diagnostic information. (Use \"-vv\" to force verbose logging in those cases.)")
	rootCmd.Flags().Bool("version", false, "Reports the version of this tool.")
	rootCmd.Flags().Bool("version-check", false, "Allow Flutter to check for updates when this command runs.")
	rootCmd.Flags().Bool("wrap", false, "Toggles output word wrapping, regardless of whether or not the output is a terminal.")
	rootCmd.Flags().String("wrap-column", "", "Sets the output wrap column. If not set, uses the width of the terminal. No wrapping occurs if not writing to a terminal. Use \"--no-wrap\" to turn off wrapping when connected to a terminal.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"device-id": action.ActionDevices(),
	})
}
