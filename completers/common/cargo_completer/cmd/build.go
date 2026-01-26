package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rust"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Compile a local package and all of its dependencies",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("build"),
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	buildCmd.Flags().Bool("all-features", false, "Activate all available features")
	buildCmd.Flags().Bool("all-targets", false, "Build all targets")
	buildCmd.Flags().String("artifact-dir", "", "Copy final artifacts to this directory (unstable)")
	buildCmd.Flags().StringSlice("bench", nil, "Build only the specified bench target")
	buildCmd.Flags().Bool("benches", false, "Build all targets that have `bench = true` set")
	buildCmd.Flags().StringSlice("bin", nil, "Build only the specified binary")
	buildCmd.Flags().Bool("bins", false, "Build all binaries")
	buildCmd.Flags().StringSlice("example", nil, "Build only the specified example")
	buildCmd.Flags().Bool("examples", false, "Build all examples")
	buildCmd.Flags().StringSlice("exclude", nil, "Exclude packages from the build")
	buildCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	buildCmd.Flags().Bool("future-incompat-report", false, "Outputs a future incompatibility report at the end of the build")
	buildCmd.Flags().BoolP("help", "h", false, "Print help")
	buildCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	buildCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs.")
	buildCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error")
	buildCmd.Flags().Bool("lib", false, "Build only this package's library")
	buildCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	buildCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	buildCmd.Flags().StringSlice("message-format", nil, "Error format")
	buildCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	buildCmd.Flags().StringSliceP("package", "p", nil, "Package to build (see `cargo help pkgid`)")
	buildCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	buildCmd.Flags().BoolP("release", "r", false, "Build artifacts in release mode, with optimizations")
	buildCmd.Flags().StringSlice("target", nil, "Build for the target triple")
	buildCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	buildCmd.Flags().StringSlice("test", nil, "Build only the specified test target")
	buildCmd.Flags().Bool("tests", false, "Build all targets that have `test = true` set")
	buildCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	buildCmd.Flags().Bool("unit-graph", false, "Output build graph in JSON (unstable)")
	buildCmd.Flags().Bool("workspace", false, "Build all packages in the workspace")
	buildCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"artifact-dir":   carapace.ActionDirectories(),
		"bench":          action.ActionTargets(buildCmd, action.TargetOpts{Bench: true}),
		"bin":            action.ActionTargets(buildCmd, action.TargetOpts{Bin: true}),
		"example":        action.ActionTargets(buildCmd, action.TargetOpts{Example: true}),
		"exclude":        action.ActionWorkspaceMembers(buildCmd),
		"features":       action.ActionFeatures(buildCmd).UniqueList(","),
		"lockfile-path":  carapace.ActionFiles(),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(buildCmd, true),
		"profile":        action.ActionProfiles(buildCmd),
		"target":         rust.ActionTargets(),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(buildCmd, action.TargetOpts{Test: true}),
	})
}
