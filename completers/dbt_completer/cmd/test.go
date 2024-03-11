package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Runs tests on data in deployed models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("defer", false, "Defer to the state variable for resolving unselected nodes")
	testCmd.Flags().BoolP("fail-fast", "x", false, "Stop execution upon a first test failure")
	testCmd.Flags().Bool("favor-state", false, "Defer to the state variable for resolving unselected nodes")
	testCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	testCmd.Flags().String("indirect-selection", "", "Select all tests that are adjacent to selected resources")
	testCmd.Flags().String("log-path", "", "Configure the 'log-path'")
	testCmd.Flags().Bool("no-defer", false, "Do not defer to the state variable for resolving unselected nodes")
	testCmd.Flags().Bool("no-favor-state", false, "If defer is set, expect standard defer behaviour")
	testCmd.Flags().Bool("no-version-check", false, "Skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	testCmd.Flags().String("selector", "", "The selector name to use")
	testCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare")
	testCmd.Flags().Bool("store-failures", false, "Store test results (failing rows) in the database")
	testCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	testCmd.Flags().String("target-path", "", "Configure the 'target-path'")
	testCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	testCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(testCmd)
	addSelectionFlags(testCmd)
	addModelsFlag(testCmd)
	addProfileFlag(testCmd)
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"indirect-selection": carapace.ActionValues("eager", "cautious", "buildable"),
		"log-path":           carapace.ActionFiles(),
		"state":              carapace.ActionDirectories(),
		"target-path":        carapace.ActionDirectories(),
	})
}
