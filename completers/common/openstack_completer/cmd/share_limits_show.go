package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_limits_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a list of share limits for a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_limits_showCmd).Standalone()

	share_limits_showCmd.Flags().Bool("absolute", false, "Get the absolute limits for the user")
	share_limits_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_limits_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_limits_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_limits_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_limits_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_limits_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_limits_showCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_limits_showCmd.Flags().Bool("rate", false, "Get the API rate limits for the user")
	share_limits_showCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_limits_showCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_limits_showCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_limitsCmd.AddCommand(share_limits_showCmd)
}
