package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cluster_agent_tokenCache_listCmd = &cobra.Command{
	Use:   "list [flags]",
	Short: "List cached GitLab Agent tokens.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_tokenCache_listCmd).Standalone()

	cluster_agent_tokenCache_listCmd.Flags().StringSlice("agent", nil, "Filter by specific agent IDs.")
	cluster_agent_tokenCache_listCmd.Flags().Bool("filesystem", false, "Include tokens from filesystem cache.")
	cluster_agent_tokenCache_listCmd.Flags().Bool("keyring", false, "Include tokens from keyring cache.")
	cluster_agent_tokenCache_listCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	cluster_agent_tokenCacheCmd.AddCommand(cluster_agent_tokenCache_listCmd)

	carapace.Gen(cluster_agent_tokenCache_listCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(cluster_agent_tokenCache_listCmd),
	})
}
