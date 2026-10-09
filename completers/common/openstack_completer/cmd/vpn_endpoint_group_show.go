package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_endpoint_group_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display endpoint group details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_endpoint_group_showCmd).Standalone()

	vpn_endpoint_group_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_endpoint_group_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_endpoint_group_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_endpoint_group_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_endpoint_group_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_endpoint_group_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	vpn_endpoint_group_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_endpoint_group_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	vpn_endpoint_groupCmd.AddCommand(vpn_endpoint_group_showCmd)
}
