package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "asciinema",
	Short: "Terminal session recorder, streamer, and player",
	Long:  "https://asciinema.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Quiet mode - suppress diagnostic messages")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
}
