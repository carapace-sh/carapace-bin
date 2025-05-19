package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reset_phase_preflightCmd = &cobra.Command{
	Use:     "preflight",
	Short:   "Run reset pre-flight checks",
	Aliases: []string{"pre-flight"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reset_phase_preflightCmd).Standalone()

	reset_phase_preflightCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	reset_phase_preflightCmd.Flags().BoolP("force", "f", false, "Reset the node without prompting for confirmation.")
	reset_phase_preflightCmd.Flags().StringSlice("ignore-preflight-errors", nil, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	reset_phaseCmd.AddCommand(reset_phase_preflightCmd)

	carapace.Gen(reset_phase_preflightCmd).FlagCompletion(carapace.ActionMap{
		"ignore-preflight-errors": action.ActionChecks().UniqueList(","),
	})
}
