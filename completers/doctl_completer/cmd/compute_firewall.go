package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_firewallCmd = &cobra.Command{
	Use:   "firewall",
	Short: "Display commands to manage cloud firewalls",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_firewallCmd).Standalone()
	computeCmd.AddCommand(compute_firewallCmd)
}
