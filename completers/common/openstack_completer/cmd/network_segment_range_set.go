package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_segment_range_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set network segment range properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_segment_range_setCmd).Standalone()

	network_segment_range_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_segment_range_setCmd.Flags().String("maximum", "", "Set network segment range maximum segment identifier")
	network_segment_range_setCmd.Flags().String("minimum", "", "Set network segment range minimum segment identifier")
	network_segment_range_setCmd.Flags().String("name", "", "Set network segment name")
	network_segment_rangeCmd.AddCommand(network_segment_range_setCmd)
}
