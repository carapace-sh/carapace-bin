package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var docCmd = &cobra.Command{
	Use:     "doc",
	Short:   "Build a package's documentation",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("doc"),
}

func init() {
	carapace.Gen(docCmd).Standalone()

	docCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	docCmd.Flags().Bool("all-features", false, "Activate all available features")
	docCmd.Flags().StringSlice("bin", nil, "Document only the specified binary")
	docCmd.Flags().Bool("bins", false, "Document all binaries")
	docCmd.Flags().Bool("document-private-items", false, "Document private items")
	docCmd.Flags().StringSlice("example", nil, "Document only the specified example")
	docCmd.Flags().Bool("examples", false, "Document all examples")
	docCmd.Flags().StringSlice("exclude", nil, "Exclude packages from the build")
	docCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	docCmd.Flags().BoolP("help", "h", false, "Print help")
	docCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	docCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs.")
	docCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error")
	docCmd.Flags().Bool("lib", false, "Document only this package's library")
	docCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	docCmd.Flags().StringSlice("message-format", nil, "Error format")
	docCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	docCmd.Flags().Bool("no-deps", false, "Don't build documentation for dependencies")
	docCmd.Flags().Bool("open", false, "Opens the docs in a browser after the operation")
	docCmd.Flags().StringSliceP("package", "p", nil, "Package to document")
	docCmd.Flags().String("profile", "", "Build artifacts with the specified profile")
	docCmd.Flags().BoolP("release", "r", false, "Build artifacts in release mode, with optimizations")
	docCmd.Flags().StringSlice("target", nil, "Build for the target triple")
	docCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	docCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	docCmd.Flags().Bool("unit-graph", false, "Output build graph in JSON (unstable)")
	docCmd.Flags().Bool("workspace", false, "Document all packages in the workspace")
	docCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(docCmd)

	carapace.Gen(docCmd).FlagCompletion(carapace.ActionMap{
		"bin":            action.ActionTargets(docCmd, action.TargetOpts{Bin: true}),
		"exclude":        action.ActionWorkspaceMembers(docCmd),
		"features":       action.ActionFeatures(docCmd).UniqueList(","),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(docCmd, true),
		"profile":        action.ActionProfiles(docCmd),
		"target-dir":     carapace.ActionDirectories(),
	})
}
