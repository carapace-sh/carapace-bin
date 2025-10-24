package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:     "fix",
	Short:   "Automatically fix lint warnings reported by rustc",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("fix"),
}

func init() {
	carapace.Gen(fixCmd).Standalone()

	fixCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	fixCmd.Flags().Bool("all-features", false, "Activate all available features")
	fixCmd.Flags().Bool("all-targets", false, "Fix all targets (default)")
	fixCmd.Flags().Bool("allow-dirty", false, "Fix code even if the working directory is dirty or has staged changes")
	fixCmd.Flags().Bool("allow-no-vcs", false, "Fix code even if a VCS was not detected")
	fixCmd.Flags().Bool("allow-staged", false, "Fix code even if the working directory has staged changes")
	fixCmd.Flags().StringSlice("bench", nil, "Fix only the specified bench target")
	fixCmd.Flags().Bool("benches", false, "Fix all targets that have `bench = true` set")
	fixCmd.Flags().StringSlice("bin", nil, "Fix only the specified binary")
	fixCmd.Flags().Bool("bins", false, "Fix all binaries")
	fixCmd.Flags().Bool("broken-code", false, "Fix code even if it already has compiler errors")
	fixCmd.Flags().Bool("edition", false, "Fix in preparation for the next edition")
	fixCmd.Flags().Bool("edition-idioms", false, "Fix warnings to migrate to the idioms of an edition")
	fixCmd.Flags().StringSlice("example", nil, "Fix only the specified example")
	fixCmd.Flags().Bool("examples", false, "Fix all examples")
	fixCmd.Flags().StringSlice("exclude", nil, "Exclude packages from the fixes")
	fixCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	fixCmd.Flags().BoolP("help", "h", false, "Print help")
	fixCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	fixCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs.")
	fixCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error")
	fixCmd.Flags().Bool("lib", false, "Fix only this package's library")
	fixCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	fixCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	fixCmd.Flags().StringSlice("message-format", nil, "Error format")
	fixCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	fixCmd.Flags().StringSliceP("package", "p", nil, "Package(s) to fix")
	fixCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	fixCmd.Flags().BoolP("release", "r", false, "Fix artifacts in release mode, with optimizations")
	fixCmd.Flags().StringSlice("target", nil, "Fix for the target triple")
	fixCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	fixCmd.Flags().StringSlice("test", nil, "Fix only the specified test target")
	fixCmd.Flags().Bool("tests", false, "Fix all targets that have `test = true` set")
	fixCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	fixCmd.Flags().Bool("workspace", false, "Fix all packages in the workspace")
	fixCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(fixCmd)

	carapace.Gen(fixCmd).FlagCompletion(carapace.ActionMap{
		"benches":        action.ActionTargets(fixCmd, action.TargetOpts{Bench: true}),
		"bin":            action.ActionTargets(fixCmd, action.TargetOpts{Bin: true}),
		"example":        action.ActionTargets(fixCmd, action.TargetOpts{Example: true}),
		"exclude":        action.ActionWorkspaceMembers(fixCmd),
		"features":       action.ActionFeatures(fixCmd).UniqueList(","),
		"lockfile-path":  carapace.ActionFiles(),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(fixCmd, true),
		"profile":        action.ActionProfiles(fixCmd),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(fixCmd, action.TargetOpts{Test: true}),
	})
}
