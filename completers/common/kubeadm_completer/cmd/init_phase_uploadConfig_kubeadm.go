package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_uploadConfig_kubeadmCmd = &cobra.Command{
	Use:   "kubeadm",
	Short: "Upload the kubeadm ClusterConfiguration to a ConfigMap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_uploadConfig_kubeadmCmd).Standalone()

	init_phase_uploadConfig_kubeadmCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_uploadConfig_kubeadmCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	init_phase_uploadConfig_kubeadmCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_uploadConfig_kubeadmCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_uploadConfigCmd.AddCommand(init_phase_uploadConfig_kubeadmCmd)

	carapace.Gen(init_phase_uploadConfig_kubeadmCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
