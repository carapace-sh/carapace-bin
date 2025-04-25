package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_node_phase_addonCmd = &cobra.Command{
	Use:   "addon",
	Short: "Upgrade the default kubeadm addons",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phase_addonCmd).Standalone()

	upgrade_node_phaseCmd.AddCommand(upgrade_node_phase_addonCmd)
}
