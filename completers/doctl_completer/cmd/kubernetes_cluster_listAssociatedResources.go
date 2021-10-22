package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_listAssociatedResourcesCmd = &cobra.Command{
	Use:   "list-associated-resources",
	Short: "Retrieve DigitalOcean resources associated with a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_listAssociatedResourcesCmd).Standalone()
	kubernetes_cluster_listAssociatedResourcesCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `Volumes`, `VolumeSnapshots`, `LoadBalancers`")
	kubernetes_cluster_listAssociatedResourcesCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_listAssociatedResourcesCmd)
}
