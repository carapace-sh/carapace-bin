package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddCompilerFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("D", "D", "", "Marks a conditional compilation flag as true")
	cmd.Flags().StringS("F", "F", "", "Add directory to framework search path")
	cmd.Flags().StringS("Fsystem", "Fsystem", "", "Add directory to system framework search path")
	cmd.Flags().StringS("I", "I", "", "Add directory to the import search path")
	cmd.Flags().StringS("Isystem", "Isystem", "", "Add directory to the system import search path")
	cmd.Flags().StringS("L", "L", "", "Add directory to library link search path")
	cmd.Flags().BoolS("O", "O", false, "Compile with optimizations")
	cmd.Flags().BoolS("Onone", "Onone", false, "Compile without any optimization")
	cmd.Flags().BoolS("Osize", "Osize", false, "Compile with optimizations and target small code size")
	cmd.Flags().BoolS("Ounchecked", "Ounchecked", false, "Compile with optimizations and remove runtime safety checks")
	cmd.Flags().BoolS("Rcache-compile-job", "Rcache-compile-job", false, "Show remarks for compiler caching")
	cmd.Flags().BoolS("Rcross-import", "Rcross-import", false, "Emit a remark if a cross-import of a module is triggered.")
	cmd.Flags().BoolS("Rindexing-system-module", "Rindexing-system-module", false, "Emit a remark when indexing a system module")
	cmd.Flags().BoolS("Rmacro-loading", "Rmacro-loading", false, "Emit remarks about loaded macro implementations")
	cmd.Flags().BoolS("Rmodule-api-import", "Rmodule-api-import", false, "Emit remarks about the import bridging in each element composing the API")
	cmd.Flags().BoolS("Rmodule-loading", "Rmodule-loading", false, "Emit remarks about loaded module")
	cmd.Flags().BoolS("Rmodule-recovery", "Rmodule-recovery", false, "Emit remarks about contextual inconsistencies in loaded modules")
	cmd.Flags().BoolS("Rmodule-serialization", "Rmodule-serialization", false, "Emit remarks about module serialization")
	cmd.Flags().StringS("Rpass", "Rpass", "", "Report performed transformations by optimization passes whose name matches the given POSIX regular expression")
	cmd.Flags().StringS("Rpass-missed", "Rpass-missed", "", "Report missed transformations by optimization passes whose name matches the given POSIX regular expression")
	cmd.Flags().BoolS("Rskip-explicit-interface-build", "Rskip-explicit-interface-build", false, "Emit a remark if an explicit module interface invocation has an early exit because the expected output is up-to-date")
	cmd.Flags().StringS("Xcc", "Xcc", "", "Pass <arg> to the C/C++/Objective-C compiler")
	cmd.Flags().StringS("Xlinker", "Xlinker", "", "Specifies an option which should be passed to the linker")
	cmd.Flags().StringS("access-notes-path", "access-notes-path", "", "Specify YAML file to override attributes on Swift declarations in this module")
	cmd.Flags().StringS("allow-availability-platforms", "allow-availability-platforms", "", "Restrict availability metadata to the given platforms, e.g. 'macOS,Swift'")
	cmd.Flags().BoolS("allow-non-resilient-access", "allow-non-resilient-access", false, "Ensures all contents are generated besides exportable decls in the binary module, so non-resilient access can be allowed")
	cmd.Flags().StringS("allowable-client", "allowable-client", "", "Module names that are allowed to import this module")
	cmd.Flags().StringS("assert-config", "assert-config", "", "Specify the assert_configuration replacement. Possible values are Debug, Release, Unchecked, DisableReplacement.")
	cmd.Flags().StringS("block-availability-platforms", "block-availability-platforms", "", "Remove the given platforms from symbol graph availability metadata, e.g. 'macOS,Swift'")
	cmd.Flags().StringS("build-id", "build-id", "", "Specify the build ID argument passed to the linker")
	cmd.Flags().BoolS("cache-compile-job", "cache-compile-job", false, "Enable compiler caching")
	cmd.Flags().BoolS("cache-disable-replay", "cache-disable-replay", false, "Skip loading the compilation result from cache")
	cmd.Flags().StringS("cas-path", "cas-path", "", "Path to CAS")
	cmd.Flags().StringS("cas-plugin-option", "cas-plugin-option", "", "Option pass to CAS Plugin")
	cmd.Flags().StringS("cas-plugin-path", "cas-plugin-path", "", "Path to CAS Plugin")
	cmd.Flags().StringS("clang-build-session-file", "clang-build-session-file", "", "Use the last modification time of <file> as the underlying Clang build session timestamp")
	cmd.Flags().StringS("clang-scanner-module-cache-path", "clang-scanner-module-cache-path", "", "Specifies the Clang dependency scanner module cache path")
	cmd.Flags().StringS("clang-target", "clang-target", "", "Separately set the target we should use for internal Clang instance")
	cmd.Flags().StringS("clang-target-variant", "clang-target-variant", "", "Separately set the target we should use for internal Clang instance for the 'zippered' code for macCatalyst")
	cmd.Flags().BoolS("color-diagnostics", "color-diagnostics", false, "Print diagnostics in color")
	cmd.Flags().BoolS("compiler-assertions", "compiler-assertions", false, "Enable internal self-checks while compiling")
	cmd.Flags().BoolS("continue-building-after-errors", "continue-building-after-errors", false, "Continue building, even after errors are encountered")
	cmd.Flags().StringS("coverage-prefix-map", "coverage-prefix-map", "", "Remap source paths in coverage info")
	cmd.Flags().StringS("cxx-interoperability-mode", "cxx-interoperability-mode", "", "Enables C++ interoperability; pass 'default' to enable or 'off' to disable")
	cmd.Flags().StringS("debug-info-format", "debug-info-format", "", "Specify the debug info format type to either 'dwarf' or 'codeview'")
	cmd.Flags().BoolS("debug-info-store-invocation", "debug-info-store-invocation", false, "Emit the compiler invocation in the debug info.")
	cmd.Flags().StringS("debug-module-path", "debug-module-path", "", "Path to this module's binary swiftmodule artifact (required by debug info)")
	cmd.Flags().StringS("debug-prefix-map", "debug-prefix-map", "", "Remap source paths in debug info")
	cmd.Flags().StringS("default-isolation", "default-isolation", "", "Specify the default actor isolation: MainActor or nonisolated. Defaults to nonisolated.")
	cmd.Flags().StringS("dependency-scan-serialize-diagnostics-path", "dependency-scan-serialize-diagnostics-path", "", "Emit a serialized diagnostics file for the dependency scanning task to <path>")
	cmd.Flags().StringS("diagnostic-style", "diagnostic-style", "", "The formatting style used when printing diagnostics ('swift' or 'llvm')")
	cmd.Flags().BoolS("disable-actor-data-race-checks", "disable-actor-data-race-checks", false, "Disable runtime checks for actor data races")
	cmd.Flags().BoolS("disable-autolinking-runtime-compatibility", "disable-autolinking-runtime-compatibility", false, "Do not use autolinking for runtime compatibility libraries")
	cmd.Flags().BoolS("disable-autolinking-runtime-compatibility-concurrency", "disable-autolinking-runtime-compatibility-concurrency", false, "Do not use autolinking for the concurrency runtime compatibility library")
	cmd.Flags().BoolS("disable-autolinking-runtime-compatibility-dynamic-replacements", "disable-autolinking-runtime-compatibility-dynamic-replacements", false, "Do not use autolinking for the dynamic replacement runtime compatibility library")
	cmd.Flags().BoolS("disable-clang-target", "disable-clang-target", false, "Disable a separately specified target triple for Clang instance to use")
	cmd.Flags().BoolS("disable-dynamic-actor-isolation", "disable-dynamic-actor-isolation", false, "Disable dynamic actor isolation checks")
	cmd.Flags().StringS("disable-experimental-feature", "disable-experimental-feature", "", "Disable an experimental feature")
	cmd.Flags().BoolS("disable-incremental-file-hashing", "disable-incremental-file-hashing", false, "Disable hashing of input and dependency file data that can prevent unnecessary invalidation")
	cmd.Flags().BoolS("disable-incremental-imports", "disable-incremental-imports", false, "Disable cross-module incremental build metadata and driver scheduling for Swift modules")
	cmd.Flags().BoolS("disable-only-one-dependency-file", "disable-only-one-dependency-file", false, "Disables incremental build optimization that only produces one dependencies file")
	cmd.Flags().BoolS("disable-sandbox", "disable-sandbox", false, "Disable using the sandbox when executing subprocesses")
	cmd.Flags().StringS("disable-upcoming-feature", "disable-upcoming-feature", "", "Disable a feature that will be introduced in an upcoming language version")
	cmd.Flags().BoolS("disallow-use-new-driver", "disallow-use-new-driver", false, "Disable using new swift-driver")
	cmd.Flags().StringS("dwarf-version", "dwarf-version", "", "DWARF debug info version to produce if requested")
	cmd.Flags().StringS("e", "e", "", "Executes a line of code provided on the command line")
	cmd.Flags().StringS("embed-tbd-for-module", "embed-tbd-for-module", "", "Embed symbols from the module in the emitted tbd file")
	cmd.Flags().StringS("emit-module-dependencies-path", "emit-module-dependencies-path", "", "Emit a discovered dependencies file for the emit-module task to <path>")
	cmd.Flags().StringS("emit-module-serialize-diagnostics-path", "emit-module-serialize-diagnostics-path", "", "Emit a serialized diagnostics file for the emit-module task to <path>")
	cmd.Flags().BoolS("enable-actor-data-race-checks", "enable-actor-data-race-checks", false, "Emit runtime checks for actor data races")
	cmd.Flags().BoolS("enable-autolinking-runtime-compatibility-bytecode-layouts", "enable-autolinking-runtime-compatibility-bytecode-layouts", false, "Enable autolinking for the bytecode layouts runtime compatibility library")
	cmd.Flags().BoolS("enable-bare-slash-regex", "enable-bare-slash-regex", false, "Enable the use of forward slash regular-expression literal syntax")
	cmd.Flags().BoolS("enable-builtin-module", "enable-builtin-module", false, "Enables the explicit import of the Builtin module")
	cmd.Flags().BoolS("enable-deterministic-check", "enable-deterministic-check", false, "Check compiler output determinism by running it twice")
	cmd.Flags().BoolS("enable-experimental-additive-arithmetic-derivation", "enable-experimental-additive-arithmetic-derivation", false, "Enable experimental 'AdditiveArithmetic' derived conformances")
	cmd.Flags().BoolS("enable-experimental-concise-pound-file", "enable-experimental-concise-pound-file", false, "Enable experimental concise '#file' identifier")
	cmd.Flags().StringS("enable-experimental-feature", "enable-experimental-feature", "", "Enable an experimental feature")
	cmd.Flags().BoolS("enable-experimental-forward-mode-differentiation", "enable-experimental-forward-mode-differentiation", false, "Enable experimental forward mode differentiation")
	cmd.Flags().BoolS("enable-incremental-file-hashing", "enable-incremental-file-hashing", false, "Enable hashing of input and dependency file data that can prevent unnecessary invalidation")
	cmd.Flags().BoolS("enable-incremental-imports", "enable-incremental-imports", false, "Enable cross-module incremental build metadata and driver scheduling for Swift modules")
	cmd.Flags().BoolS("enable-library-evolution", "enable-library-evolution", false, "Build the module to allow binary-compatible library evolution")
	cmd.Flags().BoolS("enable-only-one-dependency-file", "enable-only-one-dependency-file", false, "Enables incremental build optimization that only produces one dependencies file")
	cmd.Flags().StringS("enable-upcoming-feature", "enable-upcoming-feature", "", "Enable a feature that will be introduced in an upcoming language version")
	cmd.Flags().StringS("enforce-exclusivity", "enforce-exclusivity", "", "Enforce law of exclusivity")
	cmd.Flags().BoolS("experimental-allow-non-resilient-access", "experimental-allow-non-resilient-access", false, "Deprecated; use -allow-non-resilient-access instead")
	cmd.Flags().BoolS("experimental-emit-variant-module", "experimental-emit-variant-module", false, "When a target variant triple is specified, the same driver invocation will emit two Swift modules, one for the primary target and one for the variant.")
	cmd.Flags().BoolS("experimental-package-bypass-resilience", "experimental-package-bypass-resilience", false, "Deprecated; has no effect")
	cmd.Flags().BoolS("experimental-package-cmo", "experimental-package-cmo", false, "Deprecated; use -package-cmo instead")
	cmd.Flags().BoolS("experimental-package-cmo-abort-on-deserialization-fail", "experimental-package-cmo-abort-on-deserialization-fail", false, "Abort if a deserialization error is found while package optimization is enabled")
	cmd.Flags().StringS("explain-module-dependency", "explain-module-dependency", "", "Emit remark describing why compilation may depend on a module with a given name.")
	cmd.Flags().StringS("explain-module-dependency-detailed", "explain-module-dependency-detailed", "", "Emit remarks describing every possible dependency path that explains why compilation may depend on a module with a given name.")
	cmd.Flags().BoolS("explicit-auto-linking", "explicit-auto-linking", false, "Instead of linker-load directives, have the driver specify all link dependencies on the linker invocation. Requires '-explicit-module-build'.")
	cmd.Flags().StringS("export-as", "export-as", "", "Module name to use when referenced in clients module interfaces")
	cmd.Flags().StringS("external-plugin-path", "external-plugin-path", "", "#<plugin-server-path> Add directory to the plugin search path with a plugin server executable")
	cmd.Flags().StringS("file-compilation-dir", "file-compilation-dir", "", "The compilation directory to embed in the debug info. Coverage mapping is not supported yet.")
	cmd.Flags().StringS("file-prefix-map", "file-prefix-map", "", "Remap source paths in debug, coverage, and index info")
	cmd.Flags().StringS("framework", "framework", "", "Specifies a framework which should be linked against")
	cmd.Flags().BoolS("g", "g", false, "Emit debug info. This is the preferred setting for debugging with LLDB.")
	cmd.Flags().BoolS("gdwarf-types", "gdwarf-types", false, "Emit full DWARF type info.")
	cmd.Flags().BoolS("gline-tables-only", "gline-tables-only", false, "Emit minimal debug info for backtraces only")
	cmd.Flags().BoolS("gnone", "gnone", false, "Don't emit debug info")
	cmd.Flags().StringS("import-bridging-header", "import-bridging-header", "", "Implicitly imports a C header file")
	cmd.Flags().StringS("in-process-plugin-server-path", "in-process-plugin-server-path", "", "Path to dynamic library plugin server")
	cmd.Flags().BoolS("index-ignore-clang-modules", "index-ignore-clang-modules", false, "Avoid indexing clang modules (pcms)")
	cmd.Flags().BoolS("index-include-locals", "index-include-locals", false, "Include local definitions/references in the produced index data.")
	cmd.Flags().BoolS("index-store-compress", "index-store-compress", false, "Compress the unit and record files in the index store")
	cmd.Flags().StringS("index-store-path", "index-store-path", "", "Store indexing data to <path>")
	cmd.Flags().StringS("index-unit-output-path", "index-unit-output-path", "", "Use <path> as the output path in the produced index data.")
	cmd.Flags().StringS("internal-import-bridging-header", "internal-import-bridging-header", "", "Implicitly imports a C header file as an internal import")
	cmd.Flags().StringS("ir-output-dir", "ir-output-dir", "", "Output LLVM IR files to directory <dir> as additional output during compilation")
	cmd.Flags().IntS("j", "j", 0, "Number of commands to execute in parallel")
	cmd.Flags().StringS("l", "l", "", "Specifies a library which should be linked against")
	cmd.Flags().StringS("libc", "libc", "", "libc runtime library to use")
	cmd.Flags().BoolS("link-objc-runtime", "link-objc-runtime", false, "Deprecated")
	cmd.Flags().StringS("load-pass-plugin", "load-pass-plugin", "", "Load LLVM pass plugin from a dynamic shared object file.")
	cmd.Flags().StringS("load-plugin-executable", "load-plugin-executable", "", "#<module-names> Path to a compiler plugin executable and a comma-separated list of module names where the macro types are declared")
	cmd.Flags().StringS("load-plugin-library", "load-plugin-library", "", "Path to a dynamic library containing compiler plugins such as macros")
	cmd.Flags().StringS("load-resolved-plugin", "load-resolved-plugin", "", "#<executable-path>#<module-names> Path to resolved plugin configuration and a comma-separated list of module names where the macro types are declared. Library path and executable path can be empty if not used")
	cmd.Flags().StringS("locale", "locale", "", "Choose a language for diagnostic messages")
	cmd.Flags().StringS("localization-path", "localization-path", "", "Path to localized diagnostic messages directory")
	cmd.Flags().StringS("min-swift-runtime-version", "min-swift-runtime-version", "", "The minimum Swift runtime version that will be available at runtime")
	cmd.Flags().StringS("module-abi-name", "module-abi-name", "", "ABI name to use for the contents of this module")
	cmd.Flags().StringS("module-alias", "module-alias", "", "If a source file imports or references module <alias_name>, the <real_name> is used for the contents of the file")
	cmd.Flags().StringS("module-cache-path", "module-cache-path", "", "Specifies the module cache path")
	cmd.Flags().StringS("module-link-name", "module-link-name", "", "Library to link against when using this module")
	cmd.Flags().StringS("module-name", "module-name", "", "Name of the module to build")
	cmd.Flags().BoolS("no-color-diagnostics", "no-color-diagnostics", false, "Do not print diagnostics in color")
	cmd.Flags().BoolS("no-warnings-as-errors", "no-warnings-as-errors", false, "Treat warnings as warnings")
	cmd.Flags().BoolS("nostdimport", "nostdimport", false, "Don't search the standard library or toolchain import paths for modules")
	cmd.Flags().BoolS("nostdlibimport", "nostdlibimport", false, "Don't search the standard library import path for modules")
	cmd.Flags().IntS("num-threads", "num-threads", 0, "Enable multi-threading and specify number of threads")
	cmd.Flags().BoolS("package-cmo", "package-cmo", false, "Enable optimization to perform default CMO within a package boundary")
	cmd.Flags().StringS("package-name", "package-name", "", "Name of the package the module belongs to")
	cmd.Flags().StringS("plugin-path", "plugin-path", "", "Add directory to the plugin search path")
	cmd.Flags().BoolS("prefix-serialized-debugging-options", "prefix-serialized-debugging-options", false, "Apply debug prefix mappings to serialized debug info in Swiftmodule files")
	cmd.Flags().BoolS("pretty-print", "pretty-print", false, "Pretty-print the output JSON")
	cmd.Flags().BoolS("print-educational-notes", "print-educational-notes", false, "Include educational notes in printed diagnostic output, if available")
	cmd.Flags().BoolS("print-static-build-config", "print-static-build-config", false, "Print static build configuration that can be used to evaluate #ifs in Swift source code")
	cmd.Flags().BoolS("print-supported-features", "print-supported-features", false, "Print information about features supported by the compiler")
	cmd.Flags().BoolS("print-target-info", "print-target-info", false, "Print target information for the given target <triple>, such as x86_64-apple-macos10.9")
	cmd.Flags().StringS("project-name", "project-name", "", "Name of the project this module to build belongs to")
	cmd.Flags().StringS("public-module-name", "public-module-name", "", "Public facing module name to use in diagnostics and documentation")
	cmd.Flags().StringS("register-module-dependency", "register-module-dependency", "", "Register module for dependency scan without importing in the frontend")
	cmd.Flags().BoolS("remove-runtime-asserts", "remove-runtime-asserts", false, "Remove runtime safety checks.")
	cmd.Flags().StringS("runtime-compatibility-version", "runtime-compatibility-version", "", "Link compatibility library for Swift runtime version, or 'none'")
	cmd.Flags().StringS("save-optimization-record", "save-optimization-record", "", "Generate an optimization record file in a specific format (default: YAML)")
	cmd.Flags().StringS("save-optimization-record-passes", "save-optimization-record-passes", "", "Only include passes which match a specified regular expression in the generated optimization record (by default, include all passes)")
	cmd.Flags().StringS("save-optimization-record-path", "save-optimization-record-path", "", "Specify the file name of any generated optimization record")
	cmd.Flags().StringS("scanner-prefix-map", "scanner-prefix-map", "", "Remap paths reported by dependency scanner")
	cmd.Flags().StringS("scanner-prefix-map-paths", "scanner-prefix-map-paths", "", "Remap paths reported by dependency scanner")
	cmd.Flags().StringS("scanner-prefix-map-sdk", "scanner-prefix-map-sdk", "", "Remap paths within SDK reported by dependency scanner")
	cmd.Flags().StringS("scanner-prefix-map-toolchain", "scanner-prefix-map-toolchain", "", "Remap paths within toolchain directory reported by dependency scanner")
	cmd.Flags().StringS("sdk", "sdk", "", "Compile against <sdk>")
	cmd.Flags().StringS("sdk-module-cache-path", "sdk-module-cache-path", "", "Specifies the module cache path for explicitly-built SDK modules")
	cmd.Flags().StringS("serialize-diagnostics-path", "serialize-diagnostics-path", "", "Emit a serialized diagnostics file to <path>")
	cmd.Flags().StringS("sil-output-dir", "sil-output-dir", "", "Output SIL files to directory <dir> as additional output during compilation")
	cmd.Flags().BoolS("static-executable", "static-executable", false, "Statically link the executable")
	cmd.Flags().BoolS("static-stdlib", "static-stdlib", false, "Statically link the Swift standard library")
	cmd.Flags().StringS("strict-concurrency", "strict-concurrency", "", "Specify the how strict concurrency checking will be. The value may be 'minimal' (most 'Sendable' checking is disabled), 'targeted' ('Sendable' checking is enabled in code that uses the concurrency model, or 'complete' ('Sendable' and other checking is enabled for all code in the module)")
	cmd.Flags().BoolS("strict-memory-safety", "strict-memory-safety", false, "Enable strict memory safety checking")
	cmd.Flags().BoolS("strict-memory-safety:migrate", "strict-memory-safety:migrate", false, "Enable migration to strict memory safety checking")
	cmd.Flags().BoolS("suppress-notes", "suppress-notes", false, "Suppress all notes")
	cmd.Flags().BoolS("suppress-remarks", "suppress-remarks", false, "Suppress all remarks")
	cmd.Flags().BoolS("suppress-warnings", "suppress-warnings", false, "Suppress all warnings")
	cmd.Flags().StringS("swift-version", "swift-version", "", "Interpret input according to a specific Swift language version number")
	cmd.Flags().StringS("sysroot", "sysroot", "", "Native Platform sysroot")
	cmd.Flags().StringS("target", "target", "", "Generate code for the given target <triple>, such as x86_64-apple-macos10.9")
	cmd.Flags().StringS("target-cpu", "target-cpu", "", "Generate code for a particular CPU variant")
	cmd.Flags().StringS("target-min-inlining-version", "target-min-inlining-version", "", "Require inlinable code with no '@available' attribute to back-deploy to this version of the '-target' OS")
	cmd.Flags().StringS("target-variant", "target-variant", "", "Generate 'zippered' code for macCatalyst that can run on the specified variant target triple in addition to the main -target triple")
	cmd.Flags().StringS("use-ld", "use-ld", "", "Specifies the flavor of the linker to be used")
	cmd.Flags().StringS("user-module-version", "user-module-version", "", "Module version specified from Swift module authors")
	cmd.Flags().BoolS("v", "v", false, "Show commands to run and use verbose output")
	cmd.Flags().BoolS("validate-clang-modules-once", "validate-clang-modules-once", false, "Don't verify input files for Clang modules if the module has been successfully validated or loaded during this build session")
	cmd.Flags().StringS("vfsoverlay", "vfsoverlay", "", "Add directory to VFS overlay file")
	cmd.Flags().StringS("visualc-tools-root", "visualc-tools-root", "", "VisualC++ Tools Root")
	cmd.Flags().StringS("visualc-tools-version", "visualc-tools-version", "", "VisualC++ ToolSet Version")
	cmd.Flags().BoolS("warn-concurrency", "warn-concurrency", false, "Warn about code that is unsafe according to the Swift Concurrency model and will become ill-formed in a future language version")
	cmd.Flags().BoolS("warn-implicit-overrides", "warn-implicit-overrides", false, "Warn about implicit overrides of protocol members")
	cmd.Flags().BoolS("warnings-as-errors", "warnings-as-errors", false, "Treat warnings as errors")
	cmd.Flags().StringS("windows-sdk-root", "windows-sdk-root", "", "Windows SDK Root")
	cmd.Flags().StringS("windows-sdk-version", "windows-sdk-version", "", "Windows SDK Version")
	cmd.Flags().StringS("working-directory", "working-directory", "", "Resolve file paths relative to the specified directory")
}

