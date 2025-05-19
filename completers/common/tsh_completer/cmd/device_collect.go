package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var device_collectCmd = &cobra.Command{
	Use:    "collect",
	Short:  "Simulate enroll/authn device data collection",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_collectCmd).Standalone()

	deviceCmd.AddCommand(device_collectCmd)
}
