package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run Flutter unit tests for the current project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("branch-coverage", false, "Whether to collect branch coverage information. Implies collecting coverage data.")
	testCmd.Flags().IntP("concurrency", "j", 0, "The number of concurrent test processes to run. This will be ignored when running integration tests.")
	testCmd.Flags().Bool("coverage", false, "Whether to collect coverage information.")
	testCmd.Flags().StringArray("coverage-package", nil, "A regular expression matching packages names to include in the coverage report (if coverage is enabled). If unset, matches the current package name.")
	testCmd.Flags().String("coverage-path", "", "Where to store coverage information (if coverage is enabled).")
	testCmd.Flags().Bool("cross-origin-isolation", false, "Adds the Cross-Origin-Opener-Policy and Cross-Origin-Embedder-Policy headers to the web server. These headers are required for using APIs like SharedArrayBuffer. This is on by default for the \"skwasm\" web renderer, and this flag can be used to override the default. To disable this for the skwasm renderer, use \"--no-cross-origin-isolation\".")
	testCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	testCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	testCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	testCmd.Flags().Bool("dds", false, "Enable the Dart Developer Service (DDS). It may be necessary to disable this when attaching to an application with an existing DDS instance (e.g., attaching to an application currently connected to by \"flutter run\"), or when running certain tests. Disabling this feature may degrade IDE functionality if a DDS instance is not already connected to the target application.")
	testCmd.Flags().Int("dds-port", 0, "When this value is provided, the Dart Development Service (DDS) will be bound to the provided port. Specifying port 0 (the default) will find a random free port.")
	testCmd.Flags().Bool("deprecated-js-interop", false, "Whether to allow the deprecated JS interop libraries (such as \"dart:html\", \"dart:js\", and \"dart:js_util\") when compiling to JavaScript. With \"--no-deprecated-js-interop\", importing these libraries is a compile-time error and conditional imports on them resolve to false. Use this to find code that still needs to migrate to \"package:web\" and \"dart:js_interop\". If not specified, the Dart compiler default is used. Has no effect when compiling to WebAssembly.")
	testCmd.Flags().String("device-user", "", "Identifier number for a user or work profile on Android only. Run \"adb shell pm list users\" for available identifiers.")
	testCmd.Flags().Bool("disable-dds", false, "(deprecated; use \"--no-dds\" instead) Disable the Dart Developer Service (DDS).")
	testCmd.Flags().Bool("disable-service-auth-codes", false, "(deprecated) Allow connections to the VM service without using authentication codes. (Not recommended! This can open your device to remote code execution attacks!)")
	testCmd.Flags().Bool("disable-service-origin-check", false, "Allow connections to the VM service from any origin. (Not recommended. This can open your device to remote code execution attacks.)")
	testCmd.Flags().StringSlice("enable-experiment", nil, "The name of an experimental Dart feature to enable. Multiple experiments can be enabled by repeating \"--enable-experiment\".")
	testCmd.Flags().Bool("enable-flutter-gpu", false, "Whether to enable the Flutter GPU API (https://api.flutter.dev/flutter/flutter_gpu/). This feature is only supported with the Impeller rendering engine, which can be enabled via the \"--enable-impeller\" option.")
	testCmd.Flags().Bool("enable-hcpp", false, "Enable the use of the HCPP platform view rendering mode on the Impeller rendering backend. An explicit value takes priority over the EnableHcpp metadata in AndroidManifest.xml: build commands write it into the manifest of the artifact they produce, and \"run\", \"test\", and \"drive\" additionally apply it at launch. Without the flag, the manifest decides, and if the manifest does not set it either, the enable-hcpp feature flag does.")
	testCmd.Flags().Bool("enable-impeller", false, "Whether to enable the Impeller rendering engine. Impeller is the default renderer on iOS. On Android, Impeller is available but not the default. This flag will cause Impeller to be used on Android. On other platforms, this flag will be ignored.")
	testCmd.Flags().Bool("enable-vmservice", false, "Enables the VM service without \"--start-paused\". This flag is intended for use with tests that will use \"dart:developer\" to interact with the VM service at runtime. This flag is ignored if \"--start-paused\" or coverage are requested, as the VM service will be enabled in those cases regardless.")
	testCmd.Flags().StringArrayP("exclude-tags", "x", nil, "Run only tests that do not have the specified tags. See: https://pub.dev/packages/test#tagging-tests")
	testCmd.Flags().Bool("experimental-faster-testing", false, "Run each test in a separate lightweight Flutter Engine to speed up testing.")
	testCmd.Flags().Bool("fail-fast", false, "Stop running tests after the first failure.")
	testCmd.Flags().Bool("fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	testCmd.Flags().String("file-reporter", "", "Enable an additional reporter writing test results to a file. Should be in the form <reporter>:<filepath>, Example: \"json:reports/tests.json\".")
	testCmd.Flags().String("flavor", "", "Build a custom app flavor as defined by platform-specific build setup. Supports the use of product flavors in Android Gradle scripts, and the use of custom Xcode schemes. Overrides the value of the \"default-flavor\" entry in the flutter pubspec.")
	testCmd.Flags().String("frontend-server-starter-path", "", "When this value is provided, the frontend server will be started in JIT mode from the specified file, instead of from the AOT snapshot shipped with the Dart SDK. The specified file can either be a dart source file or a kernel file. This flag is not used by most normal users and is for backend testing only.")
	testCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	testCmd.Flags().Bool("ignore-timeouts", false, "Ignore all timeouts. Useful when testing a big application that requires a longer time to compile (e.g. running integration tests for a Flutter app).")
	testCmd.Flags().Bool("ipv6", false, "Whether to use IPv6 for the test harness server socket.")
	testCmd.Flags().Bool("machine", false, "Outputs in a machine readable structured JSON format.")
	testCmd.Flags().Bool("merge-coverage", false, "Whether to merge coverage data with \"coverage/lcov.base.info\". Implies collecting coverage data. (Requires lcov.)")
	testCmd.Flags().StringArray("name", nil, "A regular expression matching substrings of the names of tests to run.")
	testCmd.Flags().Bool("no-cross-origin-isolation", false, "Adds the Cross-Origin-Opener-Policy and Cross-Origin-Embedder-Policy headers to the web server. These headers are required for using APIs like SharedArrayBuffer. This is on by default for the \"skwasm\" web renderer, and this flag can be used to override the default. To disable this for the skwasm renderer, use \"--no-cross-origin-isolation\".")
	testCmd.Flags().Bool("no-dds", false, "Enable the Dart Developer Service (DDS). It may be necessary to disable this when attaching to an application with an existing DDS instance (e.g., attaching to an application currently connected to by \"flutter run\"), or when running certain tests. Disabling this feature may degrade IDE functionality if a DDS instance is not already connected to the target application.")
	testCmd.Flags().Bool("no-deprecated-js-interop", false, "Whether to allow the deprecated JS interop libraries (such as \"dart:html\", \"dart:js\", and \"dart:js_util\") when compiling to JavaScript. With \"--no-deprecated-js-interop\", importing these libraries is a compile-time error and conditional imports on them resolve to false. Use this to find code that still needs to migrate to \"package:web\" and \"dart:js_interop\". If not specified, the Dart compiler default is used. Has no effect when compiling to WebAssembly.")
	testCmd.Flags().Bool("no-disable-dds", false, "(deprecated; use \"--no-dds\" instead) Disable the Dart Developer Service (DDS).")
	testCmd.Flags().Bool("no-enable-flutter-gpu", false, "Whether to enable the Flutter GPU API (https://api.flutter.dev/flutter/flutter_gpu/). This feature is only supported with the Impeller rendering engine, which can be enabled via the \"--enable-impeller\" option.")
	testCmd.Flags().Bool("no-enable-hcpp", false, "Enable the use of the HCPP platform view rendering mode on the Impeller rendering backend. An explicit value takes priority over the EnableHcpp metadata in AndroidManifest.xml: build commands write it into the manifest of the artifact they produce, and \"run\", \"test\", and \"drive\" additionally apply it at launch. Without the flag, the manifest decides, and if the manifest does not set it either, the enable-hcpp feature flag does.")
	testCmd.Flags().Bool("no-enable-impeller", false, "Whether to enable the Impeller rendering engine. Impeller is the default renderer on iOS. On Android, Impeller is available but not the default. This flag will cause Impeller to be used on Android. On other platforms, this flag will be ignored.")
	testCmd.Flags().Bool("no-enable-vmservice", false, "Enables the VM service without \"--start-paused\". This flag is intended for use with tests that will use \"dart:developer\" to interact with the VM service at runtime. This flag is ignored if \"--start-paused\" or coverage are requested, as the VM service will be enabled in those cases regardless.")
	testCmd.Flags().Bool("no-fail-fast", false, "Stop running tests after the first failure.")
	testCmd.Flags().Bool("no-fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	testCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	testCmd.Flags().Bool("no-run-skipped", false, "Run skipped tests instead of skipping them.")
	testCmd.Flags().Bool("no-test-assets", false, "Whether to build the assets bundle for testing. This takes additional time before running the tests. Consider using \"--no-test-assets\" if assets are not required.")
	testCmd.Flags().Bool("no-track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	testCmd.Flags().Bool("no-uninstall", false, "Whether to uninstall the app after running integration tests. Set \"--no-uninstall\" to keep the app installed on the device.")
	testCmd.Flags().StringArray("plain-name", nil, "A plain-text substring of the names of tests to run.")
	testCmd.Flags().String("platform", "", "Selects the test backend.")
	testCmd.Flags().StringArrayP("preset", "P", nil, "The configuration preset(s) to use. Presets are defined in \"dart_test.yaml\".")
	testCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	testCmd.Flags().StringP("reporter", "r", "", "Set how to print test results. If unset, value will default to either compact or expanded.")
	testCmd.Flags().Bool("run-skipped", false, "Run skipped tests instead of skipping them.")
	testCmd.Flags().Int("shard-index", 0, "Tests can be sharded with the \"--total-shards\" and \"--shard-index\" arguments, allowing you to split up your test suites and run them separately.")
	testCmd.Flags().Bool("start-paused", false, "Start in a paused mode and wait for a debugger to connect. You must specify a single test file to run, explicitly. Instructions for connecting with a debugger are printed to the console once the test has started.")
	testCmd.Flags().StringArrayP("tags", "t", nil, "Run only tests associated with the specified tags. See: https://pub.dev/packages/test#tagging-tests")
	testCmd.Flags().Bool("test-assets", false, "Whether to build the assets bundle for testing. This takes additional time before running the tests. Consider using \"--no-test-assets\" if assets are not required.")
	testCmd.Flags().String("test-randomize-ordering-seed", "", "The seed to randomize the execution order of test cases within test files. Must be a 32bit unsigned integer or the string \"random\", which indicates that a seed should be selected randomly. By default, tests run in the order they are declared.")
	testCmd.Flags().String("timeout", "", "The default timeout for individual tests, specified either in seconds (e.g. \"60s\"), as a multiplier of the default test timeout (e.g. \"2x\"), or as the string \"none\" to disable test timeouts entirely. This value does not apply to the default test suite loading timeout.")
	testCmd.Flags().Int("total-shards", 0, "Tests can be sharded with the \"--total-shards\" and \"--shard-index\" arguments, allowing you to split up your test suites and run them separately.")
	testCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	testCmd.Flags().Bool("uninstall", false, "Whether to uninstall the app after running integration tests. Set \"--no-uninstall\" to keep the app installed on the device.")
	testCmd.Flags().Bool("update-goldens", false, "Whether \"matchesGoldenFile()\" calls within your test methods should update the golden files rather than test for an existing match.")
	testCmd.Flags().Bool("wasm", false, "Compile to WebAssembly rather than JavaScript. See https://flutter.dev/to/wasm for more information.")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"coverage-path":         carapace.ActionFiles(),
		"dart-define-from-file": carapace.ActionFiles(".json", ".env"),
		"device-user":           adb.ActionUsers(),
		"file-reporter": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("compact", "expanded", "failures-only", "github", "json", "silent").Suffix(":")
			case 1:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
		"frontend-server-starter-path": carapace.ActionFiles(),
		"platform": carapace.ActionValuesDescribed(
			"tester", "Run tests using the VM-based test environment.",
			"chrome", "(deprecated) Run tests using the Google Chrome web browser. This value is intended for testing the Flutter framework itself and may be removed at any time.",
		),
		"reporter": carapace.ActionValuesDescribed(
			"compact", "A single line, updated continuously (the default).",
			"expanded", "A separate line for each update. May be preferred when logging to a file or in continuous integration.",
			"failures-only", "A separate line for failing tests, with no output for passing tests.",
			"github", "A custom reporter for GitHub Actions (the default reporter when running on GitHub Actions).",
			"json", "A machine-readable format. See: https://dart.dev/go/test-docs/json_reporter.md",
			"silent", "A reporter with no output. May be useful when only the exit code is meaningful.",
		),
	})

	carapace.Gen(testCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".dart"),
	)
}
