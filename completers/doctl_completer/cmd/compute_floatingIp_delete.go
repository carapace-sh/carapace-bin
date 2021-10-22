package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_floatingIp_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Permanently delete a floating IP address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_floatingIp_deleteCmd).Standalone()
	compute_floatingIp_deleteCmd.Flags().BoolP("force", "f", false, "Force floating IP delete")
	compute_floatingIpCmd.AddCommand(compute_floatingIp_deleteCmd)
}
