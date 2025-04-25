package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_apply_phase_uploadConfig_kubeletCmd = &cobra.Command{
	Use:   "kubelet",
	Short: "Upload the kubelet configuration to a ConfigMap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_apply_phase_uploadConfig_kubeletCmd).Standalone()

	upgrade_apply_phase_uploadConfig_kubeletCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_apply_phase_uploadConfig_kubeletCmd.Flags().Bool("dry-run", false, "Do not change any state, just output what actions would be performed.")
	upgrade_apply_phase_uploadConfig_kubeletCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_apply_phase_uploadConfigCmd.AddCommand(upgrade_apply_phase_uploadConfig_kubeletCmd)

	carapace.Gen(upgrade_apply_phase_uploadConfig_kubeletCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
