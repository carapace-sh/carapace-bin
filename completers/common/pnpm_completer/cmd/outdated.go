package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Check for outdated package and GitHub Actions dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()

	outdatedCmd.Flags().Bool("compatible", false, "Print only versions that satisfy the ranges in package.json")
	outdatedCmd.Flags().BoolP("dev", "D", false, "Check only \"devDependencies\"")
	outdatedCmd.Flags().String("format", "table", "Output format")
	outdatedCmd.Flags().BoolP("global", "g", false, "Check globally installed packages")
	outdatedCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	outdatedCmd.Flags().Bool("include-github-actions", false, "Also check GitHub Actions dependencies in workflow and action files")
	outdatedCmd.Flags().Bool("json", false, "Shorthand for `--format json`")
	outdatedCmd.Flags().Bool("long", false, "Print details about the outdated packages (homepage, deprecation notice)")
	outdatedCmd.Flags().Bool("no-optional", false, "Don't check \"optionalDependencies\"")
	outdatedCmd.Flags().Bool("no-table", false, "Shorthand for `--format list`. Good for small consoles")
	outdatedCmd.Flags().Bool("optional", false, "Include \"optionalDependencies\"")
	outdatedCmd.Flags().BoolP("prod", "P", false, "Check only \"dependencies\" and \"optionalDependencies\"")
	outdatedCmd.Flags().Bool("production", false, "Check only \"dependencies\" and \"optionalDependencies\"")
	outdatedCmd.Flags().String("sort-by", "", "Sorting method. Currently only `name` is supported; the default sorts by the size of the version change, then by name")

	carapace.Gen(outdatedCmd).FlagCompletion(carapace.ActionMap{
		"format":  carapace.ActionValues("table", "list", "json"),
		"sort-by": carapace.ActionValues("name"),
	})

	rootCmd.AddCommand(outdatedCmd)
}
