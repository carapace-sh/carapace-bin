package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_aarCmd = &cobra.Command{
	Use:   "aar",
	Short: "Build a repository containing an AAR and a POM file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_aarCmd).Standalone()

	build_aarCmd.Flags().Bool("android-gradle-daemon", false, "Whether to enable the Gradle daemon when performing an Android build.")
	build_aarCmd.Flags().StringSliceP("android-project-arg", "P", nil, "Additional arguments specified as key=value that are passed directly to the gradle project via the -P flag. These can be accessed in build.gradle via the \"project.property\" API.")
	build_aarCmd.Flags().StringSlice("android-project-args", nil, "Additional arguments specified as key=value that are passed directly to the gradle project via the -P flag. These can be accessed in build.gradle via the \"project.property\" API.")
	build_aarCmd.Flags().String("android-project-cache-dir", "", "In an Android build, this flag allows the Gradle project cache directory to be specified to an absolute path. Setting this is roughly equivalent to setting the GRADLE_PROJECT_CACHE_DIR environment variable.")
	build_aarCmd.Flags().Bool("android-skip-build-dependency-validation", false, "Skips Android Gradle project dependency verification.")
	build_aarCmd.Flags().String("build-number", "", "An identifier used as an internal version number. Each build must have a unique identifier to differentiate it from previous builds.")
	build_aarCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_aarCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	build_aarCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_aarCmd.Flags().Bool("debug", false, "Build a debug version of the current project.")
	build_aarCmd.Flags().StringSlice("enable-experiment", nil, "The name of an experimental Dart feature to enable. Multiple experiments can be enabled by repeating \"--enable-experiment\".")
	build_aarCmd.Flags().StringSlice("extra-front-end-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart front end. For example, \"--extra-front-end-options=--enable-experiment=nonfunction-type-aliases\".")
	build_aarCmd.Flags().StringSlice("extra-gen-snapshot-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart native compiler. (Requires the \"--release\", \"--profile\", or \"--jit-release\" flag.)")
	build_aarCmd.Flags().Bool("fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_aarCmd.Flags().String("flavor", "", "Build a custom app flavor as defined by platform-specific build setup. Supports the use of product flavors in Android Gradle scripts, and the use of custom Xcode schemes. Overrides the value of the \"default-flavor\" entry in the flutter pubspec.")
	build_aarCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_aarCmd.Flags().Bool("no-android-gradle-daemon", false, "Whether to enable the Gradle daemon when performing an Android build.")
	build_aarCmd.Flags().Bool("no-android-skip-build-dependency-validation", false, "Skips Android Gradle project dependency verification.")
	build_aarCmd.Flags().Bool("no-debug", false, "Build a debug version of the current project.")
	build_aarCmd.Flags().Bool("no-fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_aarCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_aarCmd.Flags().Bool("no-profile", false, "Build a version of the current project specialized for performance profiling.")
	build_aarCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_aarCmd.Flags().Bool("no-release", false, "Build a release version of the current project.")
	build_aarCmd.Flags().Bool("no-track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	build_aarCmd.Flags().Bool("no-tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	build_aarCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_aarCmd.Flags().StringP("output", "o", "", "The absolute path to the directory where the repository is generated. By default, this is <current-directory>/build/<target-platform>. Currently supported for subcommands: aar, web.")
	build_aarCmd.Flags().String("output-dir", "", "The absolute path to the directory where the repository is generated. By default, this is <current-directory>/build/<target-platform>. Currently supported for subcommands: aar, web.")
	build_aarCmd.Flags().Bool("profile", false, "Build a version of the current project specialized for performance profiling.")
	build_aarCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_aarCmd.Flags().Bool("release", false, "Build a release version of the current project.")
	build_aarCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application. The value of the flag should be a directory where program symbol files can be stored for later use. These symbol files contain the information needed to symbolize Dart stack traces. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. This flag cannot be combined with \"--analyze-size\".")
	build_aarCmd.Flags().StringSlice("target-platform", nil, "The target platform for which the app is compiled.")
	build_aarCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	build_aarCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	buildCmd.AddCommand(build_aarCmd)

	carapace.Gen(build_aarCmd).FlagCompletion(carapace.ActionMap{
		"output":           carapace.ActionDirectories(),
		"output-dir":       carapace.ActionDirectories(),
		"split-debug-info": carapace.ActionDirectories(),
		"target-platform":  carapace.ActionValues("android-arm", "android-arm64", "android-x64"),
	})
}
