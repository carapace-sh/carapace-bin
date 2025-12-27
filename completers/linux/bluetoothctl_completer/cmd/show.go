package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [controller]",
	Short: "Show controller information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()
	rootCmd.AddCommand(showCmd)
	carapace.Gen(showCmd).PositionalCompletion(
		bluetoothctl.ActionControllers(),
	)
}
