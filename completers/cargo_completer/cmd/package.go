package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Assemble the local package into a distributable tarball",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packageCmd).Standalone()

	packageCmd.Flags().Bool("all-features", false, "Activate all available features")
	packageCmd.Flags().Bool("allow-dirty", false, "Allow dirty working directories to be packaged")
	packageCmd.Flags().StringSlice("exclude", nil, "Don't assemble specified packages")
	packageCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	packageCmd.Flags().BoolP("help", "h", false, "Print help")
	packageCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs.")
	packageCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error")
	packageCmd.Flags().BoolP("list", "l", false, "Print files included in a package without making one")
	packageCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	packageCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	packageCmd.Flags().Bool("no-metadata", false, "Ignore warnings about a lack of human-usable metadata")
	packageCmd.Flags().Bool("no-verify", false, "Don't verify the contents by building them")
	packageCmd.Flags().StringSliceP("package", "p", nil, "Package(s) to assemble")
	packageCmd.Flags().StringSlice("target", nil, "Build for the target triple")
	packageCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	packageCmd.Flags().Bool("workspace", false, "Assemble all packages in the workspace")
	rootCmd.AddCommand(packageCmd)

	carapace.Gen(packageCmd).FlagCompletion(carapace.ActionMap{
		"exclude":       action.ActionDependencies(packageCmd, false),
		"features":      action.ActionFeatures(packageCmd).UniqueList(","),
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(packageCmd, false),
		"target-dir":    carapace.ActionDirectories(),
	})
}
