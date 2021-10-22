package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_floatingIpCmd = &cobra.Command{
	Use:   "floating-ip",
	Short: "Display commands to manage floating IP addresses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_floatingIpCmd).Standalone()
	computeCmd.AddCommand(compute_floatingIpCmd)
}
