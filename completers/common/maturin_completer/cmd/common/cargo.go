package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/cargo"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rust"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

// AddCargoFlags is a subset of Cargo flags which maturin passes transparently
func AddCargoFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo")
	cmd.Flags().Bool("all-features", false, "Activate all available features")
	cmd.Flags().String("color", "", "Coloring: auto, always, never")
	cmd.Flags().String("config", "", "Override a configuration value (unstable)")
	cmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	cmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	cmd.Flags().Bool("future-incompat-report", false, "Outputs a future incompatibility report at the end of the build (unstable)")
	cmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	cmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	cmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	cmd.Flags().StringP("manifest-path", "m", "", "Path to Cargo.toml")
	cmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	cmd.Flags().Bool("offline", false, "Run without accessing the network")
	cmd.Flags().String("profile", "", "Build artifacts with the specified Cargo profile")
	cmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	cmd.Flags().String("target", "", "Build for the target triple")
	cmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	cmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	cmd.Flags().CountP("verbose", "v", "Use verbose output (-vv very verbose/build.rs output)")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"Z":             cargo.ActionNightlyFlags(),
		"color":         carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"features":      cargo.ActionFeatures("").UniqueList(","),
		"manifest-path": carapace.ActionFiles(),
		"target":        rust.ActionTargets(),
		"target-dir":    carapace.ActionDirectories(),
		"timings":       carapace.ActionValues("html", "json").UniqueList(","),
	})
}
