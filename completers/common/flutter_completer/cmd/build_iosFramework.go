package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_iosFrameworkCmd = &cobra.Command{
	Use:   "ios-framework",
	Short: "Produces .xcframeworks for a Flutter project and its plugins for integration into existing, plain iOS Xcode projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_iosFrameworkCmd).Standalone()

	build_iosFrameworkCmd.Flags().Bool("cocoapods", false, "Produce a Flutter.podspec instead of an engine Flutter.xcframework (recommended if host app uses CocoaPods).")
	build_iosFrameworkCmd.Flags().Bool("codesign", false, "Whether to code-sign XCFrameworks.")
	build_iosFrameworkCmd.Flags().String("codesign-identity", "", "The identity to use for code-signing XCFrameworks. If an identity is not provided and \"codesign\" is enabled, a code signing identity will be selected automatically from the Flutter app's Xcode project settings or Flutter config. To see a list of valid identities run \"security find-identity -p codesigning -v\".")
	build_iosFrameworkCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_iosFrameworkCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	build_iosFrameworkCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_iosFrameworkCmd.Flags().Bool("debug", false, "Whether to produce a framework for the debug build configuration. By default, all build configurations are built.")
	build_iosFrameworkCmd.Flags().StringSlice("enable-experiment", nil, "The name of an experimental Dart feature to enable. Multiple experiments can be enabled by repeating \"--enable-experiment\".")
	build_iosFrameworkCmd.Flags().StringSlice("extra-front-end-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart front end. For example, \"--extra-front-end-options=--enable-experiment=nonfunction-type-aliases\".")
	build_iosFrameworkCmd.Flags().StringSlice("extra-gen-snapshot-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart native compiler. (Requires the \"--release\", \"--profile\", or \"--jit-release\" flag.)")
	build_iosFrameworkCmd.Flags().Bool("fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_iosFrameworkCmd.Flags().String("flavor", "", "Build a custom app flavor as defined by platform-specific build setup. Supports the use of product flavors in Android Gradle scripts, and the use of custom Xcode schemes. Overrides the value of the \"default-flavor\" entry in the flutter pubspec.")
	build_iosFrameworkCmd.Flags().BoolP("force", "f", false, "Force Flutter.podspec creation on the master channel. This is only intended for testing the tool itself.")
	build_iosFrameworkCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_iosFrameworkCmd.Flags().Bool("no-cocoapods", false, "Produce a Flutter.podspec instead of an engine Flutter.xcframework (recommended if host app uses CocoaPods).")
	build_iosFrameworkCmd.Flags().Bool("no-codesign", false, "Whether to code-sign XCFrameworks.")
	build_iosFrameworkCmd.Flags().Bool("no-debug", false, "Whether to produce a framework for the debug build configuration. By default, all build configurations are built.")
	build_iosFrameworkCmd.Flags().Bool("no-fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_iosFrameworkCmd.Flags().Bool("no-force", false, "Force Flutter.podspec creation on the master channel. This is only intended for testing the tool itself.")
	build_iosFrameworkCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_iosFrameworkCmd.Flags().Bool("no-plugins", false, "Whether to produce frameworks for the plugins. This is intended for cases where plugins are already being built separately.")
	build_iosFrameworkCmd.Flags().Bool("no-profile", false, "Whether to produce a framework for the profile build configuration. By default, all build configurations are built.")
	build_iosFrameworkCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_iosFrameworkCmd.Flags().Bool("no-release", false, "Whether to produce a framework for the release build configuration. By default, all build configurations are built.")
	build_iosFrameworkCmd.Flags().Bool("no-static", false, "Build plugins as static frameworks. Link on, but do not embed these frameworks in the existing Xcode project.")
	build_iosFrameworkCmd.Flags().Bool("no-tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	build_iosFrameworkCmd.Flags().Bool("no-universal", false, "(deprecated) Produce universal frameworks that include all valid architectures.")
	build_iosFrameworkCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_iosFrameworkCmd.Flags().StringP("output", "o", "", "Location to write the frameworks.")
	build_iosFrameworkCmd.Flags().Bool("plugins", false, "Whether to produce frameworks for the plugins. This is intended for cases where plugins are already being built separately.")
	build_iosFrameworkCmd.Flags().Bool("profile", false, "Whether to produce a framework for the profile build configuration. By default, all build configurations are built.")
	build_iosFrameworkCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_iosFrameworkCmd.Flags().Bool("release", false, "Whether to produce a framework for the release build configuration. By default, all build configurations are built.")
	build_iosFrameworkCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application. The value of the flag should be a directory where program symbol files can be stored for later use. These symbol files contain the information needed to symbolize Dart stack traces. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. This flag cannot be combined with \"--analyze-size\".")
	build_iosFrameworkCmd.Flags().Bool("static", false, "Build plugins as static frameworks. Link on, but do not embed these frameworks in the existing Xcode project.")
	build_iosFrameworkCmd.Flags().StringP("target", "t", "", "The main entrypoint file of the application, as run on the device. If the \"--target\" option is omitted, but a file name is provided on the command line, then that file is used instead.")
	build_iosFrameworkCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	build_iosFrameworkCmd.Flags().Bool("universal", false, "(deprecated) Produce universal frameworks that include all valid architectures.")
	build_iosFrameworkCmd.Flags().Bool("xcframework", false, "Produce xcframeworks that include all valid architectures.")
	buildCmd.AddCommand(build_iosFrameworkCmd)

	carapace.Gen(build_iosFrameworkCmd).FlagCompletion(carapace.ActionMap{
		"output":           carapace.ActionDirectories(),
		"split-debug-info": carapace.ActionDirectories(),
		"target":           carapace.ActionFiles(".dart"),
	})
}
