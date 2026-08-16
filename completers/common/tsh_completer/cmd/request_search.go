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

	request_searchCmd.Flags().Bool("all-kube-namespaces", false, "Search Pods in every namespace.")
	request_searchCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	request_searchCmd.Flags().String("kind", "", "Resource kind to search for (node, kube_cluster, kube_resource, db, app, windows_desktop, user_group, saml_idp_service_provider, aws_ic_account, aws_ic_account_assignment, git_server).  Mutually exclusive with --roles.")
	request_searchCmd.Flags().String("kube-api-group", "", "Kubernetes API group to search for resources.")
	request_searchCmd.Flags().String("kube-cluster", "", "Kubernetes Cluster to search for Pods.")
	request_searchCmd.Flags().String("kube-kind", "", "Kubernetes resource kind name (plural) to search for. Required with --kind=\"kube_resource\" Ex: pods, deployments, namespaces, etc.")
	request_searchCmd.Flags().String("kube-namespace", "default", "Kubernetes Namespace to search for Pods.")
	request_searchCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	request_searchCmd.Flags().String("namespace", "default", "Kubernetes Namespace to search for Pods.")
	request_searchCmd.Flags().Bool("no-all-kube-namespaces", false, "Search Pods in every namespace.")
	request_searchCmd.Flags().Bool("no-roles", false, "List requestable roles instead of searching for resources. Mutually exclusive with --kind.")
	request_searchCmd.Flags().Bool("no-verbose", false, "Verbose table output, shows full label output.")
	request_searchCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	request_searchCmd.Flags().Bool("roles", false, "List requestable roles instead of searching for resources. Mutually exclusive with --kind.")
	request_searchCmd.Flags().String("search", "", "List of comma separated search keywords or phrases enclosed in quotations (e.g. --search=foo,bar,\"some phrase\").")
	request_searchCmd.Flags().BoolP("verbose", "v", false, "Verbose table output, shows full label output.")
	request_searchCmd.Flag("kube-namespace").Hidden = true
	request_searchCmd.Flag("no-all-kube-namespaces").Hidden = true
	request_searchCmd.Flag("no-roles").Hidden = true
	request_searchCmd.Flag("no-verbose").Hidden = true
	requestCmd.AddCommand(request_searchCmd)
}
