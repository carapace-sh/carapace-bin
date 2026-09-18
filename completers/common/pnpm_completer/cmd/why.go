package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var whyCmd = &cobra.Command{
	Use:   "why",
	Short: "Shows the packages that depend on `pkg`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whyCmd).Standalone()

	whyCmd.Flags().String("depth", "", "Max display depth of the reverse dependency tree")
	whyCmd.Flags().BoolP("dev", "D", false, "Display only the dependency graph for packages in `devDependencies`")
	whyCmd.Flags().Bool("exclude-peers", false, "Exclude peer dependencies")
	whyCmd.Flags().StringSlice("find-by", nil, "Search by a finder function declared in `.pnpmfile.cjs`")
	whyCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	whyCmd.Flags().Bool("json", false, "Show information in JSON format")
	whyCmd.Flags().Bool("long", false, "Show extended information")
	whyCmd.Flags().Bool("no-optional", false, "Don't display packages from `optionalDependencies`")
	whyCmd.Flags().Bool("optional", false, "Include packages from `optionalDependencies`")
	whyCmd.Flags().Bool("parseable", false, "Show parseable output instead of tree view")
	whyCmd.Flags().BoolP("prod", "P", false, "Display only the dependency graph for packages in `dependencies` and `optionalDependencies`")
	whyCmd.Flags().Bool("production", false, "Display only the dependency graph for packages in `dependencies` and `optionalDependencies`")
	rootCmd.AddCommand(whyCmd)

	carapace.Gen(whyCmd).PositionalCompletion(
		pnpm.ActionDependencyNames(),
	)
}
