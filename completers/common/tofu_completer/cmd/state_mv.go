package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_mvCmd = &cobra.Command{
	Use:     "mv [options] <SOURCE> <DESTINATION>",
	Aliases: []string{"move"},
	Short:   "Move an item in the state",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_mvCmd).Standalone()

	state_mvCmd.Flags().StringS("backup", "backup", "", "Path where OpenTofu should write the backup state.")
	state_mvCmd.Flags().StringS("backup-out", "backup-out", "", "Path to backup the prior state file after destroying.")
	state_mvCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	state_mvCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	state_mvCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	state_mvCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	state_mvCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	state_mvCmd.Flags().BoolS("dry-run", "dry-run", false, "If set, prints out what would've been moved but doesn't actually move anything.")
	state_mvCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	state_mvCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	state_mvCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	state_mvCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	state_mvCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	state_mvCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	state_mvCmd.Flags().StringS("state", "state", "", "Path to the state file to update.")
	state_mvCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\".")
	state_mvCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	state_mvCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	stateCmd.AddCommand(state_mvCmd)

	carapace.Gen(state_mvCmd).FlagCompletion(carapace.ActionMap{
		"backup":      carapace.ActionFiles(),
		"backup-out":  carapace.ActionFiles(),
		"deprecation": action.ActionDeprecationLevels(),
		"json-into":   carapace.ActionFiles(),
		"state":       carapace.ActionFiles(),
		"state-out":   carapace.ActionFiles(),
		"var-file":    carapace.ActionFiles(),
	})

	carapace.Gen(state_mvCmd).PositionalCompletion(
		action.ActionResources(state_mvCmd).MultiParts("."),
		action.ActionResources(state_mvCmd).MultiParts("."),
	)
}
