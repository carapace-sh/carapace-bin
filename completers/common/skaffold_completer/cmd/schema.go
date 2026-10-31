package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "List JSON schemas used to validate skaffold.yaml configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemaCmd).Standalone()

	applyCommonFlagCompletion(schemaCmd)
	rootCmd.AddCommand(schemaCmd)
}
