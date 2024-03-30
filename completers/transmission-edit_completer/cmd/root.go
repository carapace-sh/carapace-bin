package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "transmission-edit [flags] torrentfile...",
	Short: "A command-line utility to modify .torrent files' announce URLs",
	Long:  "https://transmissionbt.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("add", "a", "", "Add an announce URL to the torrent's announce-list")
	rootCmd.Flags().StringP("delete", "d", "", "Remove an announce URL from the torrent's announce-list")
	rootCmd.Flags().BoolP("help", "h", false, "Show a short help page and exit")
	rootCmd.Flags().StringP("replace", "r", "", "Substring search-and-replace inside a torrent's announce-list URLs")
	rootCmd.Flags().StringP("source", "s", "", "Sets the source tag within a torrent")
	rootCmd.Flags().BoolP("version", "V", false, "Show version number and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"replace": carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues()
		}),
	})
	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles(".torrent"))
}
