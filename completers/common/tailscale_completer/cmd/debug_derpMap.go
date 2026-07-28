package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_derpMapCmd = &cobra.Command{
	Use:   "derp-map",
	Short: "Print DERP map",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_derpMapCmd).Standalone()

	debugCmd.AddCommand(debug_derpMapCmd)
}
