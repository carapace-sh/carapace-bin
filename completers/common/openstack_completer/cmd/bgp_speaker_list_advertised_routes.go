package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_speaker_list_advertised_routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "List routes advertised",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speaker_list_advertised_routesCmd).Standalone()

	bgp_speaker_list_advertised_routesCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgp_speaker_list_advertised_routesCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgp_speaker_list_advertised_routesCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgp_speaker_list_advertised_routesCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgp_speaker_list_advertised_routesCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgp_speaker_list_advertised_routesCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgp_speaker_list_advertised_routesCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	bgp_speaker_list_advertised_routesCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	bgp_speaker_list_advertised_routesCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	bgp_speaker_list_advertised_routesCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	bgp_speaker_list_advertisedCmd.AddCommand(bgp_speaker_list_advertised_routesCmd)
}
