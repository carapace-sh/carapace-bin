package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_site_connection_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information of a given IPsec site connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_site_connection_showCmd).Standalone()

	vpn_ipsec_site_connection_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_ipsec_site_connection_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_ipsec_site_connection_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_ipsec_site_connection_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_ipsec_site_connection_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_ipsec_site_connection_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	vpn_ipsec_site_connection_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_ipsec_site_connection_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	vpn_ipsec_site_connectionCmd.AddCommand(vpn_ipsec_site_connection_showCmd)
}
