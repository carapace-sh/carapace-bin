package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var api_schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Print or write the bundled API schema",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(api_schemaCmd).Standalone()

	api_schemaCmd.Flags().Bool("json", false, "")
	api_schemaCmd.Flags().String("output", "", "")
	apiCmd.AddCommand(api_schemaCmd)

	carapace.Gen(api_schemaCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
