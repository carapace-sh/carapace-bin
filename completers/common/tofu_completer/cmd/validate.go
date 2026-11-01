package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:     "validate [options] [path]",
	Short:   "Check whether the configuration is valid",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	validateCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	validateCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	validateCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	validateCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	validateCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	validateCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	validateCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	validateCmd.Flags().StringArrayS("lint", "lint", nil, "Specify the linting rules to be executed.")
	validateCmd.Flags().BoolS("no-color", "no-color", false, "If specified, output won't contain any color.")
	validateCmd.Flags().BoolS("no-tests", "no-tests", false, "If specified, OpenTofu will not validate test files.")
	validateCmd.Flags().StringS("test-directory", "test-directory", "", "Set the OpenTofu test directory, defaults to \"tests\".")
	validateCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	validateCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(validateCmd)

	carapace.Gen(validateCmd).FlagCompletion(carapace.ActionMap{
		"deprecation":    action.ActionDeprecationLevels(),
		"json-into":      carapace.ActionFiles(),
		"lint":           action.ActionLintRules(),
		"test-directory": carapace.ActionDirectories(),
		"var-file":       carapace.ActionFiles(),
	})

	carapace.Gen(validateCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
