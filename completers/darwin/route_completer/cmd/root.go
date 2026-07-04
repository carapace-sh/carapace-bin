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
	rootCmd.Flags().BoolS("Debug|-d", "Debug|-d", false, "")
	rootCmd.Flags().BoolS("Noresolve|-n", "Noresolve|-n", false, "")
	rootCmd.Flags().BoolS("Quiet|-q", "Quiet|-q", false, "")
	rootCmd.Flags().BoolS("Test|-t", "Test|-t", false, "")
	rootCmd.Flags().BoolS("Verbose|-v", "Verbose|-v", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
