package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var join_phase_kubeletWaitBootstrapCmd = &cobra.Command{
	Use:   "kubelet-wait-bootstrap",
	Short: "Wait for the kubelet to bootstrap itself",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_kubeletWaitBootstrapCmd).Standalone()

	join_phase_kubeletWaitBootstrapCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	join_phase_kubeletWaitBootstrapCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	join_phase_kubeletWaitBootstrapCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	join_phaseCmd.AddCommand(join_phase_kubeletWaitBootstrapCmd)

	carapace.Gen(join_phase_kubeletWaitBootstrapCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"cri-socket": carapace.ActionFiles(),
	})
}
