package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_listCmd = &cobra.Command{
	Use:     "list [options] [address...]",
	Aliases: []string{"ls"},
	Short:   "List resources in the state",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_listCmd).Standalone()

	state_listCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	state_listCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	state_listCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	state_listCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	state_listCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	state_listCmd.Flags().StringS("id", "id", "", "Restrict output to paths with a resource having the specified ID.")
	state_listCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	state_listCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	state_listCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	state_listCmd.Flags().StringS("state", "state", "", "Path to a OpenTofu state file to use to look up instances.")
	state_listCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	state_listCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	stateCmd.AddCommand(state_listCmd)

	carapace.Gen(state_listCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"state":       carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})

	carapace.Gen(state_listCmd).PositionalAnyCompletion(
		action.ActionResources(state_listCmd).FilterArgs().MultiParts("_"),
	)
}
