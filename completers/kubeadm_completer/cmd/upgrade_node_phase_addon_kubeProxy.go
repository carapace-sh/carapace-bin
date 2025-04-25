package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_node_phase_addon_kubeProxyCmd = &cobra.Command{
	Use:   "kube-proxy",
	Short: "Upgrade the kube-proxy addon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phase_addon_kubeProxyCmd).Standalone()

	upgrade_node_phase_addon_kubeProxyCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_node_phase_addon_kubeProxyCmd.Flags().Bool("dry-run", false, "Do not change any state, just output the actions that would be performed.")
	upgrade_node_phase_addon_kubeProxyCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_node_phase_addonCmd.AddCommand(upgrade_node_phase_addon_kubeProxyCmd)

	carapace.Gen(upgrade_node_phase_addon_kubeProxyCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
