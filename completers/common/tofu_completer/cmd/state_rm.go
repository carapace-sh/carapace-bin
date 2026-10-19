package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_rmCmd = &cobra.Command{
	Use:   "rm [options] ADDRESS...",
	Short: "Remove instances from the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_rmCmd).Standalone()

	state_rmCmd.Flags().StringS("backup", "backup", "", "Path where OpenTofu should write the backup state.")
	state_rmCmd.Flags().BoolS("dry-run", "dry-run", false, "If set, prints out what would've been removed but doesn't actually remove anything.")
	state_rmCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "Continue even if remote and local OpenTofu versions are incompatible.")
	state_rmCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	state_rmCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	state_rmCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	state_rmCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	state_rmCmd.Flags().StringS("state", "state", "", "Path to the state file to update.")
	state_rmCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	state_rmCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	stateCmd.AddCommand(state_rmCmd)

	state_rmCmd.Flag("backup").NoOptDefVal = " "
	state_rmCmd.Flag("json-into").NoOptDefVal = " "
	state_rmCmd.Flag("lock-timeout").NoOptDefVal = " "
	state_rmCmd.Flag("state").NoOptDefVal = " "
	state_rmCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(state_rmCmd).FlagCompletion(carapace.ActionMap{
		"backup":    carapace.ActionFiles(),
		"json-into": carapace.ActionFiles(),
		"state":     carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(state_rmCmd).PositionalAnyCompletion(
		action.ActionResources(state_rmCmd).FilterArgs().MultiParts("."),
	)
}
