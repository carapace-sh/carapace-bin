package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var taintCmd = &cobra.Command{
	Use:   "taint [options] <address>",
	Short: "Mark a resource instance as not fully functional",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taintCmd).Standalone()

	taintCmd.Flags().BoolS("allow-missing", "allow-missing", false, "If specified, the command will succeed (exit code 0) even if the resource is missing.")
	taintCmd.Flags().StringS("backup", "backup", "", "Path to backup the existing state file before modifying.")
	taintCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	taintCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	taintCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	taintCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	taintCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	taintCmd.Flags().StringS("state", "state", "", "Path to read and save state.")
	taintCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\".")
	taintCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	taintCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(taintCmd)

	taintCmd.Flag("backup").NoOptDefVal = " "
	taintCmd.Flag("json-into").NoOptDefVal = " "
	taintCmd.Flag("lock-timeout").NoOptDefVal = " "
	taintCmd.Flag("state").NoOptDefVal = " "
	taintCmd.Flag("state-out").NoOptDefVal = " "
	taintCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(taintCmd).FlagCompletion(carapace.ActionMap{
		"backup":    carapace.ActionFiles(),
		"json-into": carapace.ActionFiles(),
		"state":     carapace.ActionFiles(),
		"state-out": carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(taintCmd).PositionalCompletion(
		action.ActionResources(taintCmd).MultiParts("."),
	)
}
