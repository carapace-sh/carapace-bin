package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kitten",
	Short: "kitten serves as a launcher for running individual kittens",
	Long:  "https://sw.kovidgoyal.net/kitty/kittens_intro/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	rootCmd.Flags().Bool("version", false, "The current kitten version")
}
