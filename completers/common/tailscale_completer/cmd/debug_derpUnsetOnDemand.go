package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_derpUnsetOnDemandCmd = &cobra.Command{
	Use:   "derp-unset-on-demand",
	Short: "Disable DERP on-demand mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_derpUnsetOnDemandCmd).Standalone()

	debugCmd.AddCommand(debug_derpUnsetOnDemandCmd)
}
