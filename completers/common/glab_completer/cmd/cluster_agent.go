package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cluster_agentCmd = &cobra.Command{
	Use:   "agent <command> [flags]",
	Short: "Manage GitLab Agents for Kubernetes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agentCmd).Standalone()

	cluster_agentCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	clusterCmd.AddCommand(cluster_agentCmd)

	carapace.Gen(cluster_agentCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(cluster_agentCmd),
	})
}
