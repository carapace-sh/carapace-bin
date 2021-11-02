package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_cluster_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_cluster_createCmd).Standalone()
	kubernetes_cluster_createCmd.Flags().StringSlice("1-clicks", []string{}, "Comma-separated list of 1-Click Applications to install on the kubernetes cluster. To see a list of 1-Click Applications available run doctl kubernetes 1-click list")
	kubernetes_cluster_createCmd.Flags().Bool("auto-upgrade", false, "A boolean flag indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window (default false). To enable automatic upgrade, supply --auto-upgrade(=true).")
	kubernetes_cluster_createCmd.Flags().Int("count", 3, "Number of nodes in the default node pool (incompatible with --node-pool)")
	kubernetes_cluster_createCmd.Flags().Bool("ha", false, "A boolean flag indicating whether the cluster will be configured with a highly-available control plane (default false). To enable the HA control plane, supply --ha(=true).")
	kubernetes_cluster_createCmd.Flags().String("maintenance-window", "any=00:00", "Sets the beginning of the four hour maintenance window for the cluster. Syntax is in the format: `day=HH:MM`, where time is in UTC. Day can be: `any`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.")
	kubernetes_cluster_createCmd.Flags().StringSlice("node-pool", []string{}, "Comma-separated list of node pools, defined using semicolon-separated configuration values and surrounded by quotes (incompatible with --size and --count)")
	kubernetes_cluster_createCmd.Flags().String("region", "nyc1", "Cluster region. Possible values: see `doctl kubernetes options regions` (required)")
	kubernetes_cluster_createCmd.Flags().Bool("set-current-context", true, "Boolean that specifies whether to set the current kubectl context to that of the new cluster")
	kubernetes_cluster_createCmd.Flags().String("size", "s-1vcpu-2gb", "Machine size to use when creating nodes in the default node pool (incompatible with --node-pool). Possible values: see `doctl kubernetes options sizes`")
	kubernetes_cluster_createCmd.Flags().Bool("surge-upgrade", true, "Boolean specifying whether to enable surge-upgrade for the cluster")
	kubernetes_cluster_createCmd.Flags().StringSlice("tag", []string{}, "Comma-separated list of tags to apply to the cluster, in addition to the default tags of `k8s` and `k8s:$K8S_CLUSTER_ID`.")
	kubernetes_cluster_createCmd.Flags().Bool("update-kubeconfig", true, "Boolean that specifies whether to add a configuration context for the new cluster to your kubectl")
	kubernetes_cluster_createCmd.Flags().String("version", "latest", "Kubernetes version. Possible values: see `doctl kubernetes options versions`")
	kubernetes_cluster_createCmd.Flags().String("vpc-uuid", "", "Kubernetes UUID. Must be the UUID of a valid VPC in the same region specified for the cluster.")
	kubernetes_cluster_createCmd.Flags().Bool("wait", true, "Boolean that specifies whether to wait for cluster creation to complete before returning control to the terminal")
	kubernetes_clusterCmd.AddCommand(kubernetes_cluster_createCmd)
}
