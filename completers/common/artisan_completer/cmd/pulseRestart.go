package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pulseRestartCmd = &cobra.Command{
	Use:   "pulse:restart",
	Short: "Restart any running \"work\" and \"check\" commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulseRestartCmd).Standalone()

	rootCmd.AddCommand(pulseRestartCmd)
}
