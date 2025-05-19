package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var config_setContextCmd = &cobra.Command{
	Use:   "set-context [NAME | --current] [--cluster=cluster_nickname] [--user=user_nickname] [--namespace=namespace]",
	Short: "Set a context entry in kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setContextCmd).Standalone()

	config_setContextCmd.Flags().String("cluster", "", "cluster for the context entry in kubeconfig")
	config_setContextCmd.Flags().Bool("current", false, "Modify the current context")
	config_setContextCmd.Flags().String("namespace", "", "namespace for the context entry in kubeconfig")
	config_setContextCmd.Flags().String("user", "", "user for the context entry in kubeconfig")
	configCmd.AddCommand(config_setContextCmd)

	carapace.Gen(config_setContextCmd).PositionalCompletion(
		kubectl.ActionContexts(),
	)
}
