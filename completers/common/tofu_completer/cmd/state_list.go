package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_listCmd = &cobra.Command{
	Use:   "list [options] [address...]",
	Short: "List resources in the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_listCmd).Standalone()

	state_listCmd.Flags().StringS("id", "id", "", "Restrict output to paths with a resource having the specified ID.")
	state_listCmd.Flags().BoolS("json", "json", false, "Produce output in a machine-readable JSON format.")
	state_listCmd.Flags().StringS("json-into", "json-into", "", "Produce the same output as -json, but sent directly to the given file.")
	state_listCmd.Flags().StringS("state", "state", "", "Path to a OpenTofu state file to use to look up instances.")
	state_listCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	state_listCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	stateCmd.AddCommand(state_listCmd)

	state_listCmd.Flag("json-into").NoOptDefVal = " "
	state_listCmd.Flag("state").NoOptDefVal = " "
	state_listCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(state_listCmd).FlagCompletion(carapace.ActionMap{
		"json-into": carapace.ActionFiles(),
		"state":     carapace.ActionFiles(),
		"var-file":  carapace.ActionFiles(),
	})

	carapace.Gen(state_listCmd).PositionalAnyCompletion(
		action.ActionResources(state_listCmd).FilterArgs().MultiParts("_"),
	)
}
