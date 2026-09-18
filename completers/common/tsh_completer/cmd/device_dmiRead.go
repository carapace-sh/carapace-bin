package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var device_dmiReadCmd = &cobra.Command{
	Use:    "dmi-read",
	Short:  "Read device DMI information.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_dmiReadCmd).Standalone()

	deviceCmd.AddCommand(device_dmiReadCmd)
}
