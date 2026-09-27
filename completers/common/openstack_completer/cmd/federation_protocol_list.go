package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federation_protocol_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List federation protocols",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federation_protocol_listCmd).Standalone()

	federation_protocol_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	federation_protocol_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	federation_protocol_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	federation_protocol_listCmd.Flags().String("identity-provider", "", "Identity provider to list (name or ID) (required)")
	federation_protocol_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	federation_protocol_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	federation_protocol_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	federation_protocol_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	federation_protocol_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	federation_protocol_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	federation_protocol_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	federation_protocol_listCmd.MarkFlagRequired("identity-provider")
	federation_protocolCmd.AddCommand(federation_protocol_listCmd)
}
