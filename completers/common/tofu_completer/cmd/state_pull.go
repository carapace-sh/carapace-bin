package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var state_pullCmd = &cobra.Command{
	Use:   "pull [options]",
	Short: "Pull current state and output to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_pullCmd).Standalone()

	state_pullCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	state_pullCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	stateCmd.AddCommand(state_pullCmd)

	state_pullCmd.Flag("var-file").NoOptDefVal = " "

	carapace.Gen(state_pullCmd).FlagCompletion(carapace.ActionMap{
		"var-file": carapace.ActionFiles(),
	})
}
