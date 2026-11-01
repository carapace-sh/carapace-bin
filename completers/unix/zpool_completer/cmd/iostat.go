package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var iostatCmd = &cobra.Command{
	Use:     "iostat [-gHLlnpPqrvwy] [-T d|u] [-c script] [pool|vdev...] [interval [count]]",
	Short:   "display I/O statistics",
	GroupID: "monitoring",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iostatCmd).Standalone()

	iostatCmd.Flags().BoolS("H", "H", false, "scripting mode")
	iostatCmd.Flags().BoolS("L", "L", false, "resolve symbolic links")
	iostatCmd.Flags().BoolS("P", "P", false, "display full paths")
	iostatCmd.Flags().StringS("T", "T", "", "timestamp format")
	iostatCmd.Flags().StringS("c", "c", "", "run script on each vdev")
	iostatCmd.Flags().BoolS("g", "g", false, "display vdev GUIDs")
	iostatCmd.Flags().BoolS("l", "l", false, "include average latency")
	iostatCmd.Flags().BoolS("n", "n", false, "print headers only once")
	iostatCmd.Flags().BoolS("p", "p", false, "display exact values")
	iostatCmd.Flags().BoolS("q", "q", false, "include active queue statistics")
	iostatCmd.Flags().BoolS("r", "r", false, "print request size histograms")
	iostatCmd.Flags().BoolS("v", "v", false, "verbose vdev statistics")
	iostatCmd.Flags().BoolS("w", "w", false, "display latency histograms")
	iostatCmd.Flags().BoolS("y", "y", false, "suppress initial statistics")

	rootCmd.AddCommand(iostatCmd)

	carapace.Gen(iostatCmd).FlagCompletion(carapace.ActionMap{
		"T": carapace.ActionValues("d", "u"),
	})

	carapace.Gen(iostatCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
