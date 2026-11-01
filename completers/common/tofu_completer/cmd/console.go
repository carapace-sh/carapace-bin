package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console [options]",
	Short: "Try OpenTofu expressions at an interactive command prompt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consoleCmd).Standalone()

	consoleCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	consoleCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	consoleCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	consoleCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	consoleCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	consoleCmd.Flags().BoolS("input", "input", false, "Disable prompting for required input variables that are not set some other way.")
	consoleCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	consoleCmd.Flags().StringS("json-into", "json-into", "", "Streams the output of the console to the given file.")
	consoleCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	consoleCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	consoleCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	consoleCmd.Flags().BoolS("show-sensitive", "show-sensitive", false, "If specified, sensitive values will be displayed.")
	consoleCmd.Flags().StringS("state", "state", "", "Legacy option for the local backend only.")
	consoleCmd.Flags().StringArrayS("var", "var", nil, "Set a variable in the OpenTofu configuration.")
	consoleCmd.Flags().StringS("var-file", "var-file", "", "Set variables in the OpenTofu configuration from a file.")

	rootCmd.AddCommand(consoleCmd)

	carapace.Gen(consoleCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"state":       carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})
}
