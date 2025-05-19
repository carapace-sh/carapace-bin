package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select <controller>",
	Short: "Select default controller",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectCmd).Standalone()
	rootCmd.AddCommand(selectCmd)
	carapace.Gen(selectCmd).PositionalCompletion(
		bluetoothctl.ActionControllers(),
	)
}
