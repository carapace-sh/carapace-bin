package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_auto_allocated_topologyCmd = &cobra.Command{
	Use:   "topology",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_auto_allocated_topologyCmd).Standalone()

	network_auto_allocatedCmd.AddCommand(network_auto_allocated_topologyCmd)
}
