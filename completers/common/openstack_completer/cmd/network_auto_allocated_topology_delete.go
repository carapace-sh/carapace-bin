package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_auto_allocated_topology_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete auto allocated topology for project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_auto_allocated_topology_deleteCmd).Standalone()

	network_auto_allocated_topology_deleteCmd.Flags().String("project", "", "Delete auto allocated topology for a given project.")
	network_auto_allocated_topology_deleteCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_auto_allocated_topologyCmd.AddCommand(network_auto_allocated_topology_deleteCmd)
}
