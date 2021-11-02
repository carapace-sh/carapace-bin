package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_firewall_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently delete a cloud firewall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_firewall_deleteCmd).Standalone()
	compute_firewall_deleteCmd.Flags().BoolP("force", "f", false, "Delete firewall without confirmation prompt")
	compute_firewallCmd.AddCommand(compute_firewall_deleteCmd)
}
