package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_segment_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new network segment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_segment_createCmd).Standalone()

	network_segment_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_segment_createCmd.Flags().String("description", "", "Network segment description")
	network_segment_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_segment_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_segment_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_segment_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_segment_createCmd.Flags().String("network", "", "Network this network segment belongs to (name or ID)")
	network_segment_createCmd.Flags().String("network-type", "", "Network type of this network segment (flat, geneve, gre, local, vlan or vxlan)")
	network_segment_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_segment_createCmd.Flags().String("physical-network", "", "Physical network name of this network segment")
	network_segment_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_segment_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_segment_createCmd.Flags().String("segment", "", "Segment identifier for this network segment which is based on the network type, VLAN ID for vlan network type and tunnel ID for geneve, gre and vxlan network types")
	network_segment_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_segment_createCmd.MarkFlagRequired("network")
	network_segment_createCmd.MarkFlagRequired("network-type")
	network_segmentCmd.AddCommand(network_segment_createCmd)
}
