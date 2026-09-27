package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quota_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete configured quota for a project and revert to defaults.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quota_deleteCmd).Standalone()

	quota_deleteCmd.Flags().Bool("all", false, "Delete project quotas for all services (default)")
	quota_deleteCmd.Flags().Bool("compute", false, "Delete compute quotas for the project (including network quotas when using nova-network)")
	quota_deleteCmd.Flags().Bool("network", false, "Delete network quotas for the project")
	quota_deleteCmd.Flags().Bool("volume", false, "Delete volume quotas for the project")
	quotaCmd.AddCommand(quota_deleteCmd)
}
