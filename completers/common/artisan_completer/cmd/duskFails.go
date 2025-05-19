package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duskFailsCmd = &cobra.Command{
	Use:   "dusk:fails [--browse] [--without-tty] [--pest]",
	Short: "Run the failing Dusk tests from the last run and stop on failure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duskFailsCmd).Standalone()

	duskFailsCmd.Flags().Bool("browse", false, "Open a browser instead of using headless mode")
	duskFailsCmd.Flags().Bool("pest", false, "Run the tests using Pest")
	duskFailsCmd.Flags().Bool("without-tty", false, "Disable output to TTY")
	rootCmd.AddCommand(duskFailsCmd)
}
