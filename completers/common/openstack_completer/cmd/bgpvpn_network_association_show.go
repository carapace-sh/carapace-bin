package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_network_association_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information of a given BGP VPN network association",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_network_association_showCmd).Standalone()

	bgpvpn_network_association_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgpvpn_network_association_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgpvpn_network_association_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgpvpn_network_association_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgpvpn_network_association_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgpvpn_network_association_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	bgpvpn_network_association_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgpvpn_network_association_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	bgpvpn_network_associationCmd.AddCommand(bgpvpn_network_association_showCmd)
}
