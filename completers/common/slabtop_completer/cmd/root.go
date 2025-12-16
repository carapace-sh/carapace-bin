package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slabtop",
	Short: "display kernel slab cache information in real time",
	Long:  "https://man7.org/linux/man-pages/man1/slabtop.1.html",
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

	rootCmd.Flags().StringP("delay", "d", "", "delay updates")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().Bool("human", false, "show human-readable output")
	rootCmd.Flags().BoolP("once", "o", false, "only display once, then exit")
	rootCmd.Flags().StringP("sort", "s", "", "specify sort criteria by character")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"sort": carapace.ActionValuesDescribed(
			"a", "sort by number of active objects",
			"b", "sort by objects per slab",
			"c", "sort by cache size",
			"l", "sort by number of slabs",
			"n", "sort by name",
			"o", "sort by number of objects (the default)",
			"p", "sort by (non display) pages per slab",
			"s", "sort by object size",
			"u", "sort by cache utilization",
			"v", "sort by (non display) number of active slabs",
		),
	})
}
