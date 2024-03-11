package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"t"},
	Short:   "Runs a package's test script, if one was provided",
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().String("changed-files-ignore-pattern", "", "Defines files to ignore when filtering for changed projects")
	testCmd.Flags().String("filter", "", "set filter                                                packages except \"foo\"")
	testCmd.Flags().String("filter-prod", "", "Restricts the scope to package names matching the given pattern")
	testCmd.Flags().String("test-pattern", "", "Defines files related to tests")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"filter":      pnpm.ActionFilter(),
		"filter-prod": pnpm.ActionFilter(),
	})

	// TODO positional completion
}
