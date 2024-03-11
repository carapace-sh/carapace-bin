package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devbox",
	Short: "Instant, easy, predictable shells and containers",
	Long:  "https://www.jetpack.io/devbox/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.PersistentFlags().Bool("debug", false, "Show full stack traces on errors")
	rootCmd.Flags().BoolP("help", "h", false, "help for devbox")
}
