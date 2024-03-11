package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var kube_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Get a list of Kubernetes clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_lsCmd).Standalone()

	kube_lsCmd.Flags().BoolP("all", "R", false, "List Kubernetes clusters from all clusters and proxies.")
	kube_lsCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	kube_lsCmd.Flags().StringP("format", "f", "", "Format output (text, json, yaml)")
	kube_lsCmd.Flags().Bool("no-all", false, "List Kubernetes clusters from all clusters and proxies.")
	kube_lsCmd.Flags().Bool("no-quiet", false, "Quiet mode.")
	kube_lsCmd.Flags().Bool("no-verbose", false, "Show an untruncated list of labels.")
	kube_lsCmd.Flags().String("query", "", "Query by predicate language enclosed in single quotes. Supports ==, !=, &&, and || (e.g. --query='labels[\"key1\"] == \"value1\" && labels[\"key2\"] != \"value2\"')")
	kube_lsCmd.Flags().BoolP("quiet", "q", false, "Quiet mode.")
	kube_lsCmd.Flags().String("search", "", "List of comma separated search keywords or phrases enclosed in quotations (e.g. --search=foo,bar,\"some phrase\")")
	kube_lsCmd.Flags().BoolP("verbose", "v", false, "Show an untruncated list of labels.")
	kube_lsCmd.Flag("no-all").Hidden = true
	kube_lsCmd.Flag("no-quiet").Hidden = true
	kube_lsCmd.Flag("no-verbose").Hidden = true
	kubeCmd.AddCommand(kube_lsCmd)

	carapace.Gen(kube_lsCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"format":  tsh.ActionFormats(),
	})
}
