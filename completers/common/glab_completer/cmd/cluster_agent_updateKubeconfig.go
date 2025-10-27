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
	cluster_agent_updateKubeconfigCmd.Flags().StringP("cache-mode", "c", "", "Mode to use for caching the token. Allowed values: keyring-filesystem-fallback, force-keyring, force-filesystem, no")
	cluster_agent_updateKubeconfigCmd.Flags().Bool("check-revoked", false, "Check if a cached token is revoked. Requires an API call to GitLab, which adds latency every time a cached token is accessed.")
	cluster_agent_updateKubeconfigCmd.Flags().String("kubeconfig", "", "Use a particular kubeconfig file.")
	cluster_agent_updateKubeconfigCmd.Flags().String("token-expiry-duration", "", "Duration for generated token's validity. Minimum is 1 day. Expires at end of day, and ignores time.")
	cluster_agent_updateKubeconfigCmd.Flags().BoolP("use-context", "u", false, "Use as default context.")
	cluster_agent_updateKubeconfigCmd.MarkFlagRequired("agent")
	cluster_agentCmd.AddCommand(cluster_agent_updateKubeconfigCmd)

	// TODO flag completion
	carapace.Gen(cluster_agent_updateKubeconfigCmd).FlagCompletion(carapace.ActionMap{
		"cache-mode": carapace.ActionValues("keyring-filesystem-fallback", "force-keyring", "force-filesystem", "no"),
		"kubeconfig": carapace.ActionFiles(),
	})
}
