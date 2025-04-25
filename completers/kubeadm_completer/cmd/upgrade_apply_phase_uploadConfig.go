package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_apply_phase_uploadConfigCmd = &cobra.Command{
	Use:     "upload-config",
	Short:   "Upload the kubeadm and kubelet configurations to ConfigMaps",
	Aliases: []string{"uploadconfig"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_apply_phase_uploadConfigCmd).Standalone()

	upgrade_apply_phaseCmd.AddCommand(upgrade_apply_phase_uploadConfigCmd)
}
