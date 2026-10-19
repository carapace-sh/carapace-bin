package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tofu"
	"github.com/spf13/cobra"
)

var state_replaceProviderCmd = &cobra.Command{
	Use:   "replace-provider [options] <FROM_PROVIDER_FQN> <TO_PROVIDER_FQN>",
	Short: "Replace provider in the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_replaceProviderCmd).Standalone()

	state_replaceProviderCmd.Flags().BoolS("auto-approve", "auto-approve", false, "Skip interactive approval.")
	state_replaceProviderCmd.Flags().StringS("backup", "backup", "", "Path where OpenTofu should write the backup state.")
	state_replaceProviderCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	state_replaceProviderCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	state_replaceProviderCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	state_replaceProviderCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	state_replaceProviderCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	state_replaceProviderCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	state_replaceProviderCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	state_replaceProviderCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	state_replaceProviderCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	state_replaceProviderCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	state_replaceProviderCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	state_replaceProviderCmd.Flags().StringS("state", "state", "", "Path to the state file to update.")
	state_replaceProviderCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	state_replaceProviderCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	stateCmd.AddCommand(state_replaceProviderCmd)

	carapace.Gen(state_replaceProviderCmd).FlagCompletion(carapace.ActionMap{
		"backup":      carapace.ActionFiles(),
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"state":       carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})

	carapace.Gen(state_replaceProviderCmd).PositionalCompletion(
		tofu.ActionProviders(),
		tofu.ActionProviders(),
	)
}
