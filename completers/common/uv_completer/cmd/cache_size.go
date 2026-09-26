package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cache_sizeCmd = &cobra.Command{
	Use:   "size",
	Short: "Show the cache size",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_sizeCmd).Standalone()

	cache_sizeCmd.Flags().BoolP("human", "H", false, "Display the cache size in human-readable format (e.g., `1.2GiB` instead of raw bytes)")
	cache_sizeCmd.Flags().Bool("human-readable", false, "Display the cache size in human-readable format (e.g., `1.2GiB` instead of raw bytes)")
	cache_sizeCmd.Flags().String("output-format", "auto", "Select the output format")
	cache_sizeCmd.Flag("human-readable").Hidden = true
	cacheCmd.AddCommand(cache_sizeCmd)
	carapace.Gen(cache_sizeCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValuesDescribed(
			"auto", "Display a human-readable size in terminals and raw bytes otherwise",
			"human", "Display the cache size in a human-readable format",
			"machine", "Display the cache size in raw bytes",
		),
	})
}
