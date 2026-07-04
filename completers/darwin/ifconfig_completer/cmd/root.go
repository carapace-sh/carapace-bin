package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ifconfig",
	Short: "Interface configuration",
	Long:  "https://keith.github.io/xcode-manpages/ifconfig.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("All|-a", "All|-a", false, "")
	rootCmd.Flags().BoolS("C|-C", "C|-C", false, "")
	rootCmd.Flags().BoolS("Down|-d", "Down|-d", false, "")
	rootCmd.Flags().BoolS("List|-l", "List|-l", false, "")
	rootCmd.Flags().BoolS("L|-L", "L|-L", false, "")
	rootCmd.Flags().BoolS("Mask|-m", "Mask|-m", false, "")
	rootCmd.Flags().BoolS("Routing|-r", "Routing|-r", false, "")
	rootCmd.Flags().BoolS("Up|-u", "Up|-u", false, "")
	rootCmd.Flags().BoolS("Verbose|-v", "Verbose|-v", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
