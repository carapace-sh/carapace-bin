package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var untaintCmd = &cobra.Command{
	Use:   "untaint [options] <address>",
	Short: "Remove the 'tainted' state from a resource instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(untaintCmd).Standalone()

	untaintCmd.Flags().BoolS("allow-missing", "allow-missing", false, "If specified, the command will succeed (exit code 0) even if the resource is missing.")
	untaintCmd.Flags().StringS("backup", "backup", "", "Path to backup the existing state file before modifying.")
	untaintCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	untaintCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	untaintCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	untaintCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	untaintCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	untaintCmd.Flags().StringS("state", "state", "", "Path to read and save state.")
	untaintCmd.Flags().StringS("state-out", "state-out", "", "Path to write state to that is different than \"-state\".")
	untaintCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	untaintCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	rootCmd.AddCommand(untaintCmd)

	untaintCmd.Flag("backup").NoOptDefVal = " "
	untaintCmd.Flag("json-into").NoOptDefVal = " "
	untaintCmd.Flag("lock-timeout").NoOptDefVal = " "
	untaintCmd.Flag("state").NoOptDefVal = " "
	untaintCmd.Flag("state-out").NoOptDefVal = " "
	untaintCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(untaintCmd).FlagCompletion(carapace.ActionMap{
		"backup":    carapace.ActionFiles(),
		"json-into": carapace.ActionFiles(),
		"state":     carapace.ActionFiles(),
		"state-out": carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(untaintCmd).PositionalCompletion(
		action.ActionResources(untaintCmd).MultiParts("."),
	)
}
