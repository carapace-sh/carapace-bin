package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Short:   "Execute all unit and integration tests and build examples of a local package",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("test"),
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	testCmd.Flags().Bool("all-features", false, "Activate all available features")
	testCmd.Flags().Bool("all-targets", false, "Test all targets (does not include doctests)")
	testCmd.Flags().StringSlice("bench", nil, "Test only the specified bench target")
	testCmd.Flags().Bool("benches", false, "Test all targets that have `bench = true` set")
	testCmd.Flags().StringSlice("bin", nil, "Test only the specified binary")
	testCmd.Flags().Bool("bins", false, "Test all binaries")
	testCmd.Flags().Bool("doc", false, "Test only this library's documentation")
	testCmd.Flags().StringSlice("example", nil, "Test only the specified example")
	testCmd.Flags().Bool("examples", false, "Test all examples")
	testCmd.Flags().StringSlice("exclude", nil, "Exclude packages from the test")
	testCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	testCmd.Flags().Bool("future-incompat-report", false, "Outputs a future incompatibility report at the end of the build")
	testCmd.Flags().BoolP("help", "h", false, "Print help")
	testCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	testCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs.")
	testCmd.Flags().Bool("lib", false, "Test only this package's library")
	testCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	testCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	testCmd.Flags().StringSlice("message-format", nil, "Error format")
	testCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	testCmd.Flags().Bool("no-fail-fast", false, "Run all tests regardless of failure")
	testCmd.Flags().Bool("no-run", false, "Compile, but don't run tests")
	testCmd.Flags().StringSliceP("package", "p", nil, "Package to run tests for")
	testCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	testCmd.Flags().BoolP("quiet", "q", false, "Display one character per test instead of one line")
	testCmd.Flags().BoolP("release", "r", false, "Build artifacts in release mode, with optimizations")
	testCmd.Flags().StringSlice("target", nil, "Build for the target triple")
	testCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	testCmd.Flags().StringSlice("test", nil, "Test only the specified test target")
	testCmd.Flags().Bool("tests", false, "Test all targets that have `test = true` set")
	testCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	testCmd.Flags().Bool("unit-graph", false, "Output build graph in JSON (unstable)")
	testCmd.Flags().Bool("workspace", false, "Test all packages in the workspace")
	testCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"bench":          action.ActionTargets(testCmd, action.TargetOpts{Bench: true}),
		"bin":            action.ActionTargets(testCmd, action.TargetOpts{Bin: true}),
		"example":        action.ActionTargets(testCmd, action.TargetOpts{Example: true}),
		"exclude":        action.ActionWorkspaceMembers(testCmd),
		"features":       action.ActionFeatures(testCmd).UniqueList(","),
		"lockfile-path":  carapace.ActionFiles(),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(testCmd, true),
		"profile":        action.ActionProfiles(testCmd),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(testCmd, action.TargetOpts{Test: true}),
	})
}
