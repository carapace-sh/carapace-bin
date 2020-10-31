package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Aliases: []string{"b"},
	Short:   "Compile a local package and all of its dependencies",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help' for details")
	buildCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	buildCmd.Flags().Bool("all-features", false, "Activate all available features")
	buildCmd.Flags().Bool("all-targets", false, "Build all targets")
	buildCmd.Flags().String("bench", "", "Build only the specified bench target")
	buildCmd.Flags().Bool("benches", false, "Build all benches")
	buildCmd.Flags().String("bin", "", "Build only the specified binary")
	buildCmd.Flags().Bool("bins", false, "Build all binaries")
	buildCmd.Flags().Bool("build-plan", false, "Output the build plan in JSON (unstable)")
	buildCmd.Flags().String("color", "", "Coloring: auto, always, never")
	buildCmd.Flags().String("example", "", "Build only the specified example")
	buildCmd.Flags().Bool("examples", false, "Build all examples")
	buildCmd.Flags().String("exclude", "", "Exclude packages from the build")
	buildCmd.Flags().String("features", "", "Space or comma separated list of features to activate")
	buildCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	buildCmd.Flags().BoolP("help", "h", false, "Prints help information")
	buildCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	buildCmd.Flags().Bool("lib", false, "Build only this package's library")
	buildCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	buildCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	buildCmd.Flags().String("message-format", "", "Error format")
	buildCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	buildCmd.Flags().Bool("offline", false, "Run without accessing the network")
	buildCmd.Flags().String("out-dir", "", "Copy final artifacts to this directory (unstable)")
	buildCmd.Flags().StringP("package", "p", "", "Package to build (see `cargo help pkgid`)")
	buildCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	buildCmd.Flags().BoolP("quiet", "q", false, "No output printed to stdout")
	buildCmd.Flags().Bool("release", false, "Build artifacts in release mode, with optimizations")
	buildCmd.Flags().String("target", "", "Build for the target triple")
	buildCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	buildCmd.Flags().String("test", "", "Build only the specified test target")
	buildCmd.Flags().Bool("tests", false, "Build all tests")
	buildCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")
	buildCmd.Flags().Bool("workspace", false, "Build all packages in the workspace")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"bench":   action.ActionTargets(buildCmd, action.TargetOpts{Bench: true}),
		"color":   action.ActionColorModes(),
		"example": action.ActionTargets(buildCmd, action.TargetOpts{Example: true}),
		"features": carapace.ActionMultiParts(",", func(args, parts []string) carapace.Action {
			return action.ActionFeatures(buildCmd).Invoke(args).Filter(parts).ToA()
		}),
		"exclude":        action.ActionWorkspaceMembers(buildCmd),
		"message-format": action.ActionMessageFormats(),
		"manifest-path":  carapace.ActionFiles(""),
		"out-dir":        carapace.ActionDirectories(),
		"package":        action.ActionDependencies(buildCmd),
		"profile":        action.ActionProfiles(buildCmd),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(buildCmd, action.TargetOpts{Test: true}),
	})
}
