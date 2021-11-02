package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_floatingIpActionCmd = &cobra.Command{
	Use:   "floating-ip-action",
	Short: "Display commands to associate floating IP addresses with Droplets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_floatingIpActionCmd).Standalone()
	computeCmd.AddCommand(compute_floatingIpActionCmd)
}
