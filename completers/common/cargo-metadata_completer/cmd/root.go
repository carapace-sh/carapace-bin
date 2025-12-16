package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/cargo"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cargo-metadata",
	Short: "Output the resolved dependencies of a package",
	Long:  "https://doc.rust-lang.org/cargo/commands/cargo-metadata.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("Z", "Z", "", "Unstable (nightly-only) flags to Cargo, see 'cargo -Z help'")
	rootCmd.Flags().Bool("all-features", false, "Activate all available features")
	rootCmd.Flags().String("color", "", "Coloring: auto, always, never")
	rootCmd.Flags().String("config", "", "Override a configuration value")
	rootCmd.Flags().StringP("features", "F", "", "Space or comma separated list of features to activate")
	rootCmd.Flags().String("filter-platform", "", "Only include resolve dependencies matching the given")
	rootCmd.Flags().String("format-version", "", "Format version [possible values: 1]")
	rootCmd.Flags().Bool("frozen", false, "Require Cargo.lock and cache are up to date")
	rootCmd.Flags().BoolP("help", "h", false, "Print help information")
	rootCmd.Flags().Bool("locked", false, "Require Cargo.lock is up to date")
	rootCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	rootCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	rootCmd.Flags().Bool("no-deps", false, "Output information only about the workspace members and don't")
	rootCmd.Flags().Bool("offline", false, "Run without accessing the network")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	rootCmd.Flags().BoolP("verbose", "v", false, "Use verbose output (-vv very verbose/build.rs output)")

	// TODO more flags
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"Z":     cargo.ActionNightlyFlags(),
		"color": carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"features": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return cargo.ActionFeatures(rootCmd.Flag("manifest-path").Value.String()).UniqueList(",")
		}),
		"format-version": carapace.ActionValues("1"),
		"manifest-path":  carapace.ActionFiles(),
	})
}
