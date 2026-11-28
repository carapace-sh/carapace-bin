package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dfuFlashCmd = &cobra.Command{
	Use:   "dfu-flash",
	Short: "Flash the DFU binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dfuFlashCmd).Standalone()
	dfuFlashCmd.Flags().String("path", "", "Specify path to DFU device")
	rootCmd.AddCommand(dfuFlashCmd)
}
