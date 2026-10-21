package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_windowsCmd = &cobra.Command{
	Use:   "windows",
	Short: "Build a Windows desktop application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_windowsCmd).Standalone()

	build_windowsCmd.Flags().Bool("analyze-size", false, "Whether to produce additional profile information for artifact output size. This flag is only supported on \"--release\" builds. When building for Android, a single ABI must be specified at a time with the \"--target-platform\" flag. When building for iOS, only the symbols from the arm64 architecture are used to analyze code size. This flag cannot be combined with \"--split-debug-info\".")
	build_windowsCmd.Flags().String("build-name", "", "A \"x.y.z\" string used as the version number shown to users. For each new version of your app, you will provide a version number to differentiate it.")
	build_windowsCmd.Flags().String("build-number", "", "An identifier used as an internal version number. Each build must have a unique identifier to differentiate it from previous builds.")
	build_windowsCmd.Flags().String("code-size-directory", "", "The location to write code size analysis files. If this is not specified, files are written to a temporary directory under the build directory.")
	build_windowsCmd.Flags().Bool("config-only", false, "Update the project configuration without performing a build. This can be used in CI/CD process that create an archive to avoid performing duplicate work.")
	build_windowsCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_windowsCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	build_windowsCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_windowsCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_windowsCmd.Flags().StringSlice("enable-experiment", nil, "The name of an experimental Dart feature to enable. Multiple experiments can be enabled by repeating \"--enable-experiment\".")
	build_windowsCmd.Flags().StringSlice("extra-front-end-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart front end. For example, \"--extra-front-end-options=--enable-experiment=nonfunction-type-aliases\".")
	build_windowsCmd.Flags().StringSlice("extra-gen-snapshot-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart native compiler. (Requires the \"--release\", \"--profile\", or \"--jit-release\" flag.)")
	build_windowsCmd.Flags().Bool("fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_windowsCmd.Flags().String("flavor", "", "Build a custom app flavor as defined by platform-specific build setup. Supports the use of product flavors in Android Gradle scripts, and the use of custom Xcode schemes. Overrides the value of the \"default-flavor\" entry in the flutter pubspec.")
	build_windowsCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_windowsCmd.Flags().Bool("jit-release", false, "Build a JIT release version of your app.")
	build_windowsCmd.Flags().Bool("no-analyze-size", false, "Whether to produce additional profile information for artifact output size. This flag is only supported on \"--release\" builds. When building for Android, a single ABI must be specified at a time with the \"--target-platform\" flag. When building for iOS, only the symbols from the arm64 architecture are used to analyze code size. This flag cannot be combined with \"--split-debug-info\".")
	build_windowsCmd.Flags().Bool("no-config-only", false, "Update the project configuration without performing a build. This can be used in CI/CD process that create an archive to avoid performing duplicate work.")
	build_windowsCmd.Flags().Bool("no-fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_windowsCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_windowsCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_windowsCmd.Flags().Bool("no-track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	build_windowsCmd.Flags().Bool("no-tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	build_windowsCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_windowsCmd.Flags().String("performance-measurement-file", "", "Output individual target performance to a JSON file.")
	build_windowsCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_windowsCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_windowsCmd.Flags().Bool("release", false, "Build a release version of your app.")
	build_windowsCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application. The value of the flag should be a directory where program symbol files can be stored for later use. These symbol files contain the information needed to symbolize Dart stack traces. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. This flag cannot be combined with \"--analyze-size\".")
	build_windowsCmd.Flags().StringP("target", "t", "", "The main entrypoint file of the application, as run on the device. If the \"--target\" option is omitted, but a file name is provided on the command line, then that file is used instead.")
	build_windowsCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	build_windowsCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	buildCmd.AddCommand(build_windowsCmd)

	carapace.Gen(build_windowsCmd).FlagCompletion(carapace.ActionMap{
		"split-debug-info": carapace.ActionDirectories(),
		"target":           carapace.ActionFiles(".dart"),
	})
}
