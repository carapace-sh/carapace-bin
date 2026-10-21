package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var kube_sessionsCmd = &cobra.Command{
	Use:   "sessions",
	Short: "Get a list of active Kubernetes sessions. (DEPRECATED: use tsh sessions ls --kind=kube instead.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_sessionsCmd).Standalone()

	kube_sessionsCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	kube_sessionsCmd.Flags().StringP("format", "f", "text", "Format output (text, json, yaml).")
	kubeCmd.AddCommand(kube_sessionsCmd)

	carapace.Gen(kube_sessionsCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
		"format":  tsh.ActionFormats(),
	})
}
