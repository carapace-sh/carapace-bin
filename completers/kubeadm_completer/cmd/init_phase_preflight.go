package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/kubeadm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var init_phase_preflightCmd = &cobra.Command{
	Use:   "preflight",
	Short: "Run pre-flight checks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_preflightCmd).Standalone()

	init_phase_preflightCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_preflightCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	init_phase_preflightCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_preflightCmd.Flags().StringSlice("ignore-preflight-errors", []string{}, "A list of checks whose errors will be shown as warnings. Example: 'IsPrivilegedUser,Swap'. Value 'all' ignores errors from all checks.")
	init_phase_preflightCmd.Flags().String("image-repository", "", "Choose a container registry to pull control plane images from")
	init_phaseCmd.AddCommand(init_phase_preflightCmd)

	carapace.Gen(init_phase_preflightCmd).FlagCompletion(carapace.ActionMap{
		"config":                  carapace.ActionFiles(),
		"ignore-preflight-errors": action.ActionChecks().UniqueList(","),
	})
}
