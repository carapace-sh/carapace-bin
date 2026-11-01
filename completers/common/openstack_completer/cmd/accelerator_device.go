package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accelerator_deviceCmd = &cobra.Command{
	Use:   "device",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_deviceCmd).Standalone()

	acceleratorCmd.AddCommand(accelerator_deviceCmd)
}
