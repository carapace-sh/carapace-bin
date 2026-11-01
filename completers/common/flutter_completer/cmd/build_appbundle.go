package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_appbundleCmd = &cobra.Command{
	Use:     "appbundle",
	Aliases: []string{"aab"},
	Short:   "Build an Android App Bundle file from your app",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_appbundleCmd).Standalone()

	build_appbundleCmd.Flags().Bool("analyze-size", false, "Whether to produce additional profile information for artifact output size. This flag is only supported on \"--release\" builds. When building for Android, a single ABI must be specified at a time with the \"--target-platform\" flag. When building for iOS, only the symbols from the arm64 architecture are used to analyze code size. This flag cannot be combined with \"--split-debug-info\".")
	build_appbundleCmd.Flags().Bool("android-gradle-daemon", false, "Whether to enable the Gradle daemon when performing an Android build.")
	build_appbundleCmd.Flags().StringSliceP("android-project-arg", "P", nil, "Additional arguments specified as key=value that are passed directly to the gradle project via the -P flag. These can be accessed in build.gradle via the \"project.property\" API.")
	build_appbundleCmd.Flags().StringSlice("android-project-args", nil, "Additional arguments specified as key=value that are passed directly to the gradle project via the -P flag. These can be accessed in build.gradle via the \"project.property\" API.")
	build_appbundleCmd.Flags().String("android-project-cache-dir", "", "In an Android build, this flag allows the Gradle project cache directory to be specified to an absolute path. Setting this is roughly equivalent to setting the GRADLE_PROJECT_CACHE_DIR environment variable.")
	build_appbundleCmd.Flags().Bool("android-skip-build-dependency-validation", false, "Skips Android Gradle project dependency verification.")
	build_appbundleCmd.Flags().String("build-name", "", "A \"x.y.z\" string used as the version number shown to users. For each new version of your app, you will provide a version number to differentiate it.")
	build_appbundleCmd.Flags().String("build-number", "", "An identifier used as an internal version number. Each build must have a unique identifier to differentiate it from previous builds.")
	build_appbundleCmd.Flags().String("code-size-directory", "", "The location to write code size analysis files. If this is not specified, files are written to a temporary directory under the build directory.")
	build_appbundleCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_appbundleCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	build_appbundleCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_appbundleCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_appbundleCmd.Flags().Bool("deferred-components", false, "Setting to false disables building with deferred components. All deferred code will be compiled into the base app, and assets act as if they were defined under the regular assets section in pubspec.yaml. This flag has no effect on non-deferred components apps.")
	build_appbundleCmd.Flags().StringSlice("enable-experiment", nil, "The name of an experimental Dart feature to enable. Multiple experiments can be enabled by repeating \"--enable-experiment\".")
	build_appbundleCmd.Flags().Bool("enable-hcpp", false, "Enable the use of the HCPP platform view rendering mode on the Impeller rendering backend. An explicit value takes priority over the EnableHcpp metadata in AndroidManifest.xml: build commands write it into the manifest of the artifact they produce, and \"run\", \"test\", and \"drive\" additionally apply it at launch. Without the flag, the manifest decides, and if the manifest does not set it either, the enable-hcpp feature flag does.")
	build_appbundleCmd.Flags().StringSlice("extra-front-end-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart front end. For example, \"--extra-front-end-options=--enable-experiment=nonfunction-type-aliases\".")
	build_appbundleCmd.Flags().StringSlice("extra-gen-snapshot-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart native compiler. (Requires the \"--release\", \"--profile\", or \"--jit-release\" flag.)")
	build_appbundleCmd.Flags().Bool("fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_appbundleCmd.Flags().String("flavor", "", "Build a custom app flavor as defined by platform-specific build setup. Supports the use of product flavors in Android Gradle scripts, and the use of custom Xcode schemes. Overrides the value of the \"default-flavor\" entry in the flutter pubspec.")
	build_appbundleCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_appbundleCmd.Flags().Bool("ignore-deprecation", false, "Indicates that the app should ignore deprecation warnings and continue to build using deprecated APIs. Use of this flag may cause your app to fail to build when deprecated APIs are removed.")
	build_appbundleCmd.Flags().Bool("jit-release", false, "Build a JIT release version of your app.")
	build_appbundleCmd.Flags().Bool("no-analyze-size", false, "Whether to produce additional profile information for artifact output size. This flag is only supported on \"--release\" builds. When building for Android, a single ABI must be specified at a time with the \"--target-platform\" flag. When building for iOS, only the symbols from the arm64 architecture are used to analyze code size. This flag cannot be combined with \"--split-debug-info\".")
	build_appbundleCmd.Flags().Bool("no-android-gradle-daemon", false, "Whether to enable the Gradle daemon when performing an Android build.")
	build_appbundleCmd.Flags().Bool("no-android-skip-build-dependency-validation", false, "Skips Android Gradle project dependency verification.")
	build_appbundleCmd.Flags().Bool("no-deferred-components", false, "Setting to false disables building with deferred components. All deferred code will be compiled into the base app, and assets act as if they were defined under the regular assets section in pubspec.yaml. This flag has no effect on non-deferred components apps.")
	build_appbundleCmd.Flags().Bool("no-enable-hcpp", false, "Enable the use of the HCPP platform view rendering mode on the Impeller rendering backend. An explicit value takes priority over the EnableHcpp metadata in AndroidManifest.xml: build commands write it into the manifest of the artifact they produce, and \"run\", \"test\", and \"drive\" additionally apply it at launch. Without the flag, the manifest decides, and if the manifest does not set it either, the enable-hcpp feature flag does.")
	build_appbundleCmd.Flags().Bool("no-fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_appbundleCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_appbundleCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_appbundleCmd.Flags().Bool("no-shrink", false, "This flag has no effect. Code shrinking is always enabled in release builds. To learn more, see: https://developer.android.com/studio/build/shrink-code")
	build_appbundleCmd.Flags().Bool("no-track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	build_appbundleCmd.Flags().Bool("no-tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	build_appbundleCmd.Flags().Bool("no-validate-deferred-components", false, "When enabled, deferred component apps will fail to build if setup problems are detected that would prevent deferred components from functioning properly. The tooling also provides guidance on how to set up the project files to pass this verification. Disabling setup verification will always attempt to fully build the app regardless of any problems detected. Builds that are part of CI testing and advanced users with custom deferred components implementations should disable setup verification. This flag has no effect on non-deferred components apps.")
	build_appbundleCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation. This flag must always be combined with \"--split-debug-info\", the mapping between the values and the original identifiers is stored in the symbol map created in the specified directory. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. Because all identifiers are renamed, methods like Object.runtimeType, Type.toString, Enum.toString, Stacktrace.toString, Symbol.toString (for constant symbols or those generated by runtime system) will return obfuscated results. Any code or tests that rely on exact names will break.")
	build_appbundleCmd.Flags().String("performance-measurement-file", "", "Output individual target performance to a JSON file.")
	build_appbundleCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_appbundleCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_appbundleCmd.Flags().Bool("release", false, "Build a release version of your app.")
	build_appbundleCmd.Flags().Bool("shrink", false, "This flag has no effect. Code shrinking is always enabled in release builds. To learn more, see: https://developer.android.com/studio/build/shrink-code")
	build_appbundleCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application. The value of the flag should be a directory where program symbol files can be stored for later use. These symbol files contain the information needed to symbolize Dart stack traces. For an app built with this flag, the \"flutter symbolize\" command with the right program symbol file is required to obtain a human readable stack trace. This flag cannot be combined with \"--analyze-size\".")
	build_appbundleCmd.Flags().StringP("target", "t", "", "The main entrypoint file of the application, as run on the device. If the \"--target\" option is omitted, but a file name is provided on the command line, then that file is used instead.")
	build_appbundleCmd.Flags().StringSlice("target-platform", nil, "The target platform for which the app is compiled.")
	build_appbundleCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	build_appbundleCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	build_appbundleCmd.Flags().Bool("validate-deferred-components", false, "When enabled, deferred component apps will fail to build if setup problems are detected that would prevent deferred components from functioning properly. The tooling also provides guidance on how to set up the project files to pass this verification. Disabling setup verification will always attempt to fully build the app regardless of any problems detected. Builds that are part of CI testing and advanced users with custom deferred components implementations should disable setup verification. This flag has no effect on non-deferred components apps.")
	buildCmd.AddCommand(build_appbundleCmd)

	carapace.Gen(build_appbundleCmd).FlagCompletion(carapace.ActionMap{
		"android-project-cache-dir":    carapace.ActionDirectories(),
		"code-size-directory":          carapace.ActionDirectories(),
		"dart-define-from-file":        carapace.ActionFiles(".json", ".env"),
		"performance-measurement-file": carapace.ActionFiles(),
		"split-debug-info":             carapace.ActionDirectories(),
		"target":                       carapace.ActionFiles(".dart"),
		"target-platform":              carapace.ActionValues("android-arm", "android-arm64", "android-x64"),
	})
}
