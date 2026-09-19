package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugRevlogStatsCmd = &cobra.Command{
	Use:    "debug-revlog-stats",
	Short:  "display statistics about revlogs in the store",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugRevlogStatsCmd).Standalone()

	debugRevlogStatsCmd.Flags().BoolP("changelog", "c", false, "Display changelog statistics")
	debugRevlogStatsCmd.Flags().BoolP("filelogs", "f", false, "Display filelogs statistics")
	debugRevlogStatsCmd.Flags().BoolP("manifest", "m", false, "Display manifest statistics")
	debugRevlogStatsCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugRevlogStatsCmd)
}
