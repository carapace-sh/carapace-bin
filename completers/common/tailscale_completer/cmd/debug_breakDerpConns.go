package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_breakDerpConnsCmd = &cobra.Command{
	Use:   "break-derp-conns",
	Short: "Break any open DERP connections from the daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_breakDerpConnsCmd).Standalone()

	debugCmd.AddCommand(debug_breakDerpConnsCmd)
}
