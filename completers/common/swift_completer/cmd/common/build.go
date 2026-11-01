package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddPackageFlags(cmd *cobra.Command) {
	cmd.Flags().String("cache-path", "", "Specify the shared cache directory path")
	cmd.Flags().String("config-path", "", "Specify the shared configuration directory path")
	cmd.Flags().String("package-path", "", "Specify the package path to operate on")
	cmd.Flags().StringArray("pkg-config-path", nil, "Specify alternative path to search for pkg-config .pc files")
	cmd.Flags().String("scratch-path", "", "Specify a custom scratch directory path (default: .build)")
	cmd.Flags().String("security-path", "", "Specify the shared security directory path")
	cmd.Flags().String("swift-sdks-path", "", "Path to the directory containing installed Swift SDKs")
	cmd.Flags().StringArray("toolset", nil, "Specify a toolset JSON file to use when building")

	cmd.Flags().Bool("disable-build-manifest-caching", false, "Disable build manifest caching")
	cmd.Flags().Bool("disable-dependency-cache", false, "Disable shared cache when fetching dependencies")
	cmd.Flags().Bool("disable-experimental-prebuilts", false, "Disable prebuilt swift-syntax libraries for macros")
	cmd.Flags().Bool("enable-build-manifest-caching", false, "Enable build manifest caching")
	cmd.Flags().Bool("enable-dependency-cache", false, "Use a shared cache when fetching dependencies")
	cmd.Flags().Bool("enable-experimental-prebuilts", false, "Use prebuilt swift-syntax libraries for macros")
	cmd.Flags().String("manifest-cache", "", "Caching mode of Package.swift manifests (shared, local, none)")

	cmd.Flags().Bool("color-diagnostics", false, "Enable color diagnostics when printing to a TTY")
	cmd.Flags().Bool("no-color-diagnostics", false, "Disable color diagnostics when printing to a TTY")
	cmd.Flags().BoolP("quiet", "q", false, "Decrease verbosity to only include error output")
	cmd.Flags().BoolP("verbose", "v", false, "Increase verbosity to include informational output")
	cmd.Flags().BoolN("very-verbose", "vv", false, "Increase verbosity to include debug output")

	cmd.Flags().Bool("disable-netrc", false, "Disable netrc file for authentication")
	cmd.Flags().Bool("disable-sandbox", false, "Disable using the sandbox when executing subprocesses")
	cmd.Flags().Bool("disable-signature-validation", false, "Do not validate signatures on package releases downloaded from the registry")
	cmd.Flags().Bool("enable-netrc", false, "Use netrc file for authentication")
	cmd.Flags().Bool("enable-signature-validation", false, "Validate signatures on package releases downloaded from the registry")
	cmd.Flags().Bool("netrc", false, "Use netrc file for authentication")
	cmd.Flags().String("netrc-file", "", "Specify the netrc file path")
	cmd.Flags().String("resolver-fingerprint-checking", "", "Fingerprint checking mode of the resolver (default: strict)")
	cmd.Flags().String("resolver-signing-entity-checking", "", "Signing entity checking mode of the resolver (default: warn)")

	cmd.Flags().String("default-registry-url", "", "Default registry URL")
	cmd.Flags().Bool("disable-automatic-resolution", false, "Only use versions from the Package.resolved file")
	cmd.Flags().Bool("disable-prefetching", false, "Disable prefetching")
	cmd.Flags().Bool("disable-scm-to-registry-transformation", false, "Disable source control to registry transformation")
	cmd.Flags().Bool("enable-prefetching", false, "Enable prefetching")
	cmd.Flags().Bool("force-resolved-versions", false, "Only use versions from the Package.resolved file")
	cmd.Flags().Bool("only-use-versions-from-resolved-file", false, "Only use versions from the Package.resolved file")
	cmd.Flags().Bool("replace-scm-with-registry", false, "Use the registry to retrieve source control dependencies")
	cmd.Flags().Bool("skip-update", false, "Skip updating dependencies during resolution")
	cmd.Flags().Bool("use-registry-identity-for-scm", false, "Look up source control dependencies in the registry")

	cmd.Flags().StringArray("Xcc", nil, "Pass flag through to all C compiler invocations")
	cmd.Flags().StringArray("Xcxx", nil, "Pass flag through to all C++ compiler invocations")
	cmd.Flags().StringArray("Xlinker", nil, "Pass flag through to all linker invocations")
	cmd.Flags().StringArray("Xswiftc", nil, "Pass flag through to all Swift compiler invocations")
	cmd.Flags().Bool("auto-index-store", false, "Enable or disable indexing-while-building feature")
	cmd.Flags().String("build-system", "", "Specify the build system to use (native, swiftbuild, xcode)")
	cmd.Flags().StringP("configuration", "c", "", "Build with configuration (debug, release)")
	cmd.Flags().String("debug-info-format", "", "The debug info format to use (dwarf, codeview, none)")
	cmd.Flags().Bool("disable-dead-strip", false, "Disable dead code stripping by the linker")
	cmd.Flags().Bool("disable-experimental-strip-products", false, "Do not strip debug symbols from the final binary")
	cmd.Flags().Bool("disable-index-store", false, "Disable indexing-while-building feature")
	cmd.Flags().Bool("disable-local-rpath", false, "Disable adding $ORIGIN/@loader_path to the rpath")
	cmd.Flags().Bool("enable-dead-strip", false, "Enable dead code stripping by the linker")
	cmd.Flags().Bool("enable-experimental-strip-products", false, "Strip debug symbols from the final binary")
	cmd.Flags().Bool("enable-index-store", false, "Enable indexing-while-building feature")
	cmd.Flags().Bool("enable-parseable-module-interfaces", false, "Enable parseable module interfaces")
	cmd.Flags().String("experimental-codesize-profile-output-dir", "", "Directory to store code size profiling output files")
	cmd.Flags().Bool("experimental-enable-codesize-profile", false, "Generate SIL, LLVM IR, and optimization record files for code size profiling")
	cmd.Flags().String("explicit-target-dependency-import-check", "", "Check explicit target dependency imports (none, warn, error)")
	cmd.Flags().IntP("jobs", "j", 0, "The number of jobs to spawn in parallel during the build process")
	cmd.Flags().String("sanitize", "", "Turn on runtime checks for erroneous behavior (address, thread, undefined, scudo, fuzzer)")
	cmd.Flags().String("sdk", "", "Specify the SDK")
	cmd.Flags().String("swift-sdk", "", "Filter for selecting a specific Swift SDK to build with")
	cmd.Flags().String("toolchain", "", "Specify the toolchain")
	cmd.Flags().String("triple", "", "Specify the target triple")
	cmd.Flags().Bool("use-integrated-swift-driver", false, "Use the integrated Swift driver")

	cmd.Flags().Bool("disable-default-traits", false, "Disables all default traits of the package")
	cmd.Flags().Bool("enable-all-traits", false, "Enables all traits of the package")
	cmd.Flags().String("traits", "", "Enables the passed traits of the package")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"build-system":      carapace.ActionValues("native", "swiftbuild", "xcode"),
		"cache-path":        carapace.ActionDirectories(),
		"config-path":       carapace.ActionDirectories(),
		"configuration":     carapace.ActionValues("debug", "release"),
		"debug-info-format": carapace.ActionValues("codeview", "dwarf", "none"),
		"experimental-codesize-profile-output-dir": carapace.ActionDirectories(),
		"explicit-target-dependency-import-check":  carapace.ActionValues("error", "none", "warn"),
		"manifest-cache":                   carapace.ActionValues("local", "none", "shared"),
		"package-path":                     carapace.ActionDirectories(),
		"pkg-config-path":                  carapace.ActionDirectories(),
		"resolver-fingerprint-checking":    carapace.ActionValues("strict", "warn"),
		"resolver-signing-entity-checking": carapace.ActionValues("strict", "warn"),
		"sanitize":                         carapace.ActionValues("address", "fuzzer", "scudo", "thread", "undefined"),
		"scratch-path":                     carapace.ActionDirectories(),
		"security-path":                    carapace.ActionDirectories(),
		"swift-sdks-path":                  carapace.ActionDirectories(),
		"toolset":                          carapace.ActionFiles(".json"),
	})
}
