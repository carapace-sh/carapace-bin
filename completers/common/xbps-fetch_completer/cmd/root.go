package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xbps-fetch [options] <url> <url+N>",
	Short: "XBPS utility to fetch files from URLs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "Enable debug messages to stderr")
	rootCmd.Flags().StringP("out", "o", "", "Rename downloaded file to <file>")
	rootCmd.Flags().BoolP("sha256", "s", false, "Output sha256sums of the files")
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "Show XBPS version")
}
