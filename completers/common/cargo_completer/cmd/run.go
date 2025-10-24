package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run a binary or example of the local package",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("run"),
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("all-features", false, "Activate all available features")
	runCmd.Flags().StringSlice("bin", nil, "Name of the bin target to run")
	runCmd.Flags().StringSlice("example", nil, "Name of the example target to run")
	runCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	runCmd.Flags().BoolP("help", "h", false, "Print help")
	runCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	runCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs.")
	runCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error")
	runCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	runCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	runCmd.Flags().StringSlice("message-format", nil, "Error format")
	runCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	runCmd.Flags().StringP("package", "p", "", "Package with the target to run")
	runCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	runCmd.Flags().BoolP("release", "r", false, "Build artifacts in release mode, with optimizations")
	runCmd.Flags().StringSlice("target", nil, "Build for the target triple")
	runCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	runCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	runCmd.Flags().Bool("unit-graph", false, "Output build graph in JSON (unstable)")
	runCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"bin":            action.ActionTargets(runCmd, action.TargetOpts{Bin: true}),
		"example":        action.ActionTargets(runCmd, action.TargetOpts{Example: true}),
		"features":       action.ActionFeatures(runCmd).UniqueList(","),
		"lockfile-path":  carapace.ActionFiles(),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(runCmd, true),
		"profile":        action.ActionProfiles(runCmd),
		"target-dir":     carapace.ActionDirectories(),
	})
}
