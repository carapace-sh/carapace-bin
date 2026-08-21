package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_dragent_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List dynamic routing agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_dragent_listCmd).Standalone()

	bgp_dragent_listCmd.Flags().String("bgp-speaker", "", "List dynamic routing agents hosting a BGP speaker (name or ID)")
	bgp_dragent_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgp_dragent_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgp_dragent_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgp_dragent_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgp_dragent_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgp_dragent_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgp_dragent_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	bgp_dragent_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	bgp_dragent_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	bgp_dragent_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	bgp_dragentCmd.AddCommand(bgp_dragent_listCmd)
}
