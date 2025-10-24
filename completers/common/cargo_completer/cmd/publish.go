package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:     "publish",
	Short:   "Upload a package to the registry",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("publish"),
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().Bool("all-features", false, "Activate all available features")
	publishCmd.Flags().Bool("allow-dirty", false, "Allow dirty working directories to be packaged")
	publishCmd.Flags().BoolP("dry-run", "n", false, "Perform all checks without uploading")
	publishCmd.Flags().StringSlice("exclude", nil, "Don't publish specified packages")
	publishCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	publishCmd.Flags().BoolP("help", "h", false, "Print help")
	publishCmd.Flags().String("index", "", "Registry index URL to upload the package to")
	publishCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs.")
	publishCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error")
	publishCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	publishCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	publishCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	publishCmd.Flags().Bool("no-verify", false, "Don't verify the contents by building them")
	publishCmd.Flags().StringSliceP("package", "p", nil, "Package(s) to publish")
	publishCmd.Flags().String("registry", "", "Registry to upload the package to")
	publishCmd.Flags().StringSlice("target", nil, "Build for the target triple")
	publishCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	publishCmd.Flags().String("token", "", "Token to use when uploading")
	publishCmd.Flags().Bool("workspace", false, "Publish all packages in the workspace")
	rootCmd.AddCommand(publishCmd)

	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"features":      action.ActionFeatures(publishCmd).UniqueList(","),
		"lockfile-path": carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(publishCmd, false),
		"registry":      action.ActionRegistries(),
		"target-dir":    carapace.ActionDirectories(),
	})
}
