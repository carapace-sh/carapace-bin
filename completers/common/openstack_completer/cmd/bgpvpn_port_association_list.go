package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_port_association_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List BGP VPN port associations for a given BGP VPN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_port_association_listCmd).Standalone()

	bgpvpn_port_association_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgpvpn_port_association_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgpvpn_port_association_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgpvpn_port_association_listCmd.Flags().Bool("long", false, "List additional fields in output")
	bgpvpn_port_association_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgpvpn_port_association_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgpvpn_port_association_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgpvpn_port_association_listCmd.Flags().String("property", "", "Filter property to apply on returned BGP VPNs (repeat to filter on multiple properties)")
	bgpvpn_port_association_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	bgpvpn_port_association_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	bgpvpn_port_association_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	bgpvpn_port_association_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	bgpvpn_port_associationCmd.AddCommand(bgpvpn_port_association_listCmd)
}
