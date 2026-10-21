package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var driveCmd = &cobra.Command{
	Use:     "drive",
	Aliases: []string{"driver"},
	Short:   "Run integration tests for the project on an attached device or emulator",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(driveCmd).Standalone()

	driveCmd.Flags().Bool("adb-log-filtering", false, "Filter adb logs so that logs not emitted by the current flutter application are not displayed.")
	driveCmd.Flags().Bool("android-emulator", false, "Whether to perform Flutter Driver testing using an Android Emulator. Works only if \"browser-name\" is set to \"android-chrome\".")
	driveCmd.Flags().Bool("android-gradle-daemon", false, "Whether to enable the Gradle daemon when performing an Android build.")
	driveCmd.Flags().StringSliceP("android-project-arg", "P", nil, "Additional arguments specified as key=value that are passed directly to the gradle project via the -P flag. These can be accessed in build.gradle via the \"project.property\" API.")
	driveCmd.Flags().StringSlice("android-project-args", nil, "Additional arguments specified as key=value that are passed directly to the gradle project via the -P flag. These can be accessed in build.gradle via the \"project.property\" API.")
	driveCmd.Flags().String("android-project-cache-dir", "", "In an Android build, this flag allows the Gradle project cache directory to be specified to an absolute path. Setting this is roughly equivalent to setting the GRADLE_PROJECT_CACHE_DIR environment variable.")
	driveCmd.Flags().Bool("android-skip-build-dependency-validation", false, "Skips Android Gradle project dependency verification.")
	driveCmd.Flags().String("base-href", "", "Overrides the href attribute of the <base> tag in web/index.html. No change is done to web/index.html if this flag is not provided. The value has to start and end with a slash \"/\". For more information: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base")
	driveCmd.Flags().String("browser-dimension", "", "The dimension of the browser when running a Flutter Web test. Format is \"width x height[@dpr]\" where dpr is optional device pixel ratio. This will affect screenshot dimensions and all offset-related actions.")
	driveCmd.Flags().String("browser-name", "", "Name of the browser where tests will be executed.")
	driveCmd.Flags().Bool("build", false, "(deprecated) If necessary, build the app before running. To use an existing app, pass the \"--use-application-binary\" flag with an existing application artifact.")
	driveCmd.Flags().Bool("cache-startup-profile", false, "Caches the CPU profile collected before the first frame for startup analysis.")
	driveCmd.Flags().String("chrome-binary", "", "Location of the Chrome binary. Works only if \"browser-name\" is set to \"chrome\".")
	driveCmd.Flags().Bool("cross-origin-isolation", false, "Adds the Cross-Origin-Opener-Policy and Cross-Origin-Embedder-Policy headers to the web server. These headers are required for using APIs like SharedArrayBuffer. This is on by default for the \"skwasm\" web renderer, and this flag can be used to override the default. To disable this for the skwasm renderer, use \"--no-cross-origin-isolation\".")
	driveCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	driveCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	driveCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	driveCmd.Flags().StringSliceP("dart-entrypoint-args", "a", nil, "Pass a list of arguments to the Dart entrypoint at application startup. By default this is main(List<String> args). Specify this option multiple times each with one argument to pass multiple arguments to the Dart entrypoint. Currently this is only supported on desktop platforms.")
	driveCmd.Flags().String("dart-flags", "", "Pass a list of comma separated flags to the Dart instance at application startup. Flags passed through this option must be present on the allowlist defined within the Flutter engine. If a disallowed flag is encountered, the process will be terminated immediately. This flag is not available on the stable channel and is only applied in debug and profile modes. This option should only be used for experiments and should not be used by typical users.")
	driveCmd.Flags().Bool("dds", false, "Enable the Dart Developer Service (DDS). It may be necessary to disable this when attaching to an application with an existing DDS instance (e.g., attaching to an application currently connected to by \"flutter run\"), or when running certain tests. Disabling this feature may degrade IDE functionality if a DDS instance is not already connected to the target application.")
	driveCmd.Flags().Int("dds-port", 0, "When this value is provided, the Dart Development Service (DDS) will be bound to the provided port. Specifying port 0 (the default) will find a random free port.")
	driveCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	driveCmd.Flags().Bool("deprecated-js-interop", false, "Whether to allow the deprecated JS interop libraries (such as \"dart:html\", \"dart:js\", and \"dart:js_util\") when compiling to JavaScript. With \"--no-deprecated-js-interop\", importing these libraries is a compile-time error and conditional imports on them resolve to false. Use this to find code that still needs to migrate to \"package:web\" and \"dart:js_interop\". If not specified, the Dart compiler default is used. Has no effect when compiling to WebAssembly.")
	driveCmd.Flags().String("device-connection", "", "Discover devices based on connection type.")
	driveCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach. Longer timeouts may be necessary for networked devices.")
	driveCmd.Flags().String("device-user", "", "Identifier number for a user or work profile on Android only. Run \"adb shell pm list users\" for available identifiers.")
	driveCmd.Flags().String("device-vmservice-port", "", "Look for vmservice connections only from the specified port. Specifying port 0 (the default) will accept the first vmservice discovered.")
	driveCmd.Flags().Bool("devtools", false, "Enable (or disable, with \"--no-devtools\") the launching of the Flutter DevTools debugger and profiler. If \"--no-devtools\" is specified, \"--devtools-server-address\" is ignored.")
	driveCmd.Flags().String("devtools-server-address", "", "When this value is provided, the Flutter tool will not spin up a new DevTools server instance, and will instead use the one provided at the given address. Ignored if \"--no-devtools\" is specified.")
	driveCmd.Flags().Bool("disable-dds", false, "(deprecated; use \"--no-dds\" instead) Disable the Dart Developer Service (DDS).")
	driveCmd.Flags().Bool("disable-service-auth-codes", false, "(deprecated) Allow connections to the VM service without using authentication codes. (Not recommended! This can open your device to remote code execution attacks!)")
	driveCmd.Flags().Bool("disable-service-origin-check", false, "Allow connections to the VM service from any origin. (Not recommended. This can open your device to remote code execution attacks.)")
	driveCmd.Flags().String("driver", "", "The test file to run on the host (as opposed to the target file to run on the device). By default, this file has the same base name as the target file, but in the \"test_driver/\" directory instead, and with \"_test\" inserted just before the extension, so e.g. if the target is \"lib/main.dart\", the driver will be \"test_driver/main_test.dart\".")
	driveCmd.Flags().Int("driver-port", 0, "The port where Webdriver server is launched at.")
	driveCmd.Flags().Bool("enable-dart-profiling", false, "Whether the Dart VM sampling CPU profiler is enabled. This flag is only meaningful in debug and profile builds.")
	driveCmd.Flags().Bool("enable-embedder-api", false, "Whether to enable the experimental embedder API on iOS.")
	driveCmd.Flags().StringSlice("enable-experiment", nil, "The name of an experimental Dart feature to enable. Multiple experiments can be enabled by repeating \"--enable-experiment\".")
	driveCmd.Flags().Bool("enable-flutter-gpu", false, "Whether to enable the Flutter GPU API (https://api.flutter.dev/flutter/flutter_gpu/). This feature is only supported with the Impeller rendering engine, which can be enabled via the \"--enable-impeller\" option.")
	driveCmd.Flags().Bool("enable-hcpp", false, "Enable the use of the HCPP platform view rendering mode on the Impeller rendering backend. An explicit value takes priority over the EnableHcpp metadata in AndroidManifest.xml: build commands write it into the manifest of the artifact they produce, and \"run\", \"test\", and \"drive\" additionally apply it at launch. Without the flag, the manifest decides, and if the manifest does not set it either, the enable-hcpp feature flag does.")
	driveCmd.Flags().Bool("enable-impeller", false, "Whether to enable the Impeller rendering engine. Impeller is the default renderer on iOS. On Android, Impeller is available but not the default. This flag will cause Impeller to be used on Android. On other platforms, this flag will be ignored.")
	driveCmd.Flags().Bool("enable-software-rendering", false, "(deprecated) Enable rendering using the Skia software backend. This is useful when testing Flutter on emulators. By default, Flutter will attempt to either use OpenGL or Vulkan and fall back to software when neither is available. This option is not supported when using the Impeller rendering engine.")
	driveCmd.Flags().Bool("enable-vulkan-validation", false, "Enable vulkan validation on the Impeller rendering backend if Vulkan is in use and the validation layers are available to the application.")
	driveCmd.Flags().Bool("endless-trace-buffer", false, "Enable tracing to an infinite buffer, instead of a ring buffer. This is useful when recording large traces. To use an endless buffer to record startup traces, combine this with \"--trace-startup\".")
	driveCmd.Flags().Bool("fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	driveCmd.Flags().String("flavor", "", "Build a custom app flavor as defined by platform-specific build setup. Supports the use of product flavors in Android Gradle scripts, and the use of custom Xcode schemes. Overrides the value of the \"default-flavor\" entry in the flutter pubspec.")
	driveCmd.Flags().Bool("headless", false, "Whether the driver browser is going to be launched in headless mode.")
	driveCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	driveCmd.Flags().String("host-vmservice-port", "", "When a device-side vmservice port is forwarded to a host-side port, use this value as the host port. Specifying port 0 (the default) will find a random free host port.")
	driveCmd.Flags().Bool("ios-profile-debugger", false, "Whether to attach the LLDB debugger when running in profile mode on a physical iOS device. Only available with Xcode 26.")
	driveCmd.Flags().Bool("ipv6", false, "Binds to IPv6 localhost instead of IPv4 when the flutter tool forwards the host port to a device port.")
	driveCmd.Flags().Bool("jit-release", false, "Build a JIT release version of your app.")
	driveCmd.Flags().Bool("keep-app-running", false, "Will keep the Flutter application running when done testing. By default, \"flutter drive\" stops the application after tests are finished, and \"--keep-app-running\" overrides this. On the other hand, if \"--use-existing-app\" is specified, then \"flutter drive\" instead defaults to leaving the application running, and \"--no-keep-app-running\" overrides it.")
	driveCmd.Flags().Bool("native-null-assertions", false, "Enables additional runtime null checks in web applications to ensure the correct nullability of native (such as in dart:html) and external (such as with JS interop) types. This is enabled by default but only takes effect in sound mode. To report an issue with a null assertion failure in dart:html or the other dart web libraries, please file a bug at: https://github.com/dart-lang/sdk/issues/labels/web-libraries")
	driveCmd.Flags().Bool("no-adb-log-filtering", false, "Filter adb logs so that logs not emitted by the current flutter application are not displayed.")
	driveCmd.Flags().Bool("no-android-emulator", false, "Whether to perform Flutter Driver testing using an Android Emulator. Works only if \"browser-name\" is set to \"android-chrome\".")
	driveCmd.Flags().Bool("no-android-gradle-daemon", false, "Whether to enable the Gradle daemon when performing an Android build.")
	driveCmd.Flags().Bool("no-android-skip-build-dependency-validation", false, "Skips Android Gradle project dependency verification.")
	driveCmd.Flags().Bool("no-build", false, "(deprecated) If necessary, build the app before running. To use an existing app, pass the \"--use-application-binary\" flag with an existing application artifact.")
	driveCmd.Flags().Bool("no-cache-startup-profile", false, "Caches the CPU profile collected before the first frame for startup analysis.")
	driveCmd.Flags().Bool("no-cross-origin-isolation", false, "Adds the Cross-Origin-Opener-Policy and Cross-Origin-Embedder-Policy headers to the web server. These headers are required for using APIs like SharedArrayBuffer. This is on by default for the \"skwasm\" web renderer, and this flag can be used to override the default. To disable this for the skwasm renderer, use \"--no-cross-origin-isolation\".")
	driveCmd.Flags().Bool("no-dds", false, "Enable the Dart Developer Service (DDS). It may be necessary to disable this when attaching to an application with an existing DDS instance (e.g., attaching to an application currently connected to by \"flutter run\"), or when running certain tests. Disabling this feature may degrade IDE functionality if a DDS instance is not already connected to the target application.")
	driveCmd.Flags().Bool("no-deprecated-js-interop", false, "Whether to allow the deprecated JS interop libraries (such as \"dart:html\", \"dart:js\", and \"dart:js_util\") when compiling to JavaScript. With \"--no-deprecated-js-interop\", importing these libraries is a compile-time error and conditional imports on them resolve to false. Use this to find code that still needs to migrate to \"package:web\" and \"dart:js_interop\". If not specified, the Dart compiler default is used. Has no effect when compiling to WebAssembly.")
	driveCmd.Flags().Bool("no-devtools", false, "Enable (or disable, with \"--no-devtools\") the launching of the Flutter DevTools debugger and profiler. If \"--no-devtools\" is specified, \"--devtools-server-address\" is ignored.")
	driveCmd.Flags().Bool("no-disable-dds", false, "(deprecated; use \"--no-dds\" instead) Disable the Dart Developer Service (DDS).")
	driveCmd.Flags().Bool("no-enable-dart-profiling", false, "Whether the Dart VM sampling CPU profiler is enabled. This flag is only meaningful in debug and profile builds.")
	driveCmd.Flags().Bool("no-enable-embedder-api", false, "Whether to enable the experimental embedder API on iOS.")
	driveCmd.Flags().Bool("no-enable-flutter-gpu", false, "Whether to enable the Flutter GPU API (https://api.flutter.dev/flutter/flutter_gpu/). This feature is only supported with the Impeller rendering engine, which can be enabled via the \"--enable-impeller\" option.")
	driveCmd.Flags().Bool("no-enable-hcpp", false, "Enable the use of the HCPP platform view rendering mode on the Impeller rendering backend. An explicit value takes priority over the EnableHcpp metadata in AndroidManifest.xml: build commands write it into the manifest of the artifact they produce, and \"run\", \"test\", and \"drive\" additionally apply it at launch. Without the flag, the manifest decides, and if the manifest does not set it either, the enable-hcpp feature flag does.")
	driveCmd.Flags().Bool("no-enable-impeller", false, "Whether to enable the Impeller rendering engine. Impeller is the default renderer on iOS. On Android, Impeller is available but not the default. This flag will cause Impeller to be used on Android. On other platforms, this flag will be ignored.")
	driveCmd.Flags().Bool("no-enable-vulkan-validation", false, "Enable vulkan validation on the Impeller rendering backend if Vulkan is in use and the validation layers are available to the application.")
	driveCmd.Flags().Bool("no-fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	driveCmd.Flags().Bool("no-headless", false, "Whether the driver browser is going to be launched in headless mode.")
	driveCmd.Flags().Bool("no-keep-app-running", false, "Will keep the Flutter application running when done testing. By default, \"flutter drive\" stops the application after tests are finished, and \"--keep-app-running\" overrides this. On the other hand, if \"--use-existing-app\" is specified, then \"flutter drive\" instead defaults to leaving the application running, and \"--no-keep-app-running\" overrides it.")
	driveCmd.Flags().Bool("no-native-null-assertions", false, "Enables additional runtime null checks in web applications to ensure the correct nullability of native (such as in dart:html) and external (such as with JS interop) types. This is enabled by default but only takes effect in sound mode. To report an issue with a null assertion failure in dart:html or the other dart web libraries, please file a bug at: https://github.com/dart-lang/sdk/issues/labels/web-libraries")
	driveCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	driveCmd.Flags().Bool("no-publish-port", false, "Publish the VM service port over mDNS. Disable to prevent the local network permission app dialog in debug and profile build modes (iOS devices only).")
	driveCmd.Flags().Bool("no-start-paused", false, "Start in a paused mode and wait for a debugger to connect.")
	driveCmd.Flags().Bool("no-test-flag", false, "No-op flag for testing purposes; use for testing flag priorities only.")
	driveCmd.Flags().Bool("no-track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	driveCmd.Flags().Bool("no-uninstall-first", false, "Uninstall previous versions of the app on the device before reinstalling. Currently only supported on iOS.")
	driveCmd.Flags().Bool("no-web-allow-expose-url", false, "Enables daemon-to-editor requests (app.exposeUrl) for exposing URLs when running on remote machines.")
	driveCmd.Flags().Bool("no-web-enable-expression-evaluation", false, "Enables expression evaluation in the debugger.")
	driveCmd.Flags().Bool("no-web-resources-cdn", false, "Use WebAssembly, CanvasKit, and other web resources from a content delivery network (CDN). Set to \"--no-web-resources-cdn\" to embed all web resources locally in the built app.")
	driveCmd.Flags().Bool("no-web-run-headless", false, "Launches the browser in headless mode. Currently only Chrome supports this option.")
	driveCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	driveCmd.Flags().String("profile-memory", "", "Launch devtools and profile application memory, writing The output data to the file path provided to this argument as JSON.")
	driveCmd.Flags().Bool("profile-microtasks", false, "Enable collection of information about each microtask. Information about completed microtasks will be written to the \"Microtask\" timeline stream. Information about queued microtasks will be accessible from Dart / Flutter DevTools.")
	driveCmd.Flags().Bool("profile-startup", false, "Make the profiler discard new samples once the profiler sample buffer is full. When this flag is not set, the profiler sample buffer is used as a ring buffer, meaning that once it is full, new samples start overwriting the oldest ones.")
	driveCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	driveCmd.Flags().Bool("publish-port", false, "Publish the VM service port over mDNS. Disable to prevent the local network permission app dialog in debug and profile build modes (iOS devices only).")
	driveCmd.Flags().Bool("purge-persistent-cache", false, "Removes all existing persistent caches. This allows reproducing shader compilation jank that normally only happens the first time an app is run, or for reliable testing of compilation jank fixes (e.g. shader warm-up).")
	driveCmd.Flags().Bool("release", false, "Build a release version of your app.")
	driveCmd.Flags().String("route", "", "Which route to load when running the app.")
	driveCmd.Flags().String("screenshot", "", "Directory location to write screenshots on test failure.")
	driveCmd.Flags().Bool("skia-deterministic-rendering", false, "(deprecated) When combined with \"--enable-software-rendering\", this should provide completely deterministic (i.e. reproducible) Skia rendering. This is useful for testing purposes (e.g. when comparing screenshots). This option is not supported when using the Impeller rendering engine.")
	driveCmd.Flags().Bool("start-paused", false, "Start in a paused mode and wait for a debugger to connect.")
	driveCmd.Flags().StringP("target", "t", "", "The main entrypoint file of the application, as run on the device. If the \"--target\" option is omitted, but a file name is provided on the command line, then that file is used instead.")
	driveCmd.Flags().StringSlice("test-arguments", nil, "Additional arguments to pass to the Dart VM running The test script. This can be used to opt-in to use \"dart test\" as a runner for the test script, which allows, among other things, changing the reporter. For example, to opt-in to the \"expanded\" reporter, pass both \"test\" and \"--reporter=expanded\". Please leave feedback at <https://github.com/flutter/flutter/issues/152409>.")
	driveCmd.Flags().Bool("test-flag", false, "No-op flag for testing purposes; use for testing flag priorities only.")
	driveCmd.Flags().String("timeout", "", "Timeout the test after the given number of seconds. If the \"--screenshot\" option is provided, a screenshot will be taken before exiting. Defaults to no timeout.")
	driveCmd.Flags().String("trace-allowlist", "", "Filters out all trace events except those that are specified in this comma separated list of allowed prefixes.")
	driveCmd.Flags().Bool("trace-skia", false, "Enable tracing of Skia code. This is useful when debugging the raster thread (formerly known as the GPU thread). By default, Flutter will not log Skia code, as it introduces significant overhead that may affect recorded performance metrics in a misleading way.")
	driveCmd.Flags().String("trace-skia-allowlist", "", "Filters out all Skia trace events except those that are specified in this comma separated list of allowed prefixes.")
	driveCmd.Flags().Bool("trace-startup", false, "Trace application startup, then exit, saving the trace to a file. By default, this will be saved in the \"build\" directory. If the FLUTTER_TEST_OUTPUTS_DIR environment variable is set, the file will be written there instead.")
	driveCmd.Flags().Bool("trace-systrace", false, "Enable tracing to the system tracer. This is only useful on platforms where such a tracer is available (Android, iOS, macOS and Fuchsia).")
	driveCmd.Flags().String("trace-to-file", "", "Write the timeline trace to a file at the specified path. The file will be in Perfetto's proto format; it will be possible to load the file into Perfetto's trace viewer.")
	driveCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	driveCmd.Flags().Bool("uninstall-first", false, "Uninstall previous versions of the app on the device before reinstalling. Currently only supported on iOS.")
	driveCmd.Flags().String("use-application-binary", "", "Specify a pre-built application binary to use when running. For Android applications, this must be the path to an APK. For iOS applications, the path to an IPA. Other device types do not yet support prebuilt application binaries.")
	driveCmd.Flags().String("use-existing-app", "", "Connect to an already running instance via the given Dart VM Service URL. If this option is given, the application will not be automatically started, and it will only be stopped if \"--no-keep-app-running\" is explicitly set.")
	driveCmd.Flags().Bool("verbose-system-logs", false, "Include verbose logging from the Flutter engine.")
	driveCmd.Flags().String("vm-service-port", "", "(deprecated; use host-vmservice-port instead) Listen to the given port for a Dart VM Service connection. Specifying port 0 (the default) will find a random free port. if the Dart Development Service (DDS) is enabled, this will not be the port of the VmService instance advertised on the command line.")
	driveCmd.Flags().String("vmservice-out-file", "", "A file to write the attached vmservice URL to after an application is started.")
	driveCmd.Flags().Bool("wasm", false, "Compile to WebAssembly (with fallback to JavaScript). See https://flutter.dev/to/wasm for more information.")
	driveCmd.Flags().Bool("web-allow-expose-url", false, "Enables daemon-to-editor requests (app.exposeUrl) for exposing URLs when running on remote machines.")
	driveCmd.Flags().Int("web-browser-debug-port", 0, "The debug port the browser should use. If not specified, a random port is selected. Currently only Chrome supports this option. It serves the Chrome DevTools Protocol (https://chromedevtools.github.io/devtools-protocol/).")
	driveCmd.Flags().StringSlice("web-browser-flag", nil, "Additional flag to pass to a browser instance at startup. Chrome: https://www.chromium.org/developers/how-tos/run-chromium-with-flags/ Firefox: https://wiki.mozilla.org/Firefox/CommandLineOptions Multiple flags can be passed by repeating \"--web-browser-flag\" multiple times.")
	driveCmd.Flags().StringArray("web-define", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors only when compiling for the web. Multiple defines can be passed by repeating \"--web-define\".")
	driveCmd.Flags().StringArray("web-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors only when compiling for the web. Multiple defines can be passed by repeating \"--web-define\".")
	driveCmd.Flags().Bool("web-enable-expression-evaluation", false, "Enables expression evaluation in the debugger.")
	driveCmd.Flags().StringArray("web-header", nil, "Additional key-value pairs that will added by the web server as headers to all responses. Multiple headers can be passed by repeating \"--web-header\" multiple times.")
	driveCmd.Flags().String("web-hostname", "", "The hostname that the web server will use to resolve an IP to serve from. The unresolved hostname is used to launch Chrome when using the chrome Device. The name \"any\" may also be used to serve on any IPV4 for either the Chrome or web-server device.")
	driveCmd.Flags().String("web-launch-url", "", "The URL to provide to the browser. Defaults to an HTTP URL with the host name of \"--web-hostname\", the port of \"--web-port\", and the path set to \"/\".")
	driveCmd.Flags().Int("web-port", 0, "The host port to serve the web application from. If not provided, the tool will select a random open port on the host.")
	driveCmd.Flags().Bool("web-resources-cdn", false, "Use WebAssembly, CanvasKit, and other web resources from a content delivery network (CDN). Set to \"--no-web-resources-cdn\" to embed all web resources locally in the built app.")
	driveCmd.Flags().Bool("web-run-headless", false, "Launches the browser in headless mode. Currently only Chrome supports this option.")
	driveCmd.Flags().String("web-server-debug-backend-protocol", "", "The protocol (SSE or WebSockets) to use for the Dart Debug Extension backend service when using the Web Server device. Using WebSockets can improve performance but may fail when connecting through some proxy servers.")
	driveCmd.Flags().String("web-server-debug-injected-client-protocol", "", "The protocol (SSE or WebSockets) to use for the injected client when using the Web Server device. Using WebSockets can improve performance but may fail when connecting through some proxy servers.")
	driveCmd.Flags().String("web-server-debug-protocol", "", "The protocol (SSE or WebSockets) to use for the debug service proxy when using the Web Server device and Dart Debug extension. This is useful for editors/debug adapters that do not support debugging over SSE (the default protocol for Web Server/Dart Debugger extension).")
	driveCmd.Flags().String("web-tls-cert-key-path", "", "The certificate key that host will use to authenticate cert. If not provided, the tool will use default http scheme.")
	driveCmd.Flags().String("web-tls-cert-path", "", "The certificate that host will use to serve using TLS connection. If not provided, the tool will use default http scheme.")
	rootCmd.AddCommand(driveCmd)

	carapace.Gen(driveCmd).FlagCompletion(carapace.ActionMap{
		"browser-name": carapace.ActionValuesDescribed(
			"android-chrome", "Chrome on Android.",
			"chrome", "Google Chrome on this computer.",
			"edge", "Microsoft Edge on this computer (Windows only).",
			"firefox", "Mozilla Firefox on this computer.",
			"ios-safari", "Apple Safari on an iOS device.",
			"safari", "Apple Safari on this computer (macOS only).",
		),
		"chrome-binary": carapace.ActionFiles(),
		"device-connection": carapace.ActionValuesDescribed(
			"attached", "Only searches for devices connected by USB or built-in (such as simulators/emulators, MacOS/Windows, Chrome)",
			"wireless", "Only searches for devices connected wirelessly. Discovering wireless devices may take longer.",
			"both", "Searches for both attached and wireless devices.",
		),
		"device-user":                       adb.ActionUsers(),
		"driver":                            carapace.ActionFiles(),
		"profile-memory":                    carapace.ActionFiles(),
		"screenshot":                        carapace.ActionDirectories(),
		"target":                            carapace.ActionFiles(".dart"),
		"web-server-debug-backend-protocol": carapace.ActionValues("sse", "ws"),
		"web-server-debug-injected-client-protocol": carapace.ActionValues("sse", "ws"),
		"web-server-debug-protocol":                 carapace.ActionValues("sse", "ws"),
	})
}
