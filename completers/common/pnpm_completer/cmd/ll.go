package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var llCmd = &cobra.Command{
	Use:     "ll",
	Short:   "List installed packages in long format",
	Aliases: []string{"la"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(llCmd).Standalone()

	llCmd.Flags().String("depth", "0", "Max display depth of the dependency tree. `0` lists direct dependencies only; `-1` lists projects only")
	llCmd.Flags().BoolP("dev", "D", false, "Display only the dependency graph for packages in `devDependencies`")
	llCmd.Flags().Bool("exclude-peers", false, "Exclude peer dependencies")
	llCmd.Flags().StringSlice("find-by", nil, "Search by a finder function declared in `.pnpmfile.cjs`")
	llCmd.Flags().BoolP("global", "g", false, "")
	llCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	llCmd.Flags().Bool("json", false, "Show information in JSON format")
	llCmd.Flags().Bool("lockfile-only", false, "List packages from the lockfile only, without checking `node_modules`")
	llCmd.Flags().Bool("long", false, "Show extended information")
	llCmd.Flags().Bool("no-optional", false, "Don't display packages from `optionalDependencies`")
	llCmd.Flags().Bool("only-projects", false, "Display only dependencies that are also projects within the workspace")
	llCmd.Flags().Bool("optional", false, "Include packages from `optionalDependencies`")
	llCmd.Flags().Bool("parseable", false, "Show parseable output instead of tree view")
	llCmd.Flags().BoolP("prod", "P", false, "Display only the dependency graph for packages in `dependencies` and `optionalDependencies`")
	llCmd.Flags().Bool("production", false, "Display only the dependency graph for packages in `dependencies` and `optionalDependencies`")
	rootCmd.AddCommand(llCmd)

	carapace.Gen(llCmd).PositionalAnyCompletion(
		pnpm.ActionDependencyNames(),
	)
}
