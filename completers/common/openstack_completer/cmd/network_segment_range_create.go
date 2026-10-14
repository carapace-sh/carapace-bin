package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_segment_range_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new network segment range",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_segment_range_createCmd).Standalone()

	network_segment_range_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_segment_range_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_segment_range_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_segment_range_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_segment_range_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_segment_range_createCmd.Flags().String("maximum", "", "Maximum segment identifier for this network segment range which is based on the network type, VLAN ID for vlan network type and tunnel ID for geneve, gre and vxlan network types")
	network_segment_range_createCmd.Flags().String("minimum", "", "Minimum segment identifier for this network segment range which is based on the network type, VLAN ID for vlan network type and tunnel ID for geneve, gre and vxlan network types")
	network_segment_range_createCmd.Flags().String("network-type", "", "Network type of this network segment range (geneve, gre, vlan or vxlan)")
	network_segment_range_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_segment_range_createCmd.Flags().String("physical-network", "", "Physical network name of this network segment range")
	network_segment_range_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_segment_range_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_segment_range_createCmd.Flags().Bool("private", false, "Network segment range is assigned specifically to the project")
	network_segment_range_createCmd.Flags().String("project", "", "Network segment range owner (name or ID).")
	network_segment_range_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_segment_range_createCmd.Flags().Bool("shared", false, "Network segment range is shared with other projects")
	network_segment_range_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_segment_range_createCmd.MarkFlagRequired("maximum")
	network_segment_range_createCmd.MarkFlagRequired("minimum")
	network_segment_range_createCmd.MarkFlagRequired("network-type")
	network_segment_rangeCmd.AddCommand(network_segment_range_createCmd)
}
