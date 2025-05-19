package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/saw"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <log group>",
	Short: "Get log events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().Bool("expand", false, "indent JSON log messages")
	getCmd.Flags().String("filter", "", "event filter pattern")
	getCmd.Flags().Bool("invert", false, "invert colors for light terminal themes")
	getCmd.Flags().String("prefix", "", "log group prefix filter")
	getCmd.Flags().Bool("pretty", false, "print timestamp and stream name prefix")
	getCmd.Flags().Bool("rawString", false, "print JSON strings without escaping")
	getCmd.Flags().String("start", "", "start getting the logs from this point")
	getCmd.Flags().String("stop", "", "stop getting the logs at this point")
	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"start": time.ActionDate(),
		"stop":  time.ActionDate(),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		saw.ActionGroups(),
	)
}
