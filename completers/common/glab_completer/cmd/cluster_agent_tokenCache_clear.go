package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cluster_agent_tokenCache_clearCmd = &cobra.Command{
	Use:   "clear [flags]",
	Short: "Clear cached GitLab Agent tokens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_tokenCache_clearCmd).Standalone()

	cluster_agent_tokenCache_clearCmd.Flags().StringSlice("agent", nil, "Clear tokens for specific agent IDs only.")
	cluster_agent_tokenCache_clearCmd.Flags().Bool("filesystem", false, "Clear tokens from filesystem cache.")
	cluster_agent_tokenCache_clearCmd.Flags().Bool("keyring", false, "Clear tokens from keyring cache.")
	cluster_agent_tokenCache_clearCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	cluster_agent_tokenCache_clearCmd.Flags().Bool("revoke", false, "Revoke tokens on GitLab server before clearing cache.")
	cluster_agent_tokenCacheCmd.AddCommand(cluster_agent_tokenCache_clearCmd)

	carapace.Gen(cluster_agent_tokenCache_clearCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(cluster_agent_tokenCache_clearCmd),
	})
}
