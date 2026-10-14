package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var availability_zone_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List availability zones and their status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(availability_zone_listCmd).Standalone()

	availability_zone_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	availability_zone_listCmd.Flags().Bool("compute", false, "List compute availability zones")
	availability_zone_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	availability_zone_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	availability_zone_listCmd.Flags().Bool("long", false, "List additional fields in output")
	availability_zone_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	availability_zone_listCmd.Flags().Bool("network", false, "List network availability zones")
	availability_zone_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	availability_zone_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	availability_zone_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	availability_zone_listCmd.Flags().Bool("share", false, "List share availability zones")
	availability_zone_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	availability_zone_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	availability_zone_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	availability_zone_listCmd.Flags().Bool("volume", false, "List volume availability zones")
	availability_zoneCmd.AddCommand(availability_zone_listCmd)
}
