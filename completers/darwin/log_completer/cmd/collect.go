package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Gather system logs into a log archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(collectCmd).Standalone()

	collectCmd.Flags().Bool("device", false, "Collect logs from first device found")
	collectCmd.Flags().String("device-name", "", "Collect logs from device with the given name")
	collectCmd.Flags().String("device-udid", "", "Collect logs from device with the given UDID")
	collectCmd.Flags().String("last", "", "Collect logs starting <num>[m|h|d] ago")
	collectCmd.Flags().String("output", "", "Output log archive to the given path")
	collectCmd.Flags().String("predicate", "", "Collect logs using a given predicate")
	collectCmd.Flags().String("size", "", "Limit log collection to the given size")
	collectCmd.Flags().String("start", "", "Collect logs starting at the given time")
	rootCmd.AddCommand(collectCmd)

	carapace.Gen(collectCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
