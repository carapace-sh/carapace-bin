package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duskCmd = &cobra.Command{
	Use:   "dusk [--browse] [--without-tty]",
	Short: "Run the Dusk tests for the application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duskCmd).Standalone()

	duskCmd.Flags().Bool("browse", false, "Open a browser instead of using headless mode")
	duskCmd.Flags().Bool("without-tty", false, "Disable output to TTY")
	rootCmd.AddCommand(duskCmd)
}
