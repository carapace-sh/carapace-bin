package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var rebuildCmd = &cobra.Command{
	Use:     "rebuild",
	Short:   "Rebuild a package",
	Aliases: []string{"rb"},
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebuildCmd).Standalone()

	rebuildCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects since the specified commit/branch")
	rebuildCmd.Flags().String("filter", "", "set filter")
	rebuildCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	rebuildCmd.Flags().Bool("pending", false, "Rebuild packages that were not build during installation")
	rebuildCmd.Flags().BoolP("recursive", "r", false, "Rebuild every package found in subdirectories or every workspace package")
	rebuildCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	rootCmd.AddCommand(rebuildCmd)

	carapace.Gen(rebuildCmd).FlagCompletion(carapace.ActionMap{
		"filter":      pnpm.ActionFilters(),
		"filter-prod": pnpm.ActionFilters(),
	})

	carapace.Gen(rebuildCmd).PositionalAnyCompletion(
		carapace.Batch(
			pnpm.ActionDependencies(),
			pnpm.ActionWorkspaceDependencies(),
		).ToA(),
	)
}
