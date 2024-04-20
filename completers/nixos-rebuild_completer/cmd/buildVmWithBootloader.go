package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildVmWithBootloaderCmd = &cobra.Command{
	Use:   "build-vm-with-bootloader",
	Short: "Like build-vm, but boots with the regular bootloader of your configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildVmWithBootloaderCmd).Standalone()
	rootCmd.AddCommand(buildVmWithBootloaderCmd)
}
