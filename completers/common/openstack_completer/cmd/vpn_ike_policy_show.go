package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ike_policy_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display IKE policy details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ike_policy_showCmd).Standalone()

	vpn_ike_policy_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_ike_policy_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_ike_policy_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_ike_policy_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_ike_policy_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_ike_policy_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	vpn_ike_policy_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_ike_policy_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	vpn_ike_policyCmd.AddCommand(vpn_ike_policy_showCmd)
}
