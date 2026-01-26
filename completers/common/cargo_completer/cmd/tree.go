package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Display a tree visualization of a dependency graph",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(treeCmd).Standalone()

	treeCmd.Flags().Bool("all-features", false, "Activate all available features")
	treeCmd.Flags().String("charset", "", "Character set to use in output")
	treeCmd.Flags().String("depth", "", "Maximum display depth of the dependency tree")
	treeCmd.Flags().Bool("duplicate", false, "Show only dependencies which come in multiple versions (implies -i)")
	treeCmd.Flags().BoolP("duplicates", "d", false, "Show only dependencies which come in multiple versions (implies -i)")
	treeCmd.Flags().StringSliceP("edges", "e", nil, "The kinds of dependencies to display")
	treeCmd.Flags().StringSlice("exclude", nil, "Exclude specific workspace members")
	treeCmd.Flags().StringSliceP("features", "F", nil, "Space or comma separated list of features to activate")
	treeCmd.Flags().StringP("format", "f", "", "Format string used for printing dependencies")
	treeCmd.Flags().BoolP("help", "h", false, "Print help")
	treeCmd.Flags().StringSliceP("invert", "i", nil, "Invert the tree direction and focus on the given package")
	treeCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	treeCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	treeCmd.Flags().Bool("no-dedupe", false, "Do not de-duplicate (repeats all shared dependencies)")
	treeCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	treeCmd.Flags().StringSliceP("package", "p", nil, "Package to be used as the root of the tree")
	treeCmd.Flags().String("prefix", "", "Change the prefix (indentation) of how each entry is displayed")
	treeCmd.Flags().StringSlice("prune", nil, "Prune the given package from the display of the dependency tree")
	treeCmd.Flags().StringSlice("target", nil, "Filter dependencies matching the given target-triple (default host platform). Pass `all` to include all targets.")
	treeCmd.Flags().Bool("workspace", false, "Display the tree for all packages in the workspace")
	treeCmd.Flag("duplicate").Hidden = true
	rootCmd.AddCommand(treeCmd)

	// TODO flag completion
	carapace.Gen(treeCmd).FlagCompletion(carapace.ActionMap{
		"charset": carapace.ActionValues("utf8", "ascii"),
		"edges": carapace.ActionValues(
			"all",
			"normal",
			"build",
			"dev",
			"features",
			"public",
			"no-normal",
			"no-build",
			"no-dev",
			"no-proc-macro",
		).UniqueList(","),
		"exclude":       action.ActionWorkspaceMembers(treeCmd).UniqueList(","),
		"features":      action.ActionFeatures(treeCmd).UniqueList(","),
		"invert":        action.ActionDependencies(treeCmd, false),
		"lockfile-path": carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(treeCmd, false),
		"prune":         action.ActionDependencies(treeCmd, false),
	})
}
