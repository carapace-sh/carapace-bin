package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stored schemas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_listCmd).Standalone()

	schemas_listCmd.Flags().StringP("database", "d", "", "Path to log database")
	schemas_listCmd.Flags().Bool("full", false, "Output full schema contents")
	schemas_listCmd.Flags().Bool("json", false, "Output as JSON")
	schemas_listCmd.Flags().Bool("nl", false, "Output as newline-delimited JSON")
	schemas_listCmd.Flags().StringP("path", "p", "", "Path to log database")
	schemas_listCmd.Flags().StringP("query", "q", "", "Search for schemas matching this string")
	schemasCmd.AddCommand(schemas_listCmd)
}
