package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_showJoinCommandCmd = &cobra.Command{
	Use:   "show-join-command",
	Short: "Show the join command for control-plane and worker node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_showJoinCommandCmd).Standalone()

	init_phaseCmd.AddCommand(init_phase_showJoinCommandCmd)
}
