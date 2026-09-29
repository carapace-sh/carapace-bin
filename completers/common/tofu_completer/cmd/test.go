package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [options]",
	Short: "Execute integration tests for OpenTofu modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	testCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	testCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	testCmd.Flags().StringArrayS("filter", "filter", nil, "If specified, OpenTofu will only execute the test files that match the given pattern.")
	testCmd.Flags().BoolS("json", "json", false, "If specified, machine readable output will be printed in JSON format.")
	testCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	testCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	testCmd.Flags().StringS("test-directory", "test-directory", "", "Set the OpenTofu test directory, defaults to \"tests\".")
	testCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	testCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	testCmd.Flags().BoolS("verbose", "verbose", false, "Print the plan or state for each test run block as it executes.")
	rootCmd.AddCommand(testCmd)

	testCmd.Flag("json-into").NoOptDefVal = " "
	testCmd.Flag("test-directory").NoOptDefVal = " "
	testCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"json-into":      carapace.ActionFiles(),
		"test-directory": carapace.ActionDirectories(),
		"var-file":       carapace.ActionFiles(),
	})
}
