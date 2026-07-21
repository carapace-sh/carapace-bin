package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "uninstall", "un"},
	Short:   "Removes packages from `node_modules` and from the project's `package.json`",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects since the specified commit/branch")
	removeCmd.Flags().String("filter", "", "set filter")
	removeCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	removeCmd.Flags().Bool("global-dir", false, "Specify a custom directory to store global packages")
	removeCmd.Flags().BoolP("recursive", "r", false, "Remove from every package found in subdirectories or from every workspace package")
	removeCmd.Flags().BoolP("save-dev", "D", false, "Remove the dependency only from \"devDependencies\"")
	removeCmd.Flags().BoolP("save-optional", "O", false, "Remove the dependency only from \"optionalDependencies\"")
	removeCmd.Flags().BoolP("save-prod", "P", false, "Remove the dependency only from \"dependencies\"")
	removeCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"filter":      pnpm.ActionFilters(),
		"filter-prod": pnpm.ActionFilters(),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		carapace.Batch(
			pnpm.ActionDependencies(),
			pnpm.ActionWorkspaceDependencies(),
		).ToA(),
	)
}
