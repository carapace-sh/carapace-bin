package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_node_phase_addon_corednsCmd = &cobra.Command{
	Use:   "coredns",
	Short: "Upgrade the CoreDNS addon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phase_addon_corednsCmd).Standalone()

	upgrade_node_phase_addon_corednsCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_node_phase_addon_corednsCmd.Flags().Bool("dry-run", false, "Do not change any state, just output the actions that would be performed.")
	upgrade_node_phase_addon_corednsCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_node_phase_addon_corednsCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\". For example, \"kube-apiserver0+merge.yaml\" or just \"etcd.json\". \"target\" can be one of \"kube-apiserver\", \"kube-controller-manager\", \"kube-scheduler\", \"etcd\", \"kubeletconfiguration\", \"corednsdeployment\". \"patchtype\" can be one of \"strategic\", \"merge\" or \"json\" and they match the patch formats supported by kubectl. The default \"patchtype\" is \"strategic\". \"extension\" must be either \"json\" or \"yaml\". \"suffix\" is an optional string that can be used to determine which patches are applied first alpha-numerically.")
	upgrade_node_phase_addonCmd.AddCommand(upgrade_node_phase_addon_corednsCmd)
}
