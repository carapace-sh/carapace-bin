package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kube_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to a Kubernetes cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_loginCmd).Standalone()

	kube_loginCmd.Flags().Bool("all", false, "Generate a kubeconfig with every cluster the user has access to. Mutually exclusive with --labels or --query.")
	kube_loginCmd.Flags().String("as", "", "Configure custom Kubernetes user impersonation.")
	kube_loginCmd.Flags().String("as-groups", "", "Configure custom Kubernetes group impersonation.")
	kube_loginCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	kube_loginCmd.Flags().Bool("disable-access-request", false, "Disable automatic resource access requests.")
	kube_loginCmd.Flags().String("kube-namespace", "", "Configure the default Kubernetes namespace.")
	kube_loginCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	kube_loginCmd.Flags().StringP("namespace", "n", "", "Configure the default Kubernetes namespace.")
	kube_loginCmd.Flags().Bool("no-all", false, "Generate a kubeconfig with every cluster the user has access to. Mutually exclusive with --labels or --query.")
	kube_loginCmd.Flags().Bool("no-disable-access-request", false, "Disable automatic resource access requests.")
	kube_loginCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	kube_loginCmd.Flags().String("request-reason", "", "Reason for requesting access.")
	kube_loginCmd.Flags().String("set-context-name", "{{.ClusterName}}-{{.KubeName}}", "Define a custom context name. To use it with --all include \"{{.KubeName}}\".")
	kube_loginCmd.Flag("kube-namespace").Hidden = true
	kube_loginCmd.Flag("no-all").Hidden = true
	kube_loginCmd.Flag("no-disable-access-request").Hidden = true
	kubeCmd.AddCommand(kube_loginCmd)
}
