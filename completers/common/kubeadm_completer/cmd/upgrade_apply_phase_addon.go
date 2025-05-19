package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_apply_phase_addonCmd = &cobra.Command{
	Use:   "addon",
	Short: "Upgrade the default kubeadm addons",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_apply_phase_addonCmd).Standalone()

	upgrade_apply_phaseCmd.AddCommand(upgrade_apply_phase_addonCmd)
}
