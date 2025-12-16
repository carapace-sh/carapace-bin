package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "transmission-create [flags] PATH",
	Short: "A command-line utility to create .torrent files",
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

	rootCmd.Flags().Bool("anonymize", false, "Omit the optional \"created by\" and \"created date\" keys")
	rootCmd.Flags().StringArrayP("comment", "c", nil, "Add a comment to the torrent file")
	rootCmd.Flags().BoolP("help", "h", false, "Show a short help page and exit")
	rootCmd.Flags().StringP("outfile", "o", "", "Name of output torrent file")
	rootCmd.Flags().IntP("piecesize", "s", 0, "Set the size of the torrent pieces in KiB")
	rootCmd.Flags().BoolP("private", "p", false, "Flag the torrent as intended for use on private trackers")
	rootCmd.Flags().StringP("source", "r", "", "Set the torrent's source for private trackers")
	rootCmd.Flags().StringArrayP("tracker", "t", nil, "Add a tracker's announce URL to the torrent")
	rootCmd.Flags().BoolP("version", "V", false, "Show version number and exit")
	rootCmd.Flags().StringArrayP("webseed", "w", nil, "Add a webseed URL")

	rootCmd.MarkFlagsMutuallyExclusive("source", "tracker")
	rootCmd.MarkFlagsOneRequired("source", "tracker")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"outfile": carapace.ActionFiles(".torrent"),
	})
	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
