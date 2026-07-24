package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_break_derp_connsCmd = &cobra.Command{
	Use:   "break-derp-conns",
	Short: "Break any open DERP connections from the daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_break_derp_connsCmd).Standalone()

	debugCmd.AddCommand(debug_break_derp_connsCmd)
}
