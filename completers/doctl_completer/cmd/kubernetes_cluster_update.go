package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a Kubernetes cluster's configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_updateCmd).Standalone()
	kubernetes_cluster_updateCmd.Flags().Bool("auto-upgrade", false, "A boolean flag indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window (default false). To enable automatic upgrade, supply --auto-upgrade(=true).")
	kubernetes_cluster_updateCmd.Flags().String("cluster-name", "", "Specifies a new cluster name")
	kubernetes_cluster_updateCmd.Flags().String("maintenance-window", "any=00:00", "Sets the beginning of the four hour maintenance window for the cluster. Syntax is in the format: 'day=HH:MM', where time is in UTC. Day can be: `any`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.")
	kubernetes_cluster_updateCmd.Flags().Bool("set-current-context", true, "Boolean specifying whether to set the current kubectl context to that of the new cluster")
	kubernetes_cluster_updateCmd.Flags().Bool("surge-upgrade", false, "Boolean specifying whether to enable surge-upgrade for the cluster")
	kubernetes_cluster_updateCmd.Flags().StringSlice("tag", []string{}, "A comma-separated list of tags to apply to the cluster. Existing user-provided tags will be removed from the cluster if they are not specified.")
	kubernetes_cluster_updateCmd.Flags().Bool("update-kubeconfig", true, "Boolean specifying whether to update the cluster in your kubeconfig")
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_updateCmd)
}
