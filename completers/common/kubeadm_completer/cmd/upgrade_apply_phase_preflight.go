package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_apply_phase_preflightCmd = &cobra.Command{
	Use:   "preflight",
	Short: "Run preflight checks before upgrade",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_apply_phase_preflightCmd).Standalone()

	upgrade_apply_phase_preflightCmd.Flags().Bool("allow-experimental-upgrades", false, "Show unstable versions of Kubernetes as an upgrade alternative and allow upgrading to an alpha/beta/release candidate versions of Kubernetes.")
	upgrade_apply_phase_preflightCmd.Flags().Bool("allow-release-candidate-upgrades", false, "Show release candidate versions of Kubernetes as an upgrade alternative and allow upgrading to a release candidate versions of Kubernetes.")
	upgrade_apply_phase_preflightCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_apply_phase_preflightCmd.Flags().Bool("dry-run", false, "Do not change any state, just output what actions would be performed.")
	upgrade_apply_phase_preflightCmd.Flags().BoolP("force", "f", false, "Force upgrading although some requirements might not be met. This also implies non-interactive mode.")
	upgrade_apply_phase_preflightCmd.Flags().StringSlice("ignore-preflight-errors", nil, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	upgrade_apply_phase_preflightCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	upgrade_apply_phase_preflightCmd.Flags().BoolP("yes", "y", false, "Perform the upgrade and do not prompt for confirmation (non-interactive mode).")
	upgrade_apply_phaseCmd.AddCommand(upgrade_apply_phase_preflightCmd)

	carapace.Gen(upgrade_apply_phase_preflightCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
