package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_derp_set_on_demandCmd = &cobra.Command{
	Use:   "derp-set-on-demand",
	Short: "Enable DERP on-demand mode (breaks reachability)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_derp_set_on_demandCmd).Standalone()

	debugCmd.AddCommand(debug_derp_set_on_demandCmd)
}
