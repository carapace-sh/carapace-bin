package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var state_pushCmd = &cobra.Command{
	Use:   "push [options] PATH",
	Short: "Update remote state from a local state file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_pushCmd).Standalone()

	state_pushCmd.Flags().BoolS("force", "force", false, "Write the state even if lineages don't match or the remote serial is higher.")
	state_pushCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	state_pushCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	state_pushCmd.Flags().BoolS("lock", "lock", false, "Don't hold a state lock during the operation.")
	state_pushCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	state_pushCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	state_pushCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	stateCmd.AddCommand(state_pushCmd)

	state_pushCmd.Flag("json-into").NoOptDefVal = " "
	state_pushCmd.Flag("lock-timeout").NoOptDefVal = " "
	state_pushCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(state_pushCmd).FlagCompletion(carapace.ActionMap{
		"json-into": carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(state_pushCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
