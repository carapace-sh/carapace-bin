package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_dslCmd = &cobra.Command{
	Use:   "dsl",
	Short: "Convert LLM's schema DSL to a JSON schema",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_dslCmd).Standalone()

	schemas_dslCmd.Flags().Bool("multi", false, "Wrap in an array")
	schemasCmd.AddCommand(schemas_dslCmd)
}
