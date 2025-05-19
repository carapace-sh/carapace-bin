package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rustcCmd = &cobra.Command{
	Use:   "rustc",
	Short: "Compile a package, and pass extra options to the compiler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rustcCmd).Standalone()

	rustcCmd.Flags().Bool("all-features", false, "Activate all available features")
	rustcCmd.Flags().Bool("all-targets", false, "Build all targets")
	rustcCmd.Flags().StringSlice("bench", nil, "Build only the specified bench target")
	rustcCmd.Flags().Bool("benches", false, "Build all bench targets")
	rustcCmd.Flags().StringSlice("bin", nil, "Build only the specified binary")
	rustcCmd.Flags().Bool("bins", false, "Build all binaries")
	rustcCmd.Flags().StringSlice("crate-type", nil, "Comma separated list of types of crates for the compiler to emit")
	rustcCmd.Flags().StringSlice("example", nil, "Build only the specified example")
	rustcCmd.Flags().Bool("examples", false, "Build all examples")
	rustcCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	rustcCmd.Flags().Bool("future-incompat-report", false, "Outputs a future incompatibility report at the end of the build")
	rustcCmd.Flags().BoolP("help", "h", false, "Print help")
	rustcCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	rustcCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs.")
	rustcCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error")
	rustcCmd.Flags().Bool("lib", false, "Build only this package's library")
	rustcCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	rustcCmd.Flags().StringSlice("message-format", nil, "Error format")
	rustcCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	rustcCmd.Flags().StringP("package", "p", "", "Package to build")
	rustcCmd.Flags().String("print", "", "Output compiler information without compiling")
	rustcCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	rustcCmd.Flags().BoolP("release", "r", false, "Build artifacts in release mode, with optimizations")
	rustcCmd.Flags().StringSlice("target", nil, "Target triple which compiles will be for")
	rustcCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	rustcCmd.Flags().StringSlice("test", nil, "Build only the specified test target")
	rustcCmd.Flags().Bool("tests", false, "Build all test targets")
	rustcCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	rustcCmd.Flags().Bool("unit-graph", false, "Output build graph in JSON (unstable)")
	rustcCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(rustcCmd)

	// TODO flags
	carapace.Gen(rustcCmd).FlagCompletion(carapace.ActionMap{
		"bench":          action.ActionTargets(rustcCmd, action.TargetOpts{Bench: true}),
		"bin":            action.ActionTargets(rustcCmd, action.TargetOpts{Bin: true}),
		"crate-type":     carapace.ActionValues("bin", "lib", "rlib", "dylib", "cdylib", "staticlib", "and proc-macr").UniqueList(","),
		"example":        action.ActionTargets(rustcCmd, action.TargetOpts{Example: true}),
		"features":       action.ActionFeatures(rustcCmd).UniqueList(","),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(rustcCmd, false),
		"print":          carapace.ActionValues("crate-name", "file-names", "sysroot", "target-libdir", "cfg", "calling-conventions", "target-list", "target-cpus", "target-features", "relocation-models", "code-models", "tls-models", "target-spec-json", "native-static-libs", "stack-protector-strategies", "link-args"),
		"profile":        action.ActionProfiles(rustcCmd),
		"test":           action.ActionTargets(rustcCmd, action.TargetOpts{Test: true}),
	})
}
