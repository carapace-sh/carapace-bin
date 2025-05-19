package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var device_keygetCmd = &cobra.Command{
	Use:    "keyget",
	Short:  "Get information about the device key",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_keygetCmd).Standalone()

	deviceCmd.AddCommand(device_keygetCmd)
}
