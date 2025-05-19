package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_updateKubeconfigCmd = &cobra.Command{
	Use:   "update-kubeconfig [flags]",
	Short: "Update selected kubeconfig.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_updateKubeconfigCmd).Standalone()

	cluster_agent_updateKubeconfigCmd.Flags().StringP("agent", "a", "", "The numeric agent ID to create the kubeconfig entry for.")
	cluster_agent_updateKubeconfigCmd.PersistentFlags().String("kubeconfig", "", "Use a particular kubeconfig file.")
	cluster_agent_updateKubeconfigCmd.PersistentFlags().BoolP("use-context", "u", false, "Use as default context.")
	cluster_agent_updateKubeconfigCmd.MarkFlagRequired("agent")
	cluster_agentCmd.AddCommand(cluster_agent_updateKubeconfigCmd)

	// TODO flag completion
}
