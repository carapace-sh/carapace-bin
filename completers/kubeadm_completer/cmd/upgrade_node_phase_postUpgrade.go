package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_node_phase_postUpgradeCmd = &cobra.Command{
	Use:   "post-upgrade",
	Short: "Run post upgrade tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phase_postUpgradeCmd).Standalone()

	upgrade_node_phase_postUpgradeCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_node_phase_postUpgradeCmd.Flags().Bool("dry-run", false, "Do not change any state, just output the actions that would be performed.")
	upgrade_node_phase_postUpgradeCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_node_phaseCmd.AddCommand(upgrade_node_phase_postUpgradeCmd)

	carapace.Gen(upgrade_node_phase_postUpgradeCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
