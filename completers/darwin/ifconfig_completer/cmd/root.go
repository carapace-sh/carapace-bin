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

	rootCmd.Flags().BoolS("C", "C", false, "Show interface configuration")
	rootCmd.Flags().BoolS("L", "L", false, "Show link-layer address lifetime")
	rootCmd.Flags().BoolS("a", "a", false, "Display all interfaces")
	rootCmd.Flags().BoolS("d", "d", false, "Display interfaces that are down")
	rootCmd.Flags().BoolS("l", "l", false, "List interface names")
	rootCmd.Flags().BoolS("m", "m", false, "Display all supported media types")
	rootCmd.Flags().BoolS("r", "r", false, "Show additional reference count")
	rootCmd.Flags().BoolS("u", "u", false, "Display interfaces that are up")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
