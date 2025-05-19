package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgrade_apply_phaseCmd = &cobra.Command{
	Use:   "phase",
	Short: "Use this command to invoke single phase of the \"apply\" workflow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_apply_phaseCmd).Standalone()

	upgrade_applyCmd.AddCommand(upgrade_apply_phaseCmd)
}
