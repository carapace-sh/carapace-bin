package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var custom_devicesCmd = &cobra.Command{
	Use:   "custom-devices",
	Short: "List, reset, add and delete custom devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(custom_devicesCmd).Standalone()

	custom_devicesCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(custom_devicesCmd)
}
