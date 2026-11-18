package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Calculate statistics for your history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statsCmd).Standalone()

	statsCmd.Flags().StringP("count", "c", "10", "How many top commands to list")
	statsCmd.Flags().String("filter-mode", "", "Filter commands by scope [global, host, session, directory, workspace]")
	statsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	statsCmd.Flags().StringP("ngram-size", "n", "1", "The number of consecutive commands to consider")
	rootCmd.AddCommand(statsCmd)

	carapace.Gen(statsCmd).FlagCompletion(carapace.ActionMap{
		"filter-mode": carapace.ActionValues("directory", "global", "host", "session", "session-preload", "workspace"),
	})
}
