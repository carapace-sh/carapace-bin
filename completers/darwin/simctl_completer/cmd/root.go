package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simctl",
	Short: "control the iOS Simulator",
	Long:  "https://keith.github.io/xcode-manpages/simctl.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().String("profiles", "", "Specify a custom profiles path")
	rootCmd.Flags().String("set", "", "Specify a custom set path for simulator devices")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")
}
