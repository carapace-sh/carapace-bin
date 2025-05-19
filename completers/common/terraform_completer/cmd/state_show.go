package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_showCmd = &cobra.Command{
	Use:   "show [options] ADDRESS",
	Short: "Show a resource in the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_showCmd).Standalone()
	state_showCmd.Flags().StringS("state", "state", "", "Path to a Terraform state file")
	stateCmd.AddCommand(state_showCmd)

	state_showCmd.Flag("state").NoOptDefVal = " "

	carapace.Gen(state_showCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionFiles(),
	})

	carapace.Gen(state_showCmd).PositionalCompletion(
		action.ActionResources(state_showCmd).MultiParts("."),
	)
}
