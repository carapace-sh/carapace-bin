package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var join_phase_kubeletStartCmd = &cobra.Command{
	Use:   "kubelet-start [api-server-endpoint]",
	Short: "Write kubelet settings, certificates and (re)start the kubelet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_kubeletStartCmd).Standalone()

	join_phase_kubeletStartCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	join_phase_kubeletStartCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	join_phase_kubeletStartCmd.Flags().String("discovery-file", "", "For file-based discovery, a file or URL from which to load cluster information.")
	join_phase_kubeletStartCmd.Flags().String("discovery-token", "", "For token-based discovery, the token used to validate cluster information fetched from the API server.")
	join_phase_kubeletStartCmd.Flags().StringSlice("discovery-token-ca-cert-hash", []string{}, "For token-based discovery, validate that the root CA public key matches this hash (format: \"<type>:<value>\").")
	join_phase_kubeletStartCmd.Flags().Bool("discovery-token-unsafe-skip-ca-verification", false, "For token-based discovery, allow joining without --discovery-token-ca-cert-hash pinning.")
	join_phase_kubeletStartCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	join_phase_kubeletStartCmd.Flags().String("node-name", "", "Specify the node name.")
	join_phase_kubeletStartCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\". For example, \"kube-apiserver0+merge.yaml\" or just \"etcd.json\". \"target\" can be one of \"kube-apiserver\", \"kube-controller-manager\", \"kube-scheduler\", \"etcd\", \"kubeletconfiguration\", \"corednsdeployment\". \"patchtype\" can be one of \"strategic\", \"merge\" or \"json\" and they match the patch formats supported by kubectl. The default \"patchtype\" is \"strategic\". \"extension\" must be either \"json\" or \"yaml\". \"suffix\" is an optional string that can be used to determine which patches are applied first alpha-numerically.")
	join_phase_kubeletStartCmd.Flags().String("tls-bootstrap-token", "", "Specify the token used to temporarily authenticate with the Kubernetes Control Plane while joining the node.")
	join_phase_kubeletStartCmd.Flags().String("token", "", "Use this token for both discovery-token and tls-bootstrap-token when those values are not provided.")
	join_phaseCmd.AddCommand(join_phase_kubeletStartCmd)

	carapace.Gen(join_phase_kubeletStartCmd).FlagCompletion(carapace.ActionMap{
		"config":         carapace.ActionFiles(),
		"discovery-file": carapace.ActionFiles(),
	})
}
