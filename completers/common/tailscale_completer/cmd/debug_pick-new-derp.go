package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_pick_new_derpCmd = &cobra.Command{
	Use:   "pick-new-derp",
	Short: "Switch to some other random DERP home region for a short time",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_pick_new_derpCmd).Standalone()

	debugCmd.AddCommand(debug_pick_new_derpCmd)
}
