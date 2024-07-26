package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pulseCheckCmd = &cobra.Command{
	Use:   "pulse:check [--once]",
	Short: "Take a snapshot of the current server's pulse",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulseCheckCmd).Standalone()

	pulseCheckCmd.Flags().Bool("once", false, "Take a single snapshot")
	rootCmd.AddCommand(pulseCheckCmd)
}
