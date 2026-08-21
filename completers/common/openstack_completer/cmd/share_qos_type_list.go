package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_qos_type_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List qos types",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_qos_type_listCmd).Standalone()

	share_qos_type_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_qos_type_listCmd.Flags().String("description", "", "Filter results by description.")
	share_qos_type_listCmd.Flags().String("description~", "", "Filter results matching a qos description pattern.")
	share_qos_type_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_qos_type_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_qos_type_listCmd.Flags().String("limit", "", "Limit the number of qos types returned.")
	share_qos_type_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_qos_type_listCmd.Flags().String("name", "", "Filter results by name.")
	share_qos_type_listCmd.Flags().String("name~", "", "Filter results matching a qos name pattern. ")
	share_qos_type_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_qos_type_listCmd.Flags().String("offset", "", "Start position of qos type records listing.")
	share_qos_type_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_qos_type_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_qos_type_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_qos_type_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_qos_type_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_qos_type_listCmd.Flags().String("sort_dir", "", "Sort direction, available values are ('asc', 'desc').")
	share_qos_type_listCmd.Flags().String("sort_key", "", "Key to be sorted with, available keys are ('id', 'name', 'created_at', 'updated_at').")
	share_qos_typeCmd.AddCommand(share_qos_type_listCmd)
}
