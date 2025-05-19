package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_apply_phase_uploadConfig_kubeadmCmd = &cobra.Command{
	Use:   "kubeadm",
	Short: "Upload the kubeadm ClusterConfiguration to a ConfigMap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_apply_phase_uploadConfig_kubeadmCmd).Standalone()

	upgrade_apply_phase_uploadConfig_kubeadmCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_apply_phase_uploadConfig_kubeadmCmd.Flags().Bool("dry-run", false, "Do not change any state, just output what actions would be performed.")
	upgrade_apply_phase_uploadConfig_kubeadmCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_apply_phase_uploadConfigCmd.AddCommand(upgrade_apply_phase_uploadConfig_kubeadmCmd)

	carapace.Gen(upgrade_apply_phase_uploadConfig_kubeadmCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
