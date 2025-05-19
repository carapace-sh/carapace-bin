package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/saw"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch <log group>",
	Short: "Continuously stream log events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(watchCmd).Standalone()

	watchCmd.Flags().Bool("expand", false, "indent JSON log messages")
	watchCmd.Flags().String("filter", "", "event filter pattern")
	watchCmd.Flags().Bool("invert", false, "invert colors for light terminal themes")
	watchCmd.Flags().String("prefix", "", "log stream prefix filter")
	watchCmd.Flags().Bool("raw", false, "print raw log event without timestamp or stream prefix")
	watchCmd.Flags().Bool("rawString", false, "print JSON strings without escaping")
	rootCmd.AddCommand(watchCmd)

	carapace.Gen(watchCmd).PositionalCompletion(
		saw.ActionGroups(),
	)
}
