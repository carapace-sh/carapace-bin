package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testsCmd = &cobra.Command{
	Use:     "tests",
	Short:   "Run Homebrew's unit and integration tests",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testsCmd).Standalone()

	testsCmd.Flags().Bool("byebug", false, "Enable debugging using byebug.")
	testsCmd.Flags().Bool("changed", false, "Only runs tests on files that were changed from the master branch.")
	testsCmd.Flags().Bool("coverage", false, "Generate code coverage reports.")
	testsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	testsCmd.Flags().Bool("fail-fast", false, "Exit early on the first failing test.")
	testsCmd.Flags().Bool("generic", false, "Run only OS-agnostic tests.")
	testsCmd.Flags().Bool("help", false, "Show this message.")
	testsCmd.Flags().Bool("online", false, "Include tests that use the GitHub API and tests that use any of the taps for official external commands.")
	testsCmd.Flags().Bool("only", false, "Run only <test_script>`_spec.rb`. Appending `:`<line_number> will start at a specific line.")
	testsCmd.Flags().Bool("profile", false, "Run the test suite serially to find the <n> slowest tests.")
	testsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	testsCmd.Flags().Bool("seed", false, "Randomise tests with the specified <value> instead of a random seed.")
	testsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(testsCmd)
}
