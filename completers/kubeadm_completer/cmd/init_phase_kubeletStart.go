package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_kubeletStartCmd = &cobra.Command{
	Use:   "kubelet-start",
	Short: "Write kubelet settings and (re)start the kubelet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeletStartCmd).Standalone()

	init_phase_kubeletStartCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_kubeletStartCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	init_phase_kubeletStartCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_kubeletStartCmd.Flags().String("image-repository", "", "Choose a container registry to pull control plane images from")
	init_phase_kubeletStartCmd.Flags().String("node-name", "", "Specify the node name.")
	init_phase_kubeletStartCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\". For example, \"kube-apiserver0+merge.yaml\" or just \"etcd.json\". \"target\" can be one of \"kube-apiserver\", \"kube-controller-manager\", \"kube-scheduler\", \"etcd\", \"kubeletconfiguration\", \"corednsdeployment\". \"patchtype\" can be one of \"strategic\", \"merge\" or \"json\" and they match the patch formats supported by kubectl. The default \"patchtype\" is \"strategic\". \"extension\" must be either \"json\" or \"yaml\". \"suffix\" is an optional string that can be used to determine which patches are applied first alpha-numerically.")
	init_phaseCmd.AddCommand(init_phase_kubeletStartCmd)

	carapace.Gen(init_phase_kubeletStartCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
