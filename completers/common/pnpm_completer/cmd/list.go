package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List installed packages",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().String("depth", "0", "Max display depth of the dependency tree. `0` lists direct dependencies only; `-1` lists projects only")
	listCmd.Flags().BoolP("dev", "D", false, "Display only the dependency graph for packages in `devDependencies`")
	listCmd.Flags().Bool("exclude-peers", false, "Exclude peer dependencies")
	listCmd.Flags().StringSlice("find-by", nil, "Search by a finder function declared in `.pnpmfile.cjs`")
	listCmd.Flags().BoolP("global", "g", false, "")
	listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	listCmd.Flags().Bool("json", false, "Show information in JSON format")
	listCmd.Flags().Bool("lockfile-only", false, "List packages from the lockfile only, without checking `node_modules`")
	listCmd.Flags().Bool("long", false, "Show extended information")
	listCmd.Flags().Bool("no-optional", false, "Don't display packages from `optionalDependencies`")
	listCmd.Flags().Bool("only-projects", false, "Display only dependencies that are also projects within the workspace")
	listCmd.Flags().Bool("optional", false, "Include packages from `optionalDependencies`")
	listCmd.Flags().Bool("parseable", false, "Show parseable output instead of tree view")
	listCmd.Flags().BoolP("prod", "P", false, "Display only the dependency graph for packages in `dependencies` and `optionalDependencies`")
	listCmd.Flags().Bool("production", false, "Display only the dependency graph for packages in `dependencies` and `optionalDependencies`")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalAnyCompletion(
		pnpm.ActionDependencyNames(),
	)
}
