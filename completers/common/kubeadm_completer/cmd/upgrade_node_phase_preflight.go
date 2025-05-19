package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgrade_node_phase_preflightCmd = &cobra.Command{
	Use:   "preflight",
	Short: "Run upgrade node pre-flight checks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phase_preflightCmd).Standalone()

	upgrade_node_phase_preflightCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	upgrade_node_phase_preflightCmd.Flags().StringSlice("ignore-preflight-errors", nil, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	upgrade_node_phaseCmd.AddCommand(upgrade_node_phase_preflightCmd)

	carapace.Gen(upgrade_node_phase_preflightCmd).FlagCompletion(carapace.ActionMap{
		"ignore-preflight-errors": action.ActionChecks().UniqueList(","),
	})
}
