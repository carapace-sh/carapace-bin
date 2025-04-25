package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var join_phase_waitControlPlaneCmd = &cobra.Command{
	Use:   "wait-control-plane",
	Short: "EXPERIMENTAL: Wait for the control plane to start",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_waitControlPlaneCmd).Standalone()

	join_phaseCmd.AddCommand(join_phase_waitControlPlaneCmd)
}
