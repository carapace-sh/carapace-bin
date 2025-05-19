package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "run Go test(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().BoolS("asan", "asan", false, "run in address sanitizer builder mode")
	testCmd.Flags().StringS("banner", "banner", "", "banner prefix")
	testCmd.Flags().BoolS("compile-only", "compile-only", false, "compile tests, but don't run them")
	testCmd.Flags().BoolS("k", "k", false, "keep going even when error occurred")
	testCmd.Flags().BoolS("list", "list", false, "list available tests")
	testCmd.Flags().BoolS("msan", "msan", false, "run in memory sanitizer builder mode")
	testCmd.Flags().BoolS("no-rebuild", "no-rebuild", false, "overrides -rebuild (historical dreg)")
	testCmd.Flags().BoolS("race", "race", false, "run in race builder mode (different set of tests)")
	testCmd.Flags().BoolS("rebuild", "rebuild", false, "rebuild everything first")
	testCmd.Flags().StringS("run", "run", "", "run only those tests matching the regular expression")
	testCmd.Flags().BoolS("v", "v", false, "verbosity")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"run": golang.ActionTests(golang.TestOpts{}.Default()),
	})
}
