package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var request_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for resources to request access to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(request_searchCmd).Standalone()

	request_searchCmd.Flags().Bool("all-kube-namespaces", false, "Search Pods in every namespace")
	request_searchCmd.Flags().String("kind", "", "Resource kind to search for (node, kube_cluster, db, app, windows_desktop, user_group, pod, secret, configmap, namespace, service, serviceaccount, kube_node, persistentvolume, persistentvolumeclaim, deployment, replicaset, statefulset, daemonset, clusterrole, role, clusterrolebinding, rolebinding, cronjob, job, certificatesigningrequest, ingress)")
	request_searchCmd.Flags().String("kube-cluster", "", "Kubernetes Cluster to search for Pods")
	request_searchCmd.Flags().String("kube-namespace", "", "Kubernetes Namespace to search for Pods")
	request_searchCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2)")
	request_searchCmd.Flags().Bool("no-all-kube-namespaces", false, "Search Pods in every namespace")
	request_searchCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"')")
	request_searchCmd.Flags().String("search", "", "List of comma separated search keywords or phrases enclosed in quotations (e.g. --search=foo,bar,\"some phrase\")")
	request_searchCmd.MarkFlagRequired("kind")
	request_searchCmd.Flag("no-all-kube-namespaces").Hidden = true
	requestCmd.AddCommand(request_searchCmd)
}
