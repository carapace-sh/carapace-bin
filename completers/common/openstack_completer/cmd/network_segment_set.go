package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_segment_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set network segment properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_segment_setCmd).Standalone()

	network_segment_setCmd.Flags().String("description", "", "Set network segment description")
	network_segment_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_segment_setCmd.Flags().String("name", "", "Set network segment name")
	network_segmentCmd.AddCommand(network_segment_setCmd)
}
