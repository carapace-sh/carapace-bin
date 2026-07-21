package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:     "outdated",
	Short:   "Check for outdated packages",
	GroupID: "review",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()

	outdatedCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects")
	outdatedCmd.Flags().Bool("compatible", false, "Print only versions that satisfy specs in package.json")
	outdatedCmd.Flags().BoolP("dev", "D", false, "Check only \"devDependencies\"")
	outdatedCmd.Flags().String("filter", "", "set filter")
	outdatedCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	outdatedCmd.Flags().String("format", "", "Output format (table, list, json)")
	outdatedCmd.Flags().BoolP("global", "g", false, "List outdated globally installed packages")
	outdatedCmd.Flags().Bool("long", false, "Show more details about the outdated packages")
	outdatedCmd.Flags().Bool("no-optional", false, "Don't check \"optionalDependencies\"")
	outdatedCmd.Flags().BoolP("prod", "P", false, "Check only \"dependencies\" and \"optionalDependencies\"")
	outdatedCmd.Flags().BoolP("recursive", "r", false, "Check for outdated dependencies in every package found in subdirectories")
	outdatedCmd.Flags().String("sort-by", "", "Sort outdated packages list")
	outdatedCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	rootCmd.AddCommand(outdatedCmd)

	carapace.Gen(outdatedCmd).FlagCompletion(carapace.ActionMap{
		"filter":      pnpm.ActionFilters(),
		"filter-prod": pnpm.ActionFilters(),
		"format":      carapace.ActionValues("table", "list", "json"),
		"sort-by":     carapace.ActionValues("name"),
	})

	carapace.Gen(outdatedCmd).PositionalAnyCompletion(
		pnpm.ActionDependencies(),
	)
}
