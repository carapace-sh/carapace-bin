package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List available devices, with an optional property as the filter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicesCmd).Standalone()
	rootCmd.AddCommand(devicesCmd)
	carapace.Gen(devicesCmd).PositionalCompletion(
		carapace.ActionStyledValues(
			"Paired", style.Yellow,
			"Bonded", style.Magenta,
			"Trusted", style.Green,
			"Connected", style.Blue,
		),
	)
}
