package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netstat",
	Short: "Network statistics",
	Long:  "https://keith.github.io/xcode-manpages/netstat.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("Address|-a", "Address|-a", false, "")
	rootCmd.Flags().BoolS("All|-A", "All|-A", false, "")
	rootCmd.Flags().BoolS("Interfaces|-i", "Interfaces|-i", false, "")
	rootCmd.Flags().BoolS("Listen|-l", "Listen|-l", false, "")
	rootCmd.Flags().BoolS("Memory|-m", "Memory|-m", false, "")
	rootCmd.Flags().BoolS("Numeric|-n", "Numeric|-n", false, "")
	rootCmd.Flags().BoolS("Routing|-r", "Routing|-r", false, "")
	rootCmd.Flags().BoolS("Stats|-s", "Stats|-s", false, "")
	rootCmd.Flags().BoolS("Verbose|-v", "Verbose|-v", false, "")
	rootCmd.Flags().BoolS("Wide|-W", "Wide|-W", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
