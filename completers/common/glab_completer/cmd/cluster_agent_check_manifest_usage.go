package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cluster_agent_check_manifest_usageCmd = &cobra.Command{
	Use:   "check_manifest_usage [flags]",
	Short: "Check agent configuration files for built-in GitOps manifests usage. (EXPERIMENTAL.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cluster_agent_check_manifest_usageCmd).Standalone()

	cluster_agent_check_manifest_usageCmd.Flags().StringP("agent-page", "a", "", "Page number for projects.")
	cluster_agent_check_manifest_usageCmd.Flags().StringP("agent-per-page", "A", "", "Number of projects to list per page.")
	cluster_agent_check_manifest_usageCmd.Flags().StringP("group", "g", "", "Group ID to check.")
	cluster_agent_check_manifest_usageCmd.Flags().StringP("page", "p", "", "Page number for projects.")
	cluster_agent_check_manifest_usageCmd.Flags().StringP("per-page", "P", "", "Number of projects to list per page.")
	cluster_agent_check_manifest_usageCmd.Flags().BoolP("recursive", "r", false, "Recursively check subgroups.")
	cluster_agent_check_manifest_usageCmd.MarkFlagRequired("group")
	cluster_agentCmd.AddCommand(cluster_agent_check_manifest_usageCmd)

	// TODO flag completion
}
