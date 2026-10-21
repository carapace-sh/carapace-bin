package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_ipaCmd = &cobra.Command{
	Use:   "ipa",
	Short: "Build an iOS archive bundle and IPA for distribution",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_ipaCmd).Standalone()

	build_ipaCmd.Flags().Bool("analyze-size", false, "Whether to produce additional profile information for artifact output size. This flag is only supported on \"--release\" builds. When building for Android, a single ABI must be specified at a time with the \"--target-platform\" flag. When building for iOS, only the symbols from the arm64 architecture are used to analyze code size. This flag cannot be combined with \"--split-debug-info\".")
	build_ipaCmd.Flags().String("build-name", "", "A \"x.y.z\" string used as the version number shown to users. For each new version of your app, you will provide a version number to differentiate it.")
	build_ipaCmd.Flags().String("build-number", "", "An identifier used as an internal version number. Each build must have a unique identifier to differentiate it from previous builds.")
	build_ipaCmd.Flags().String("code-size-directory", "", "The location to write code size analysis files. If this is not specified, files are written to a temporary directory under the build directory.")
	build_ipaCmd.Flags().Bool("codesign", false, "Codesign the application bundle (only available on device builds).")
	build_ipaCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_ipaCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	build_ipaCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_ipaCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_ipaCmd.Flags().StringSlice("enable-experiment", nil, "The name of an experimental Dart feature to enable. Multiple experiments can be enabled by repeating \"--enable-experiment\".")
	build_ipaCmd.Flags().String("export-method", "", "Specify how the IPA will be distributed.")
	build_ipaCmd.Flags().String("export-options-plist", "", "Export an IPA with these options. See \"xcodebuild -h\" for available exportOptionsPlist keys.")
	build_ipaCmd.Flags().StringSlice("extra-front-end-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart front end. For example, \"--extra-front-end-options=--enable-experiment=nonfunction-type-aliases\".")
	build_ipaCmd.Flags().StringSlice("extra-gen-snapshot-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart native compiler. (Requires the \"--release\", \"--profile\", or \"--jit-release\" flag.)")
	build_ipaCmd.Flags().Bool("fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_ipaCmd.Flags().String("flavor", "", "Build a custom app flavor as defined by platform-specific build setup. Supports the use of product flavors in Android Gradle scripts, and the use of custom Xcode schemes. Overrides the value of the \"default-flavor\" entry in the flutter pubspec.")
	build_ipaCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_ipaCmd.Flags().Bool("jit-release", false, "Build a JIT release version of your app.")
	build_ipaCmd.Flags().Bool("no-analyze-size", false, "Whether to produce additional profile information for artifact output size. This flag is only supported on \"--release\" builds. When building for Android, a single ABI must be specified at a time with the \"--target-platform\" flag. When building for iOS, only the symbols from the arm64 architecture are used to analyze code size. This flag cannot be combined with \"--split-debug-info\".")
	build_ipaCmd.Flags().Bool("no-codesign", false, "Codesign the application bundle (only available on device builds).")
	build_ipaCmd.Flags().Bool("no-fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_ipaCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_ipaCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_ipaCmd.Flags().Bool("no-tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	build_ipaCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_ipaCmd.Flags().String("performance-measurement-file", "", "Output individual target performance to a JSON file.")
	build_ipaCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_ipaCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_ipaCmd.Flags().Bool("release", false, "Build a release version of your app.")
	build_ipaCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application. The value of the flag should be a directory where program symbol files can be stored for later use. These symbol files contain the information needed to symbolize Dart stack traces. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. This flag cannot be combined with \"--analyze-size\".")
	build_ipaCmd.Flags().StringP("target", "t", "", "The main entrypoint file of the application, as run on the device. If the \"--target\" option is omitted, but a file name is provided on the command line, then that file is used instead.")
	build_ipaCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	buildCmd.AddCommand(build_ipaCmd)

	carapace.Gen(build_ipaCmd).FlagCompletion(carapace.ActionMap{
		"export-method": carapace.ActionValuesDescribed(
			"app-store", "Upload to the App Store.",
			"ad-hoc", "Test on designated devices that do not need to be registered with the Apple developer account. Requires a distribution certificate.",
			"development", "Test only on development devices registered with the Apple developer account.",
			"enterprise", "Distribute an app registered with the Apple Developer Enterprise Program.",
		),
		"split-debug-info": carapace.ActionDirectories(),
		"target":           carapace.ActionFiles(".dart"),
	})
}
