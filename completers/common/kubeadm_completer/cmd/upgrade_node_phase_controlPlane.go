package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_node_phase_controlPlaneCmd = &cobra.Command{
	Use:   "control-plane",
	Short: "Upgrade the control plane instance deployed on this node, if any",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phase_controlPlaneCmd).Standalone()

	upgrade_node_phase_controlPlaneCmd.Flags().Bool("certificate-renewal", false, "Perform the renewal of certificates used by component changed during upgrades.")
	upgrade_node_phase_controlPlaneCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_node_phase_controlPlaneCmd.Flags().Bool("dry-run", false, "Do not change any state, just output the actions that would be performed.")
	upgrade_node_phase_controlPlaneCmd.Flags().Bool("etcd-upgrade", false, "Perform the upgrade of etcd.")
	upgrade_node_phase_controlPlaneCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_node_phase_controlPlaneCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\". For example, \"kube-apiserver0+merge.yaml\" or just \"etcd.json\". \"target\" can be one of \"kube-apiserver\", \"kube-controller-manager\", \"kube-scheduler\", \"etcd\", \"kubeletconfiguration\", \"corednsdeployment\". \"patchtype\" can be one of \"strategic\", \"merge\" or \"json\" and they match the patch formats supported by kubectl. The default \"patchtype\" is \"strategic\". \"extension\" must be either \"json\" or \"yaml\". \"suffix\" is an optional string that can be used to determine which patches are applied first alpha-numerically.")
	upgrade_node_phaseCmd.AddCommand(upgrade_node_phase_controlPlaneCmd)

	carapace.Gen(upgrade_node_phase_controlPlaneCmd).FlagCompletion(carapace.ActionMap{
		"kubeconfig": carapace.ActionFiles(),
		"patches":    carapace.ActionDirectories(),
	})
}
