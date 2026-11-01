package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_copyDetectionCmd = &cobra.Command{
	Use:   "copy-detection",
	Short: "Show information about file copies detected",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_copyDetectionCmd).Standalone()

	debug_copyDetectionCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_copyDetectionCmd)
}
