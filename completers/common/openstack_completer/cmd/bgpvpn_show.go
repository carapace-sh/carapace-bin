package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information of a given BGP VPN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_showCmd).Standalone()

	bgpvpn_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgpvpn_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgpvpn_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgpvpn_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgpvpn_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgpvpn_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	bgpvpn_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgpvpn_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	bgpvpnCmd.AddCommand(bgpvpn_showCmd)
}
