package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cluster_agent_checkManifestUsageCmd = &cobra.Command{
	Use:     "check-manifest-usage [flags]",
	Short:   "Find agents using deprecated GitOps manifest settings. (EXPERIMENTAL)",
	Aliases: []string{"check_manifest_usage"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_checkManifestUsageCmd).Standalone()

	cluster_agent_checkManifestUsageCmd.Flags().StringP("agent-page", "a", "1", "Page number for agents.")
	cluster_agent_checkManifestUsageCmd.Flags().StringP("agent-per-page", "A", "30", "Number of agents to list per page.")
	cluster_agent_checkManifestUsageCmd.Flags().StringP("group", "g", "", "Group ID to check.")
	cluster_agent_checkManifestUsageCmd.Flags().StringP("page", "p", "1", "Page number for projects.")
	cluster_agent_checkManifestUsageCmd.Flags().StringP("per-page", "P", "30", "Number of projects to list per page.")
	cluster_agent_checkManifestUsageCmd.Flags().BoolP("recursive", "r", false, "Recursively check subgroups.")
	cluster_agent_checkManifestUsageCmd.MarkFlagRequired("group")
	cluster_agentCmd.AddCommand(cluster_agent_checkManifestUsageCmd)

	carapace.Gen(cluster_agent_checkManifestUsageCmd).FlagCompletion(carapace.ActionMap{
		"group": action.ActionGroups(cluster_agent_checkManifestUsageCmd),
	})
}
