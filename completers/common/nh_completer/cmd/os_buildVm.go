package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_buildVmCmd = &cobra.Command{
	Use:   "build-vm",
	Short: "Build a NixOS VM image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_buildVmCmd).Standalone()

	addOsRebuildFlags(os_buildVmCmd)
	os_buildVmCmd.Flags().BoolP("with-bootloader", "B", false, "Build with bootloader")
	os_buildVmCmd.Flags().BoolP("run", "r", false, "Run the VM immediately after building")

	osCmd.AddCommand(os_buildVmCmd)
}
