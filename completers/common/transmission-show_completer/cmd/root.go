package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "transmission-show [flags] torrentfile",
	Short: "A command-line utility to show .torrent file metadata",
	Long:  "https://transmissionbt.com/",
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

	rootCmd.Flags().BoolP("help", "h", false, "Show a short help page and exit")
	rootCmd.Flags().BoolP("magnet", "m", false, "Show a magnet link for the specified .torrent file")
	rootCmd.Flags().BoolP("scrape", "s", false, "Ask the torrent's trackers how many peers are in the torrent's swarm")
	rootCmd.Flags().BoolP("version", "V", false, "Show version number and exit")

	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionFiles(".torrent"))
}
