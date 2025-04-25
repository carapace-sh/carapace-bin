package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_kubeletFinalize_enableClientCertRotationCmd = &cobra.Command{
	Use:   "enable-client-cert-rotation",
	Short: "Enable kubelet client certificate rotation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeletFinalize_enableClientCertRotationCmd).Standalone()

	init_phase_kubeletFinalize_enableClientCertRotationCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_kubeletFinalize_enableClientCertRotationCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_kubeletFinalize_enableClientCertRotationCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_kubeletFinalizeCmd.AddCommand(init_phase_kubeletFinalize_enableClientCertRotationCmd)

	carapace.Gen(init_phase_kubeletFinalize_enableClientCertRotationCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
