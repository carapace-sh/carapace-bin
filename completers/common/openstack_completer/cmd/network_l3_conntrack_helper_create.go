package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_l3_conntrack_helper_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new L3 conntrack helper",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_l3_conntrack_helper_createCmd).Standalone()

	network_l3_conntrack_helper_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_l3_conntrack_helper_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_l3_conntrack_helper_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_l3_conntrack_helper_createCmd.Flags().String("helper", "", "The netfilter conntrack helper module")
	network_l3_conntrack_helper_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_l3_conntrack_helper_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_l3_conntrack_helper_createCmd.Flags().String("port", "", "The network port for the netfilter conntrack target rule")
	network_l3_conntrack_helper_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_l3_conntrack_helper_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_l3_conntrack_helper_createCmd.Flags().String("protocol", "", "The network protocol for the netfilter conntrack target rule")
	network_l3_conntrack_helper_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_l3_conntrack_helper_createCmd.MarkFlagRequired("helper")
	network_l3_conntrack_helper_createCmd.MarkFlagRequired("port")
	network_l3_conntrack_helper_createCmd.MarkFlagRequired("protocol")
	network_l3_conntrack_helperCmd.AddCommand(network_l3_conntrack_helper_createCmd)
}
