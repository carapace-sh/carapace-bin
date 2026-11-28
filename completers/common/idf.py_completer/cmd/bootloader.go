package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootloaderCmd = &cobra.Command{
	Use:   "bootloader",
	Short: "Build only bootloader",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootloaderCmd).Standalone()
	rootCmd.AddCommand(bootloaderCmd)
}
