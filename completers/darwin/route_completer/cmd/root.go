package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "route",
	Short: "Routing table management",
	Long:  "https://keith.github.io/xcode-manpages/route.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "Debug mode")
	rootCmd.Flags().BoolS("n", "n", false, "Do not resolve host names")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("t", "t", false, "Test mode (do not modify routing table)")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")
}
