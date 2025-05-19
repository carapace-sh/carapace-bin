package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_apply_phase_bootstrapTokenCmd = &cobra.Command{
	Use:   "bootstrap-token",
	Short: "Configures bootstrap token and cluster-info RBAC rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_apply_phase_bootstrapTokenCmd).Standalone()

	upgrade_apply_phase_bootstrapTokenCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_apply_phase_bootstrapTokenCmd.Flags().Bool("dry-run", false, "Do not change any state, just output what actions would be performed.")
	upgrade_apply_phase_bootstrapTokenCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_apply_phaseCmd.AddCommand(upgrade_apply_phase_bootstrapTokenCmd)

	carapace.Gen(upgrade_apply_phase_bootstrapTokenCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
