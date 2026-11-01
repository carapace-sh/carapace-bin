package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:     "status [-DdegijLpPstvx] [pool...] [interval [count]]",
	Short:   "display detailed pool health status",
	GroupID: "monitoring",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().BoolS("D", "D", false, "display dedup statistics")
	statusCmd.Flags().BoolS("L", "L", false, "resolve symbolic links")
	statusCmd.Flags().BoolS("P", "P", false, "display full paths")
	statusCmd.Flags().StringS("T", "T", "", "timestamp format")
	statusCmd.Flags().StringS("c", "c", "", "run script on each vdev")
	statusCmd.Flags().BoolS("d", "d", false, "display Direct I/O errors")
	statusCmd.Flags().BoolS("e", "e", false, "show only unhealthy vdevs")
	statusCmd.Flags().BoolS("g", "g", false, "display vdev GUIDs")
	statusCmd.Flags().BoolS("i", "i", false, "display initialization status")
	statusCmd.Flags().BoolP("json", "j", false, "JSON output")
	statusCmd.Flags().Bool("json-flat-vdevs", false, "display vdevs in flat hierarchy")
	statusCmd.Flags().Bool("json-int", false, "display numbers as integers")
	statusCmd.Flags().Bool("json-pool-key-guid", false, "use pool GUID as key for pool objects")
	statusCmd.Flags().BoolS("p", "p", false, "display exact values")
	statusCmd.Flags().Bool("power", false, "display vdev enclosure slot power status")
	statusCmd.Flags().BoolS("s", "s", false, "display slow I/O count")
	statusCmd.Flags().BoolS("t", "t", false, "display TRIM status")
	statusCmd.Flags().BoolS("v", "v", false, "verbose error information")
	statusCmd.Flags().BoolS("x", "x", false, "show only pools with errors")

	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"T": carapace.ActionValues("d", "u"),
	})

	carapace.Gen(statusCmd).PositionalCompletion(
		zfs.ActionPools(),
	)
}
