package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_segment_rangeCmd = &cobra.Command{
	Use:   "range",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_segment_rangeCmd).Standalone()

	network_segmentCmd.AddCommand(network_segment_rangeCmd)
}
