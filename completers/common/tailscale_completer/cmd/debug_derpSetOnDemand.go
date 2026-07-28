package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_derpSetOnDemandCmd = &cobra.Command{
	Use:   "derp-set-on-demand",
	Short: "Enable DERP on-demand mode (breaks reachability)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_derpSetOnDemandCmd).Standalone()

	debugCmd.AddCommand(debug_derpSetOnDemandCmd)
}
