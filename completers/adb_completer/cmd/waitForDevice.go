package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waitForDeviceCmd = &cobra.Command{
	Use:   "wait-for-device",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForDeviceCmd).Standalone()

	rootCmd.AddCommand(waitForDeviceCmd)
}
