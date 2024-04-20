package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_vm_with_bootloaderCmd = &cobra.Command{
	Use:   "build-vm-with-bootloader",
	Short: "Like build-vm, but boots with the regular bootloader of your configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_vm_with_bootloaderCmd).Standalone()
	rootCmd.AddCommand(build_vm_with_bootloaderCmd)
}
