package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_uploadConfig_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Upload all configuration to a config map",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_uploadConfig_allCmd).Standalone()

	init_phase_uploadConfig_allCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_uploadConfig_allCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	init_phase_uploadConfig_allCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_uploadConfig_allCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_uploadConfigCmd.AddCommand(init_phase_uploadConfig_allCmd)

	carapace.Gen(init_phase_uploadConfig_allCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
