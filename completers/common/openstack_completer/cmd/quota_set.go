package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quota_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set quotas for project or class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quota_setCmd).Standalone()

	quota_setCmd.Flags().String("backup-gigabytes", "", "New value for the backup-gigabytes quota")
	quota_setCmd.Flags().String("backups", "", "New value for the backups quota")
	quota_setCmd.Flags().Bool("check-limit", false, "==SUPPRESS==")
	quota_setCmd.Flags().Bool("class", false, "**Deprecated** Set quotas for <class>.")
	quota_setCmd.Flags().String("cores", "", "New value for the cores quota")
	quota_setCmd.Flags().Bool("default", false, "Set default quotas for <project>")
	quota_setCmd.Flags().String("floating-ips", "", "New value for the floating-ips quota")
	quota_setCmd.Flags().Bool("force", false, "Force quota update (only supported by compute and network)")
	quota_setCmd.Flags().String("gigabytes", "", "New value for the gigabytes quota")
	quota_setCmd.Flags().String("groups", "", "New value for the groups quota")
	quota_setCmd.Flags().String("injected-file-size", "", "New value for the injected-file-size quota")
	quota_setCmd.Flags().String("injected-files", "", "New value for the injected-files quota")
	quota_setCmd.Flags().String("injected-path-size", "", "New value for the injected-path-size quota")
	quota_setCmd.Flags().String("instances", "", "New value for the instances quota")
	quota_setCmd.Flags().String("key-pairs", "", "New value for the key-pairs quota")
	quota_setCmd.Flags().String("networks", "", "New value for the networks quota")
	quota_setCmd.Flags().Bool("no-force", false, "Do not force quota update (only supported by compute and network) (default)")
	quota_setCmd.Flags().String("per-volume-gigabytes", "", "New value for the per-volume-gigabytes quota")
	quota_setCmd.Flags().String("ports", "", "New value for the ports quota")
	quota_setCmd.Flags().String("properties", "", "New value for the properties quota")
	quota_setCmd.Flags().String("ram", "", "New value for the ram quota")
	quota_setCmd.Flags().String("rbac-policies", "", "New value for the rbac-policies quota")
	quota_setCmd.Flags().String("router-routes", "", "New value for the router-routes quota")
	quota_setCmd.Flags().String("routers", "", "New value for the routers quota")
	quota_setCmd.Flags().String("secgroup-rules", "", "New value for the secgroup-rules quota")
	quota_setCmd.Flags().String("secgroups", "", "New value for the secgroups quota")
	quota_setCmd.Flags().String("server-group-members", "", "New value for the server-group-members quota")
	quota_setCmd.Flags().String("server-groups", "", "New value for the server-groups quota")
	quota_setCmd.Flags().String("snapshots", "", "New value for the snapshots quota")
	quota_setCmd.Flags().String("subnetpools", "", "New value for the subnetpools quota")
	quota_setCmd.Flags().String("subnets", "", "New value for the subnets quota")
	quota_setCmd.Flags().String("volume-type", "", "Set quotas for a specific <volume-type>")
	quota_setCmd.Flags().String("volumes", "", "New value for the volumes quota")
	quotaCmd.AddCommand(quota_setCmd)
}
