package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Display commands for managing Kubernetes clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_clusterCmd).Standalone()
	kubernetesCmd.AddCommand(kubernetes_clusterCmd)
}