func CompilerFlagCompletion(cmd *cobra.Command) {
	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"access-notes-path":         carapace.ActionFiles(".yaml"),
		"assert-config":             carapace.ActionValues("Debug", "Release", "Unchecked", "DisableReplacement"),
		"cas-path":                  carapace.ActionDirectories(),
		"cas-plugin-path":           carapace.ActionFiles(),
		"cxx-interoperability-mode": carapace.ActionValues("default", "off"),
		"debug-info-format":         carapace.ActionValues("dwarf", "codeview"),
		"debug-module-path":         carapace.ActionFiles(),
		"default-isolation":         carapace.ActionValues("MainActor", "nonisolated"),
		"dependency-scan-serialize-diagnostics-path": carapace.ActionFiles(),
		"diagnostic-style":                           carapace.ActionValues("swift", "llvm"),
		"emit-module-dependencies-path":              carapace.ActionFiles(),
		"emit-module-serialize-diagnostics-path":     carapace.ActionFiles(),
		"enforce-exclusivity":                        carapace.ActionValues("always", "runtime-only", "none"),
		"external-plugin-path":                       carapace.ActionFiles(),
		"file-compilation-dir":                       carapace.ActionDirectories(),
		"import-bridging-header":                     carapace.ActionFiles(".h"),
		"in-process-plugin-server-path":              carapace.ActionFiles(),
		"index-store-path":                           carapace.ActionDirectories(),
		"index-unit-output-path":                     carapace.ActionFiles(),
		"internal-import-bridging-header":            carapace.ActionFiles(".h"),
		"ir-output-dir":                              carapace.ActionDirectories(),
		"load-pass-plugin":                           carapace.ActionFiles(),
		"load-plugin-executable":                     carapace.ActionFiles(),
		"load-plugin-library":                        carapace.ActionFiles(),
		"load-resolved-plugin":                       carapace.ActionFiles(),
		"localization-path":                          carapace.ActionDirectories(),
		"module-cache-path":                          carapace.ActionDirectories(),
		"plugin-path":                                carapace.ActionDirectories(),
		"runtime-compatibility-version":              carapace.ActionValues("none"),
		"save-optimization-record":                   carapace.ActionValues("yaml", "bitstream", "json"),
		"save-optimization-record-passes":            carapace.ActionValues(),
		"save-optimization-record-path":              carapace.ActionFiles(),
		"scanner-prefix-map-paths":                   carapace.ActionValues(),
		"scanner-prefix-map-sdk":                     carapace.ActionDirectories(),
		"scanner-prefix-map-toolchain":               carapace.ActionDirectories(),
		"sdk":                                        carapace.ActionDirectories(),
		"sdk-module-cache-path":                      carapace.ActionDirectories(),
		"serialize-diagnostics-path":                 carapace.ActionFiles(),
		"sil-output-dir":                             carapace.ActionDirectories(),
		"strict-concurrency":                         carapace.ActionValues("minimal", "targeted", "complete"),
		"swift-version":                              carapace.ActionValues("4", "4.2", "5", "6"),
		"sysroot":                                    carapace.ActionDirectories(),
		"use-ld":                                     carapace.ActionValues("lld", "gold", "bfd"),
		"vfsoverlay":                                 carapace.ActionFiles(),
		"visualc-tools-root":                         carapace.ActionDirectories(),
		"windows-sdk-root":                           carapace.ActionDirectories(),
		"working-directory":                          carapace.ActionDirectories(),
	})
}
