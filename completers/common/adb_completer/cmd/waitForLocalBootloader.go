package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waitForLocalBootloaderCmd = &cobra.Command{
	Use:   "wait-for-local-bootloader",
	Short: "wait for device to be in given state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitForLocalBootloaderCmd).Standalone()

	rootCmd.AddCommand(waitForLocalBootloaderCmd)
}
