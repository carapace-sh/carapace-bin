package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_getTokenCmd = &cobra.Command{
	Use:   "get-token [flags]",
	Short: "Create and return a k8s_proxy-scoped personal access token to authenticate with a GitLab Agents for Kubernetes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_getTokenCmd).Standalone()

	cluster_agent_getTokenCmd.Flags().StringP("agent", "a", "", "The numerical Agent ID to connect to.")
	cluster_agent_getTokenCmd.Flags().StringP("cache-mode", "c", "", "Mode to use for caching the token. Allowed values: keyring-filesystem-fallback, force-keyring, force-filesystem, no")
	cluster_agent_getTokenCmd.Flags().Bool("check-revoked", false, "Check if a cached token is revoked. This requires an API call to GitLab which adds latency every time a cached token is accessed.")
	cluster_agent_getTokenCmd.Flags().String("token-expiry-duration", "", "Duration for how long the generated tokens should be valid for. Minimum is 1 day and the effective expiry is always at the end of the day, the time is ignored.")
	cluster_agent_getTokenCmd.MarkFlagRequired("agent")
	cluster_agentCmd.AddCommand(cluster_agent_getTokenCmd)

	// TODO flag completion
	carapace.Gen(cluster_agent_getTokenCmd).FlagCompletion(carapace.ActionMap{
		"cache-mode": carapace.ActionValues("keyring-filesystem-fallback", "force-keyring", "force-filesystem", "no"),
	})
}
