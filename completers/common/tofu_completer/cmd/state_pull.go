package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_pullCmd = &cobra.Command{
	Use:   "pull [options]",
	Short: "Pull current state and output to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_pullCmd).Standalone()

	state_pullCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	state_pullCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	state_pullCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	state_pullCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	state_pullCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	state_pullCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	state_pullCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	state_pullCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	stateCmd.AddCommand(state_pullCmd)

	carapace.Gen(state_pullCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"var-file":    carapace.ActionFiles(),
	})
}
