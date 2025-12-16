package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vmstat",
	Short: "Report virtual memory statistics",
	Long:  "https://man7.org/linux/man-pages/man8/vmstat.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("active", "a", false, "active/inactive memory")
	rootCmd.Flags().BoolP("disk", "d", false, "disk statistics")
	rootCmd.Flags().BoolP("disk-sum", "D", false, "summarize disk statistics")
	rootCmd.Flags().BoolP("forks", "f", false, "number of forks since boot")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("no-first", "y", false, "skips first line of output")
	rootCmd.Flags().BoolP("one-header", "n", false, "do not redisplay header")
	rootCmd.Flags().StringP("partition", "p", "", "partition specific statistics")
	rootCmd.Flags().BoolP("slabs", "m", false, "slabinfo")
	rootCmd.Flags().BoolP("stats", "s", false, "event counter statistics")
	rootCmd.Flags().BoolP("timestamp", "t", false, "show timestamp")
	rootCmd.Flags().StringP("unit", "S", "", "define display unit")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")
	rootCmd.Flags().BoolP("wide", "w", false, "wide output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"partition": fs.ActionBlockDevices(),
		"unit": carapace.ActionValuesDescribed(
			"k", "1000",
			"K", "1024",
			"m", "1000000",
			"M", "1048576",
		),
	})
}
