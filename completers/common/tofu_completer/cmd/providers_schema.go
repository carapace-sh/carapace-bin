package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var providers_schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Show schemas for the providers used in the configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providers_schemaCmd).Standalone()

	providers_schemaCmd.Flags().BoolS("json", "json", false, "Prints out a json representation of the providers used in the current configuration.")
	providers_schemaCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	providers_schemaCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")

	providers_schemaCmd.Flag("var-file").NoOptDefVal = " "

	providersCmd.AddCommand(providers_schemaCmd)

	carapace.Gen(providers_schemaCmd).FlagCompletion(carapace.ActionMap{
		"var-file": carapace.ActionFiles(),
	})
}
