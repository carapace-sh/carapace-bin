package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available devices, device types, runtimes, or device pairs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	listCmd.Flags().BoolP("json", "j", false, "Produce JSON output")
	rootCmd.AddCommand(listCmd)
	carapace.Gen(listCmd).PositionalCompletion(
		carapace.ActionValues("devices", "devicetypes", "runtimes", "pairs"),
	)
}
