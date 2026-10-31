package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schema_listCmd = &cobra.Command{
	Use:    "list",
	Short:  "List skaffold.yaml's json schema versions",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schema_listCmd).Standalone()

	schema_listCmd.Flags().StringP("output", "o", "plain", "Type of output: `plain` or `json`.")
	applyCommonFlagCompletion(schema_listCmd)
	schemaCmd.AddCommand(schema_listCmd)
}
