package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_showCmd = &cobra.Command{
	Use:   "show [options] <ADDRESS>",
	Short: "Show a resource in the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_showCmd).Standalone()

	state_showCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	state_showCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	state_showCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	state_showCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	state_showCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	state_showCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	state_showCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	state_showCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	state_showCmd.Flags().BoolS("show-sensitive", "show-sensitive", false, "If specified, sensitive values will be displayed.")
	state_showCmd.Flags().StringS("state", "state", "", "Path to a OpenTofu state file to use to look up instances.")
	state_showCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	state_showCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	stateCmd.AddCommand(state_showCmd)

	carapace.Gen(state_showCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"state":       carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})

	carapace.Gen(state_showCmd).PositionalCompletion(
		action.ActionResources(state_showCmd).MultiParts("."),
	)
}
