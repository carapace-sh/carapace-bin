package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_restartCmd = &cobra.Command{
	Use:   "restart [NAME [NAME...] | --all]",
	Short: "Restart existing k3d cluster(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_restartCmd).Standalone()

	cluster_restartCmd.Flags().BoolP("all", "a", false, "Restart all existing clusters")
	cluster_restartCmd.Flags().String("timeout", "", "Maximum waiting time for '--wait' before canceling/returning.")
	cluster_restartCmd.Flags().Bool("wait", false, "Wait for the server(s) (and loadbalancer) to be ready before returning.")
	clusterCmd.AddCommand(cluster_restartCmd)
}
