package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_segmentCmd = &cobra.Command{
	Use:   "segment",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_segmentCmd).Standalone()

	networkCmd.AddCommand(network_segmentCmd)
}
