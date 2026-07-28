package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_controlKnobsCmd = &cobra.Command{
	Use:   "control-knobs",
	Short: "See current control knobs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_controlKnobsCmd).Standalone()

	debugCmd.AddCommand(debug_controlKnobsCmd)
}
