package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_segment_range_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network segment range(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_segment_range_deleteCmd).Standalone()

	network_segment_rangeCmd.AddCommand(network_segment_range_deleteCmd)
}
