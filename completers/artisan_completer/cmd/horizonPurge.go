package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var horizonPurgeCmd = &cobra.Command{
	Use:   "horizon:purge [--signal [SIGNAL]]",
	Short: "Terminate any rogue Horizon processes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonPurgeCmd).Standalone()

	horizonPurgeCmd.Flags().String("signal", "", "The signal to send to the rogue processes")
	rootCmd.AddCommand(horizonPurgeCmd)

	carapace.Gen(horizonPurgeCmd).FlagCompletion(carapace.ActionMap{
		"signal": ps.ActionKillSignals(),
	})
}
