package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_derp_unset_on_demandCmd = &cobra.Command{
	Use:   "derp-unset-on-demand",
	Short: "Disable DERP on-demand mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_derp_unset_on_demandCmd).Standalone()

	debugCmd.AddCommand(debug_derp_unset_on_demandCmd)
}
