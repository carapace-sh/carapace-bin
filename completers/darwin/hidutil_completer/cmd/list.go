package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List HID Event System services and devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().StringP("matching", "m", "", "Set matching services/devices (JSON dictionary or device string)")
	listCmd.Flags().BoolP("ndjson", "n", false, "Print service/device information in ndjson format")
	rootCmd.AddCommand(listCmd)
}
