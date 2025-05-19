package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonPauseCmd = &cobra.Command{
	Use:   "horizon:pause",
	Short: "Pause the master supervisor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonPauseCmd).Standalone()

	rootCmd.AddCommand(horizonPauseCmd)
}
