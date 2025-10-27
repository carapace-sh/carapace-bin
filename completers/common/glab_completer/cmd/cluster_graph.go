package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_graphCmd = &cobra.Command{
	Use:   "graph [flags]",
	Short: "Queries the Kubernetes object graph, using the GitLab Agent for Kubernetes. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_graphCmd).Standalone()

	cluster_graphCmd.Flags().StringP("agent", "a", "", "The numerical Agent ID to connect to.")
	cluster_graphCmd.Flags().Bool("apps", false, "Watch deployments, replicasets, daemonsets, and statefulsets in apps/v1 group.")
	cluster_graphCmd.Flags().Bool("batch", false, "Watch jobs and cronjobs in the batch/v1 group.")
	cluster_graphCmd.Flags().Bool("cluster-rbac", false, "Watch clusterroles and clusterrolebindings in the rbac.authorization.k8s.io/v1 group.")
	cluster_graphCmd.Flags().Bool("core", false, "Watch pods, secrets, configmaps, and serviceaccounts in the core/v1 group")
	cluster_graphCmd.Flags().Bool("crd", false, "Watch customresourcedefinitions in the apiextensions.k8s.io/v1 group.")
	cluster_graphCmd.Flags().Bool("ignore-arc-direction", false, "Ignore arc direction when evaluating roots connectivity. GitLab and agent 18.3+ required.")
	cluster_graphCmd.Flags().String("listen-addr", "", "Address to listen on.")
	cluster_graphCmd.Flags().String("listen-net", "", "Network on which to listen for connections.")
	cluster_graphCmd.Flags().Bool("log-watch-request", false, "Log watch request to stdout. Helpful for debugging.")
	cluster_graphCmd.Flags().StringSliceP("namespace", "n", nil, "Namespaces to watch. If not specified, all namespaces are watched with label and field selectors filtering.")
	cluster_graphCmd.Flags().String("ns-expression", "", "CEL expression to select namespaces. Evaluated before a namespace is watched and on any updates for the namespace object.")
	cluster_graphCmd.Flags().String("ns-field-selector", "", "Field selector to select namespaces.")
	cluster_graphCmd.Flags().String("ns-label-selector", "", "Label selector to select namespaces.")
	cluster_graphCmd.Flags().Bool("rbac", false, "Watch roles, and rolebindings in the rbac.authorization.k8s.io/v1 group.")
	cluster_graphCmd.Flags().StringSliceP("resource", "r", nil, "A list of resources to watch. You can see the list of resources your cluster supports by running 'kubectl api-resources'.")
	cluster_graphCmd.Flags().StringSlice("root-expression", nil, "CEL expression to select root objects. GitLab and agent 18.3+ required.")
	cluster_graphCmd.Flags().Bool("stdin", false, "Read watch request from standard input.")
	cluster_graphCmd.MarkFlagRequired("agent")
	clusterCmd.AddCommand(cluster_graphCmd)
}
