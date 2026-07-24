package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_derp_mapCmd = &cobra.Command{
	Use:   "derp-map",
	Short: "Print DERP map",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_derp_mapCmd).Standalone()

	debugCmd.AddCommand(debug_derp_mapCmd)
}
