package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proxy_kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Start local proxy for Kubernetes access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_kubeCmd).Standalone()

	proxy_kubeCmd.Flags().String("as", "", "Configure custom Kubernetes user impersonation.")
	proxy_kubeCmd.Flags().String("as-groups", "", "Configure custom Kubernetes group impersonation.")
	proxy_kubeCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	proxy_kubeCmd.Flags().Bool("exec", false, "Run the proxy in the background and reexec into a new shell with $KUBECONFIG already pointed to our config file.")
	proxy_kubeCmd.Flags().String("exec-arg", "", "Arguments to pass to the executed command (can be specified multiple times).")
	proxy_kubeCmd.Flags().String("exec-cmd", "", "Command to execute when --exec is enabled. If not specified, defaults to $SHELL or /bin/bash. Implicitly enables exec mode.")
	proxy_kubeCmd.Flags().StringP("format", "f", "unix", "Optional format to print the commands for setting environment variables, one of: unix, command-prompt, powershell, text. Default is unix.")
	proxy_kubeCmd.Flags().String("kube-namespace", "", "Configure the default Kubernetes namespace.")
	proxy_kubeCmd.Flags().String("labels", "", "List of comma separated labels to filter by labels (e.g. key1=value1,key2=value2).")
	proxy_kubeCmd.Flags().StringP("namespace", "n", "", "Configure the default Kubernetes namespace.")
	proxy_kubeCmd.Flags().Bool("no-exec", false, "Run the proxy in the background and reexec into a new shell with $KUBECONFIG already pointed to our config file.")
	proxy_kubeCmd.Flags().StringP("port", "p", "", "Specifies the source port used by the proxy listener.")
	proxy_kubeCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"').")
	proxy_kubeCmd.Flags().String("set-context-name", "{{.ClusterName}}-{{.KubeName}}", "Define a custom context name or template.")
	proxy_kubeCmd.Flag("kube-namespace").Hidden = true
	proxy_kubeCmd.Flag("no-exec").Hidden = true
	proxyCmd.AddCommand(proxy_kubeCmd)
}
