package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_peer_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List BGP peers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_peer_listCmd).Standalone()

	bgp_peer_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgp_peer_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgp_peer_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgp_peer_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgp_peer_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgp_peer_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgp_peer_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	bgp_peer_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	bgp_peer_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	bgp_peer_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	bgp_peerCmd.AddCommand(bgp_peer_listCmd)
}
