package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var debug_loadbalancer_getConfigCmd = &cobra.Command{
	Use:   "get-config CLUSTERNAME",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_loadbalancer_getConfigCmd).Standalone()

	debug_loadbalancerCmd.AddCommand(debug_loadbalancer_getConfigCmd)

	carapace.Gen(debug_loadbalancer_getConfigCmd).PositionalCompletion(
		k3d.ActionClusters(),
	)
}
