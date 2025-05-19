package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_kubeletFinalize_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Run all kubelet-finalize phases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeletFinalize_allCmd).Standalone()

	init_phase_kubeletFinalize_allCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_kubeletFinalize_allCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_kubeletFinalize_allCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_kubeletFinalizeCmd.AddCommand(init_phase_kubeletFinalize_allCmd)

	carapace.Gen(init_phase_kubeletFinalize_allCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
