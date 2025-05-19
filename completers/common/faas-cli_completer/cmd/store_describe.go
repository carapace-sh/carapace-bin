package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var store_describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Show details of OpenFaaS function from a store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_describeCmd).Standalone()
	store_describeCmd.Flags().BoolP("verbose", "v", false, "Verbose output for the field values")
	storeCmd.AddCommand(store_describeCmd)

	carapace.Gen(store_describeCmd).PositionalCompletion(
		action.ActionFunctionsStore(),
	)
}
