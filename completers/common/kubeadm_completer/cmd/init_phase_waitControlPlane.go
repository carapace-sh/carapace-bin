package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_waitControlPlaneCmd = &cobra.Command{
	Use:   "wait-control-plane",
	Short: "Wait for the control plane to start",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_waitControlPlaneCmd).Standalone()

	init_phaseCmd.AddCommand(init_phase_waitControlPlaneCmd)
}
