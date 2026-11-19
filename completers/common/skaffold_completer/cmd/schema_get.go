package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schema_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print a given skaffold.yaml's json schema",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schema_getCmd).Standalone()

	applyCommonFlagCompletion(schema_getCmd)
	schemaCmd.AddCommand(schema_getCmd)
}
