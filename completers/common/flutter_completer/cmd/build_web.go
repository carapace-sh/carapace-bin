package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_webCmd = &cobra.Command{
	Use:   "web",
	Short: "Build a web application bundle",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_webCmd).Standalone()

	build_webCmd.Flags().String("base-href", "", "Overrides the href attribute of the <base> tag in web/index.html. No change is done to web/index.html if this flag is not provided. The value has to start and end with a slash \"/\". For more information: https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base")
	build_webCmd.Flags().String("build-name", "", "A \"x.y.z\" string used as the version number shown to users. For each new version of your app, you will provide a version number to differentiate it.")
	build_webCmd.Flags().String("build-number", "", "An identifier used as an internal version number. Each build must have a unique identifier to differentiate it from previous builds.")
	build_webCmd.Flags().Bool("csp", false, "Disable dynamic generation of code in the generated output. This is necessary to satisfy CSP restrictions (see http://www.w3.org/TR/CSP/).")
	build_webCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_webCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	build_webCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_webCmd.Flags().String("dart2js-optimization", "", "Sets the optimization level used for Dart compilation to JavaScript. Deprecated: Please use \"-O=<level>\" / \"--optimization-level=<level>\".")
	build_webCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_webCmd.Flags().Bool("deprecated-js-interop", false, "Whether to allow the deprecated JS interop libraries (such as \"dart:html\", \"dart:js\", and \"dart:js_util\") when compiling to JavaScript. With \"--no-deprecated-js-interop\", importing these libraries is a compile-time error and conditional imports on them resolve to false. Use this to find code that still needs to migrate to \"package:web\" and \"dart:js_interop\". If not specified, the Dart compiler default is used. Has no effect when compiling to WebAssembly.")
	build_webCmd.Flags().Bool("dump-info", false, "Passes \"--dump-info\" to the Javascript compiler which generates information about the generated code in main.dart.js.info.json.")
	build_webCmd.Flags().StringSlice("enable-experiment", nil, "The name of an experimental Dart feature to enable. Multiple experiments can be enabled by repeating \"--enable-experiment\".")
	build_webCmd.Flags().Bool("enable-wasm-deferred-loading", false, "Enable multi-module deferred loading for Wasm.")
	build_webCmd.Flags().Bool("fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_webCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_webCmd.Flags().Bool("jit-release", false, "Build a JIT release version of your app.")
	build_webCmd.Flags().Bool("minify-js", false, "Generate minified output for js. If not explicitly set, uses the compilation mode (debug, profile, release).")
	build_webCmd.Flags().Bool("minify-wasm", false, "Generate minified output for wasm. If not explicitly set, uses the compilation mode (debug, profile, release).")
	build_webCmd.Flags().Bool("native-null-assertions", false, "Enables additional runtime null checks in web applications to ensure the correct nullability of native (such as in dart:html) and external (such as with JS interop) types. This is enabled by default but only takes effect in sound mode. To report an issue with a null assertion failure in dart:html or the other dart web libraries, please file a bug at: https://github.com/dart-lang/sdk/issues/labels/web-libraries")
	build_webCmd.Flags().Bool("no-deprecated-js-interop", false, "Whether to allow the deprecated JS interop libraries (such as \"dart:html\", \"dart:js\", and \"dart:js_util\") when compiling to JavaScript. With \"--no-deprecated-js-interop\", importing these libraries is a compile-time error and conditional imports on them resolve to false. Use this to find code that still needs to migrate to \"package:web\" and \"dart:js_interop\". If not specified, the Dart compiler default is used. Has no effect when compiling to WebAssembly.")
	build_webCmd.Flags().Bool("no-enable-wasm-deferred-loading", false, "Enable multi-module deferred loading for Wasm.")
	build_webCmd.Flags().Bool("no-fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_webCmd.Flags().Bool("no-frequency-based-minification", false, "Disables the frequency based minifier. Useful for comparing the output between builds.")
	build_webCmd.Flags().Bool("no-minify-js", false, "Generate minified output for js. If not explicitly set, uses the compilation mode (debug, profile, release).")
	build_webCmd.Flags().Bool("no-minify-wasm", false, "Generate minified output for wasm. If not explicitly set, uses the compilation mode (debug, profile, release).")
	build_webCmd.Flags().Bool("no-native-null-assertions", false, "Enables additional runtime null checks in web applications to ensure the correct nullability of native (such as in dart:html) and external (such as with JS interop) types. This is enabled by default but only takes effect in sound mode. To report an issue with a null assertion failure in dart:html or the other dart web libraries, please file a bug at: https://github.com/dart-lang/sdk/issues/labels/web-libraries")
	build_webCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_webCmd.Flags().Bool("no-source-maps", false, "Generate a sourcemap file. These can be used by browsers to view and debug the original source code of a compiled and minified Dart application.")
	build_webCmd.Flags().Bool("no-strip-wasm", false, "Whether to strip the resulting wasm file of static symbol names.")
	build_webCmd.Flags().Bool("no-tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	build_webCmd.Flags().Bool("no-wasm-dry-run", false, "Compiles wasm in dry run mode during JS only compilations. Disable to suppress warnings.")
	build_webCmd.Flags().Bool("no-web-content-hash", false, "Include a content hash in the filenames of compiled web entrypoints (for example, \"main.dart.<hash>.js\") and static assets in \"assets/\", and emit a \"precache_manifest.json\" file so that browsers and custom service workers fetch new versions after a deploy instead of serving stale cached files. The web server must still serve \"index.html\" and \"flutter_bootstrap.js\" with revalidation (for example, \"Cache-Control: no-cache\") for a new deploy to be picked up. When \"--no-web-resources-cdn\" is used, locally bundled \"canvaskit/**\" files are not content-hashed (\"urlHashed: false\" in \"precache_manifest.json\") and must not be served with \"immutable\" caching. Not supported with deferred imports. Custom \"index.html\" files that reference \"main.dart.js\" directly, and the deprecated \"FlutterLoader.loadEntrypoint\" JavaScript API, are incompatible with this flag.")
	build_webCmd.Flags().Bool("no-web-resources-cdn", false, "Use WebAssembly, CanvasKit, and other web resources from a content delivery network (CDN). Set to \"--no-web-resources-cdn\" to embed all web resources locally in the built app.")
	build_webCmd.Flags().StringP("optimization-level", "O", "", "Sets the optimization level used for Dart compilation to JavaScript/Wasm.")
	build_webCmd.Flags().StringP("output", "o", "", "The absolute path to the directory where the repository is generated. By default, this is <current-directory>/build/<target-platform>. Currently supported for subcommands: aar, web.")
	build_webCmd.Flags().String("output-dir", "", "The absolute path to the directory where the repository is generated. By default, this is <current-directory>/build/<target-platform>. Currently supported for subcommands: aar, web.")
	build_webCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_webCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_webCmd.Flags().String("pwa-strategy", "", "This option is deprecated and will be removed in a future Flutter release. The caching strategy to be used by the PWA service worker.")
	build_webCmd.Flags().Bool("release", false, "Build a release version of your app.")
	build_webCmd.Flags().Bool("source-maps", false, "Generate a sourcemap file. These can be used by browsers to view and debug the original source code of a compiled and minified Dart application.")
	build_webCmd.Flags().String("static-assets-url", "", "Used when serving the static assets from a different domain the application is hosted on. The value has to end with a slash \"/\". When this is set, it will replace all $FLUTTER_STATIC_ASSETS_URL in web/index.html for the given value.")
	build_webCmd.Flags().Bool("strip-wasm", false, "Whether to strip the resulting wasm file of static symbol names.")
	build_webCmd.Flags().StringP("target", "t", "", "The main entrypoint file of the application, as run on the device. If the \"--target\" option is omitted, but a file name is provided on the command line, then that file is used instead.")
	build_webCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application are included.")
	build_webCmd.Flags().Bool("wasm", false, "Compile to WebAssembly (with fallback to JavaScript). See https://flutter.dev/to/wasm for more information.")
	build_webCmd.Flags().Bool("wasm-dry-run", false, "Compiles wasm in dry run mode during JS only compilations. Disable to suppress warnings.")
	build_webCmd.Flags().Bool("web-content-hash", false, "Include a content hash in the filenames of compiled web entrypoints (for example, \"main.dart.<hash>.js\") and static assets in \"assets/\", and emit a \"precache_manifest.json\" file so that browsers and custom service workers fetch new versions after a deploy instead of serving stale cached files. The web server must still serve \"index.html\" and \"flutter_bootstrap.js\" with revalidation (for example, \"Cache-Control: no-cache\") for a new deploy to be picked up. When \"--no-web-resources-cdn\" is used, locally bundled \"canvaskit/**\" files are not content-hashed (\"urlHashed: false\" in \"precache_manifest.json\") and must not be served with \"immutable\" caching. Not supported with deferred imports. Custom \"index.html\" files that reference \"main.dart.js\" directly, and the deprecated \"FlutterLoader.loadEntrypoint\" JavaScript API, are incompatible with this flag.")
	build_webCmd.Flags().StringArray("web-define", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors only when compiling for the web. Multiple defines can be passed by repeating \"--web-define\".")
	build_webCmd.Flags().StringSlice("web-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables only when compiling for the web. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--web-define-from-file\".")
	build_webCmd.Flags().StringArray("web-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors only when compiling for the web. Multiple defines can be passed by repeating \"--web-define\".")
	build_webCmd.Flags().Bool("web-resources-cdn", false, "Use WebAssembly, CanvasKit, and other web resources from a content delivery network (CDN). Set to \"--no-web-resources-cdn\" to embed all web resources locally in the built app.")
	buildCmd.AddCommand(build_webCmd)

	carapace.Gen(build_webCmd).FlagCompletion(carapace.ActionMap{
		"dart2js-optimization": carapace.ActionValues("O1", "O2", "O3", "O4"),
		"optimization-level":   carapace.ActionValues("0", "1", "2", "3", "4"),
		"output":               carapace.ActionDirectories(),
		"output-dir":           carapace.ActionDirectories(),
		"pwa-strategy": carapace.ActionValuesDescribed(
			"offline-first", "Attempt to cache the application shell eagerly and then lazily cache all subsequent assets as they are loaded. When making a network request for an asset, the offline cache will be preferred.",
			"none", "Generate a service worker with no body. This is useful for local testing or in cases where the service worker caching functionality is not desirable",
		),
		"target": carapace.ActionFiles(".dart"),
	})

}
